package admin

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/amankhys/multi_vendor_ecommerce_go/pkg/utils"
	"github.com/amankhys/multi_vendor_ecommerce_go/pkg/validators"
	"github.com/amankhys/multi_vendor_ecommerce_go/repository/db"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Admin struct{ DB *db.Queries }

func (a *Admin) AdminAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	var resp struct {
		Data    []db.GetAllUsersRow `json:"data"`
		Message string              `json:"message"`
	}
	data, err := a.DB.GetAllUsers(context.TODO())
	if err == sql.ErrNoRows {
		message := "no current users available to display"
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(message))
	} else if err != nil {
		log.Warn(err)
		http.Error(w, fmt.Errorf("failed : %w", err).Error(), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Message = "successfully fetched all users"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (a *Admin) AdminUsersHandler(w http.ResponseWriter, r *http.Request) {
	var resp struct {
		Data    []db.GetAllUsersByRoleUserRow `json:"data"`
		Message string                        `json:"message"`
	}
	data, err := a.DB.GetAllUsersByRoleUser(context.TODO(), "user")
	if err == sql.ErrNoRows {
		w.Header().Set("Content-Type", "text/plain")
		message := "no current users available"
		w.Write([]byte(message))
	} else if err != nil {
		log.Warn(err)
		http.Error(w, fmt.Errorf("failed : %w", err).Error(), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Message = "successfully fetched all users"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (a *Admin) AdminSellersHandler(w http.ResponseWriter, r *http.Request) {
	var resp struct {
		Data    []db.GetAllUsersByRoleSellerRow `json:"data"`
		Message string                          `json:"message"`
	}
	data, err := a.DB.GetAllUsersByRoleSeller(context.TODO(), "seller")
	if err == sql.ErrNoRows {
		w.Header().Set("Content-Type", "text/plain")
		message := "no sellers available"
		w.Write([]byte(message))
	} else if err != nil {
		log.Warn(err)
		http.Error(w, fmt.Errorf("failed : %w", err).Error(), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Message = "successfully fetched all sellers"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
func (a *Admin) VerifySellerHandler(w http.ResponseWriter, r *http.Request) {
	// check if the request is in correct format
	var req struct {
		Email string `json:"email"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest)
		return
	} else if !validators.ValidateEmail(req.Email) {
		http.Error(w, "invalid email format:", http.StatusBadRequest)
		return
	}

	// check if the seller exists and is not verified
	user, err := a.DB.GetUserByEmail(context.TODO(), req.Email)
	if err == sql.ErrNoRows {
		http.Error(w, "seller does not exist.", http.StatusBadRequest)
		return
	} else if err != nil {
		log.Warn("error fetching user by email")
		http.Error(w, "internal server error while fetching user", http.StatusInternalServerError)
		return
	} else if user.Role != utils.SellerRole {
		http.Error(w, "the given details is not that of a seller", http.StatusBadRequest)
		return
	} else if user.UserVerified {
		http.Error(w, "seller already verified", http.StatusBadRequest)
		return
	} else if !user.EmailVerified {
		http.Error(w, "seller email not yet verified.", http.StatusUnauthorized)
		return
	}

	// verify seller
	seller, err := a.DB.VerifySellerByID(context.TODO(), user.ID)
	if err != nil {
		log.Warn("verify seller by id failed for a valid seller.")
		http.Error(w, "internal server error while verifying seller", http.StatusInternalServerError)
		return
	}
	// make errors and messages slice for the response
	var Err []string
	var Messages []string
	// add wallet for the seller
	wallet, err := a.DB.AddWalletByUserID(context.TODO(), seller.ID)
	if err != nil {
		log.Warn("error adding wallet for seller after verifying account:", err.Error())
		Err = append(Err, "error adding wallet for seller after verifying account")
	} else {
		Messages = append(Messages, "successfully added wallet for seller")
		Messages = append(Messages, "walletID:", wallet.ID.String(), fmt.Sprintf("savings: %v", wallet.Savings))
	}
	var resp struct {
		Data     db.VerifySellerByIDRow `json:"data"`
		Messages []string               `json:"message"`
		Err      []string               `json:"errors"`
	}
	resp.Data = seller
	resp.Messages = append(Messages, "successfully verified seller")
	resp.Err = Err
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (a *Admin) AdminProductsHandler(w http.ResponseWriter, r *http.Request) {
	var resp struct {
		Data    []db.Product `json:"data"`
		Message string       `json:"message"`
	}
	data, err := a.DB.GetAllProductsForAdmin(context.TODO())
	if err != nil {
		log.Warn(err)
		http.Error(w, fmt.Errorf("failed : %w", err).Error(), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Message = "successfully fetched all products"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func (a *Admin) AdminCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var resp struct {
		Data    []db.Category `json:"data"`
		Message string        `json:"message"`
	}
	data, err := a.DB.GetAllCategoriesForAdmin(context.TODO())
	if err != nil {
		log.Warn(err)
		http.Error(w, fmt.Errorf("failed : %w", err).Error(), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Message = "successfully fetched all categories"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (a *Admin) BlockUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID uuid.UUID `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid data format", http.StatusBadRequest)
		return
	}

	user, err := a.DB.GetUserById(context.TODO(), req.UserID)
	if err == sql.ErrNoRows {
		http.Error(w, "invalid userID", http.StatusBadRequest)
		return
	} else if err != nil {
		log.Warn("error taking user from db: ", err)
		http.Error(w, "error fetching user", http.StatusInternalServerError)
		return
	} else if user.IsBlocked {
		http.Error(w, "user already blocked", http.StatusBadRequest)
		return
	} else if user.Role == utils.AdminRole {
		http.Error(w, "trying to block admin: invalid request", http.StatusBadRequest)
		return
	}

	blockedUser, err := a.DB.BlockUserByID(context.TODO(), req.UserID)
	if err != nil {
		log.Warn(err)
		http.Error(w, "error blocking user", http.StatusInternalServerError)
		return
	}
	log.Infof("blocked user: %s", blockedUser.Email)
	message := fmt.Sprintf("succesfully blocked user: %s", blockedUser.ID.String())
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func (a *Admin) UnblockUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID uuid.UUID `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid data format", http.StatusBadRequest)
		return
	}

	user, err := a.DB.UnblockUserByID(context.TODO(), req.UserID)
	if err == sql.ErrNoRows {
		http.Error(w, "invalid user data", http.StatusBadRequest)
		return
	} else if err != nil {
		log.Warn(err)
		http.Error(w, "error unblocking user", http.StatusInternalServerError)
		return
	}
	log.Infof("unblocked user: %s", user.ID.String())
	message := fmt.Sprintf("succesfully unblocked user: %s", user.ID.String())
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func (a *Admin) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ProductID uuid.UUID `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid data format", http.StatusBadRequest)
		return
	}
	product, err := a.DB.DeleteProductByID(context.TODO(), req.ProductID)
	if err == sql.ErrNoRows {
		http.Error(w, "invalid productID", http.StatusBadRequest)
		return
	} else if err != nil {
		log.Warn(err)
		http.Error(w, "error deleting product", http.StatusInternalServerError)
		return
	}
	log.Infof("deleted product: %s", product.ID.String())
	w.Header().Set("Content-Type", "application/json")
	message := fmt.Sprintf("product: %s deleted", product.Name)
	w.Write([]byte(message))
}

func (a *Admin) AddCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid data format", http.StatusBadRequest)
		return
	}
	req.Name = strings.ToLower(req.Name)
	category, err := a.DB.AddCateogry(context.TODO(), req.Name)
	if err != nil {
		log.Warn(err)
		http.Error(w, fmt.Errorf("failed to add cateogry: %w", err).Error(), http.StatusBadRequest)
		return
	}
	log.Infof("added category: %s", category.Name)
	w.Header().Set("Content-Type", "text/plain")
	message := fmt.Sprintf("category: %s added", category.Name)
	w.Write([]byte(message))
}

func (a *Admin) EditCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var req db.EditCategoryNameByNameParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid data format", http.StatusBadRequest)
		return
	}
	req.Name = strings.ToLower(req.Name)
	req.NewName = strings.ToLower(req.NewName)
	category, err := a.DB.EditCategoryNameByName(context.TODO(), req)
	if err != nil {
		log.Warn(err)
		http.Error(w, fmt.Errorf("failed to rename cateogry: %w", err).Error(), http.StatusBadRequest)
		return
	}
	log.Infof("renamed category: %s", category.Name)
	w.Header().Set("Content-Type", "text/plain")
	message := fmt.Sprintf("category: %s renamed to %s", req.Name, category.Name)
	w.Write([]byte(message))
}

func (a *Admin) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CategoryName string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid data format", http.StatusBadRequest)
		return
	}
	req.CategoryName = strings.ToLower(req.CategoryName)
	category, err := a.DB.DeleteCategoryByName(context.TODO(), req.CategoryName)
	if err == sql.ErrNoRows {
		http.Error(w, "invalid category name", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, fmt.Errorf("failed to delete category: %w", err).Error(), http.StatusBadRequest)
		return
	}
	log.Infof("deleted category: %s", category.Name)
	w.Header().Set("Content-Type", "text/plain")
	message := fmt.Sprintf("category: %s deleted", category.Name)
	w.Write([]byte(message))
}

func (a *Admin) GetOrderItemsHandler(w http.ResponseWriter, r *http.Request) {
	user := helper.GetUserHelper(w, r)
	if user.ID == uuid.Nil {
		return
	}

	// fetch order_items
	orderItems, err := a.DB.GetAllOrderForAdmin(context.TODO())
	if err != nil {
		log.Warn("error fetching orders for admin in GetOrderItemsHandler:", err.Error())
		http.Error(w, "internal server error fetching orders for admin", http.StatusInternalServerError)
		return
	}

	// make resp orderItem struct
	type respOrderItem struct {
		ID          uuid.UUID `json:"id"`
		OrderID     uuid.UUID `json:"order_id"`
		Status      string    `json:"status"`
		ProductID   uuid.UUID `json:"product_id"`
		Price       float64   `json:"price"`
		Quantity    int       `json:"quantity"`
		TotalAmount float64   `json:"total_amount"`
	}

	// respOrderItems slice
	var respOrderItems []respOrderItem
	for _, v := range orderItems {
		var temp respOrderItem
		temp.ID = v.ID
		temp.OrderID = v.OrderID
		temp.Status = v.Status
		temp.ProductID = v.ProductID
		temp.Price = v.Price
		temp.Quantity = int(v.Quantity)
		temp.TotalAmount = v.TotalAmount

		respOrderItems = append(respOrderItems, temp)
	}

	// give response
	var resp struct {
		Data    []respOrderItem `json:"data"`
		Message string          `json:"message"`
	}
	resp.Data = respOrderItems
	resp.Message = "successfully fetched all orderItems"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (a *Admin) DeliverOrderItemHandler(w http.ResponseWriter, r *http.Request) {
	user := helper.GetUserHelper(w, r)
	if user.ID == uuid.Nil {
		return
	}

	var req struct {
		OrderItemIDStr string `json:"order_item_id"`
	}
	req.OrderItemIDStr = r.URL.Query().Get("order_item_id")
	orderID, err := uuid.Parse(req.OrderItemIDStr)
	if err != nil {
		http.Error(w, "not a valid orderItemID", http.StatusBadRequest)
		return
	}

	orderItem, err := a.DB.GetOrderItemByID(context.TODO(), orderID)
	if err == sql.ErrNoRows {
		http.Error(w, "not a valid orderItemID", http.StatusBadRequest)
		return
	} else if err != nil {
		log.Warn("error fetching orderItemByID in admin to change orderStatus:", err.Error())
		http.Error(w, "internal server error fetching orderItem", http.StatusInternalServerError)
		return
	}

	// checking order item status to change the status to delivered
	if orderItem.Status == utils.StatusOrderCancelled ||
		orderItem.Status == utils.StatusOrderPending ||
		orderItem.Status == utils.StatusOrderProcessing ||
		orderItem.Status == utils.StatusOrderDelivered {
		msg := fmt.Sprintf("order %s. Cannot change to status to delivered", orderItem.Status)
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}
	// no need to check the orderItem status is shipped since it is the
	// default case if all the cases above failed

	var arg db.ChangeOrderItemStatusByIDParams
	arg.ID = orderItem.ID
	arg.Status = utils.StatusOrderDelivered
	updatedOrderItem, err := a.DB.ChangeOrderItemStatusByID(context.TODO(), arg)
	if err != nil {
		log.Warn("error changing order status to delivered in DeliverOrderItemHandler in admin:", err.Error())
		http.Error(w, "internal error changing order status to delivered", http.StatusInternalServerError)
		return
	}
	type respOrderItem struct {
		OrderItemID uuid.UUID `json:"order_item_id"`
		Status      string    `json:"status"`
		ProductID   uuid.UUID `json:"product_id"`
		Price       float64   `json:"price"`
		Qauntity    int       `json:"quantity"`
		TotalAmount float64   `json:"total_amount"`
	}
	var respUpdatedOrderItem respOrderItem
	respUpdatedOrderItem.OrderItemID = updatedOrderItem.ID
	respUpdatedOrderItem.ProductID = updatedOrderItem.ProductID
	respUpdatedOrderItem.Price = updatedOrderItem.Price
	respUpdatedOrderItem.Qauntity = int(updatedOrderItem.Quantity)
	respUpdatedOrderItem.TotalAmount = updatedOrderItem.TotalAmount
	respUpdatedOrderItem.Status = updatedOrderItem.Status
	var resp struct {
		Data    respOrderItem `json:"data"`
		Message string
	}
	resp.Data = respUpdatedOrderItem
	resp.Message = "successfully updated the orderItem to status delivered"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
