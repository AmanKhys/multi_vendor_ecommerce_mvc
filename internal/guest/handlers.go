package guest

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/amankhys/multi_vendor_ecommerce_go/pkg/mail"
	"github.com/amankhys/multi_vendor_ecommerce_go/pkg/sessions"
	"github.com/amankhys/multi_vendor_ecommerce_go/pkg/utils"
	"github.com/amankhys/multi_vendor_ecommerce_go/pkg/validators"
	"github.com/amankhys/multi_vendor_ecommerce_go/repository/db"
	log "github.com/sirupsen/logrus"
)

var RoleSeller = "seller"
var RoleUser = "user"

type Guest struct {
	DB *db.Queries
}

func (g *Guest) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	message := "hello there"
	w.Write([]byte(message))
}

func (g *Guest) UserSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var req db.AddUserParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid data format:"+err.Error(), http.StatusBadRequest)
		return
	} else if !validators.ValidateEmail(req.Email) {
		http.Error(w, "invalid email format:", http.StatusBadRequest)
		return
	} else if !validators.ValidateName(req.Name) {
		http.Error(w, "invalid Name format:", http.StatusBadRequest)
		return
	} else if !validators.ValidatePassword(req.Password) {
		http.Error(w, "invalid password format:", http.StatusBadRequest)
		return
	} else if req.Phone.Valid && !validators.ValidatePhone(strconv.Itoa(int(req.Phone.Int64))) {
		http.Error(w, "invalid phone format:", http.StatusBadRequest)
		return
	}
	user, _ := g.DB.GetUserByEmail(context.TODO(), req.Email)
	if req.Email == user.Email && user.UserVerified {
		http.Error(w, "user already exists and verified", http.StatusBadRequest)
		return
	} else if req.Email == user.Email && !user.UserVerified {
		http.Error(w, "user already exists and not verified. visit /user_signup_otp and verify user", http.StatusBadRequest)
		return
	}

	log.Info("reached here mf")
	hashed, err := utils.HashPassword(req.Password)
	log.Info("reached here mf again")
	if err != nil {
		log.Warn("error hashing password")
		http.Error(w, "error hashing password"+err.Error(), http.StatusInternalServerError)
		return
	}
	var arg = db.AddUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
	}
	if req.Phone.Valid {
		arg.Phone = req.Phone
	}
	var respUser db.AddUserRow
	respUser, err = g.DB.AddUser(context.TODO(), arg)
	if err != nil {
		log.Warn("user not added")
		http.Error(w, "internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	type resp struct {
		Data    db.AddUserRow `json:"data"`
		Message string        `json:"message"`
	}
	var response = resp{
		Data:    respUser,
		Message: "successfully added user. Now you need to verify it",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	signupOTP, err := g.DB.AddOTP(context.TODO(), respUser.ID)
	if err == sql.ErrNoRows {
		log.Warn("otp not generated")
		return
	}
	err = mail.SendOTPMail(int(signupOTP.Otp), signupOTP.ExpiresAt, respUser.Email)
	if err != nil {
		log.Warn("failed to send otp:", err.Error())
	}
}

func (g *Guest) SellerSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var req db.AddSellerParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid data format:"+err.Error(), http.StatusBadRequest)
		return
	} else if !validators.ValidateEmail(req.Email) {
		http.Error(w, "invalid email format:", http.StatusBadRequest)
		return
	} else if !validators.ValidateName(req.Name) {
		http.Error(w, "invalid Name format:", http.StatusBadRequest)
		return
	} else if !validators.ValidatePassword(req.Password) {
		http.Error(w, "invalid password format:", http.StatusBadRequest)
		return
	} else if !validators.ValidatePhone(strconv.Itoa(int(req.Phone.Int64))) || !req.Phone.Valid {
		http.Error(w, "invalid phone format:", http.StatusBadRequest)
		return
	} else if !validators.ValidateGSTNo(req.GstNo.String) || !req.GstNo.Valid {
		http.Error(w, "invalid gst_no format:", http.StatusBadRequest)
		return
	} else if req.About.Valid && req.GstNo.String == "" {
		http.Error(w, "invalid about format: about empty", http.StatusBadRequest)
		return
	}
	user, _ := g.DB.GetUserByEmail(context.TODO(), req.Email)
	if user.Role != "" && user.Role != RoleSeller {
		http.Error(w, "signing up on a already existing user. not allowed", http.StatusBadRequest)
		return
	}
	if req.Email == user.Email && user.EmailVerified {
		http.Error(w, "seller already exists and email verified", http.StatusBadRequest)
		return
	} else if req.Email == user.Email && !user.EmailVerified {
		http.Error(w, "seller already exists and email not verified. visit /user_signup_otp and verify user", http.StatusBadRequest)
		return
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Warn("error hashing password")
		http.Error(w, "error hashing password"+err.Error(), http.StatusInternalServerError)
		return
	}
	var arg = db.AddSellerParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
		Phone:    req.Phone,
		GstNo:    req.GstNo,
		About:    req.About,
	}
	var respSeller db.AddSellerRow
	respSeller, err = g.DB.AddSeller(context.TODO(), arg)
	if err != nil {
		log.Warn("user not added")
		http.Error(w, "internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}
	type resp struct {
		Data    db.AddSellerRow `json:"data"`
		Message string          `json:"message"`
	}
	var response = resp{
		Data:    respSeller,
		Message: "successfully added user. Now you need to verify it. Check email for otp.",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	signupOTP, err := g.DB.AddOTP(context.TODO(), respSeller.ID)
	if err == sql.ErrNoRows {
		log.Warn("otp not generated")
		return
	} else if err != nil {
		log.Warn("error processing AddOTP query")
		return
	}

	err = mail.SendOTPMail(int(signupOTP.Otp), signupOTP.ExpiresAt, user.Email)
	if err != nil {
		log.Warn("failed to send otp", err.Error())
		result, err := g.DB.DeleteOTPByEmail(context.TODO(), user.Email)
		if err != nil {
			log.Warn("error deleting otp by email")
		}
		k, err := result.RowsAffected()
		if err != nil {
			log.Warn("error fetching the rows affected from DeleteOTPByEmail query resut")
		}
		if k == 0 {
			log.Warn("no rows affected while operating DeleteOTPByEmail query")
		}
	}
}

func (g *Guest) UserSignUpOTPHandler(w http.ResponseWriter, r *http.Request) {
	// get req.Body and check if it's in correct format
	var req struct {
		Email string `json:"email"`
		Otp   int    `json:"otp"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	if !validators.ValidateEmail(req.Email) {
		http.Error(w, "invalid email format:", http.StatusBadRequest)
		return
	} else if !validators.ValidateOTP(req.Otp) {
		http.Error(w, "invalid OTP format:", http.StatusBadRequest)
		return
	}

	user, err := g.DB.GetUserByEmail(context.TODO(), req.Email)
	if err != nil {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	} else if user.EmailVerified {
		http.Error(w, "user email already verified", http.StatusBadRequest)
		return
	}

	otp, err := g.DB.GetValidOTPByUserID(context.TODO(), user.ID)
	if err == sql.ErrNoRows {
		http.Error(w, "no valid otp available. generating another otp", http.StatusBadRequest)
		otp, err := g.DB.AddOTP(context.TODO(), user.ID)
		if err == sql.ErrNoRows {
			log.Warn("no otp generated")
			return
		}
		err = mail.SendOTPMail(int(otp.Otp), otp.ExpiresAt, user.Email)
		if err != nil {
			log.Warn("error sending otp:", err.Error())
			http.Error(w, "error sending otp:", http.StatusInternalServerError)
			return
		}
		log.Info("testing otp generated: ", otp) // for testing
		return
	} else if err != nil {
		log.Warn("error fetching otp")
		http.Error(w, "internal server error fetching otp", http.StatusInternalServerError)
		return
	}
	if req.Otp != int(otp.Otp) {
		http.Error(w, "invalid otp", http.StatusBadRequest)
		return
	}

	// verify user
	respUser, err := g.DB.VerifyUserByID(context.TODO(), user.ID)
	if err != nil {
		log.Warn("error verifying a valid user")
		http.Error(w, "internal server error verifying user", http.StatusInternalServerError)
		return
	}

	// send response
	var resp struct {
		Data    db.VerifyUserByIDRow `json:"data"`
		Message string               `json:"message"`
	}
	resp.Data = respUser
	resp.Message = "user verified successfully"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (g *Guest) SellerSignUpOTPHandler(w http.ResponseWriter, r *http.Request) {
	// get req.Body and check if it's in correct format
	var req struct {
		Email string `json:"email"`
		Otp   int    `json:"otp"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	if !validators.ValidateEmail(req.Email) {
		http.Error(w, "invalid email format:", http.StatusBadRequest)
		return
	} else if !validators.ValidateOTP(req.Otp) {
		http.Error(w, "invalid OTP format:", http.StatusBadRequest)
		return
	}

	user, err := g.DB.GetUserByEmail(context.TODO(), req.Email)
	if err != nil {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	} else if user.EmailVerified {
		http.Error(w, "user email already verified", http.StatusBadRequest)
		return
	}

	// validate otp
	otp, err := g.DB.GetValidOTPByUserID(context.TODO(), user.ID)
	if err == sql.ErrNoRows {
		http.Error(w, "no valid otp available. generating another otp", http.StatusBadRequest)
		otp, err := g.DB.AddOTP(context.TODO(), user.ID)
		if err == sql.ErrNoRows {
			log.Warn("no otp generated")
			return
		}
		err = mail.SendOTPMail(int(otp.Otp), otp.ExpiresAt, user.Email)
		if err != nil {
			result, err := g.DB.DeleteOTPByEmail(context.TODO(), user.Email)
			if err != nil {
				log.Warn("error deleting otp:", err)
			}
			k, err := result.RowsAffected()
			if err != nil {
				log.Warn("error fetching rows affected from sql.Result:", err)
			} else if k == 0 {
				log.Warn("no otp deleted after successful query execution")
			}
			log.Warn("error sending otp email to seller", err.Error())
			http.Error(w, "error sending otp email", http.StatusInternalServerError)
			return
		}
		log.Info("testing otp generated: ", otp) // for testing
		return
	} else if err != nil {
		log.Warn("error fetching otp")
		http.Error(w, "internal server error fetching otp", http.StatusInternalServerError)
		return
	}
	if req.Otp != int(otp.Otp) {
		http.Error(w, "invalid otp", http.StatusBadRequest)
		return
	}

	// verify user
	respSeller, err := g.DB.VerifySellerEmailByID(context.TODO(), user.ID)
	if err != nil {
		log.Warn("error verifying a valid user")
		http.Error(w, "internal server error verifying user", http.StatusInternalServerError)
		return
	}

	// send response
	var resp struct {
		Data    db.VerifySellerEmailByIDRow `json:"data"`
		Message string                      `json:"message"`
	}
	resp.Data = respSeller
	resp.Message = "seller email verified successfully. Wait for admin to verify the seller."

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func (g *Guest) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// take request
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	if !validators.ValidateEmail(req.Email) {
		http.Error(w, "invalid email format", http.StatusBadRequest)
		return
	} else if !validators.ValidatePassword(req.Password) {
		http.Error(w, "invalid password format", http.StatusBadRequest)
		return
	}

	// compare email and password
	user, err := g.DB.GetUserWithPasswordByEmail(context.TODO(), req.Email)
	if err != nil {
		http.Error(w, "invalid email", http.StatusUnauthorized)
		return
	}

	err = utils.ComparePassword(req.Password, user.Password)
	if err != nil {
		log.Warn(err)
		http.Error(w, "wrong password", http.StatusUnauthorized)
		return
	} else if !user.UserVerified {
		message := fmt.Sprintf("%s not verified", user.Role)
		http.Error(w, message, http.StatusUnauthorized)
		return
	}

	var arg = db.AddSessionParams{
		UserID:    user.ID,
		IpAddress: utils.GetClientIPString(r),
		UserAgent: utils.GetUserAgent(r),
	}
	session, err := g.DB.AddSession(context.TODO(), arg)
	if err != nil {
		http.Error(w, "invalid data format", http.StatusBadRequest)
		return
	}
	sessions.SetSessionCookie(w, session.ID.String())
	w.Header().Set("Content-Type", "text/plain")
	message := fmt.Sprintf("%s of id: %s has successfully logged in\n", user.Role, user.ID.String())
	w.Write([]byte(message))
}

func (g *Guest) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := sessions.GetSessionCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sessions.DeleteSessionCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
