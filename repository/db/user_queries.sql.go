// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user_queries.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addAndVerifyUser = `-- name: AddAndVerifyUser :one
insert into users
(name, email, password, role, email_verified, user_verified)
values  ($1, $2, $3, 'user', true, true)
returning id, name, email, role, is_blocked, email_verified, user_verified, created_at, updated_at
`

type AddAndVerifyUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddAndVerifyUserRow struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Role          string    `json:"role"`
	IsBlocked     bool      `json:"is_blocked"`
	EmailVerified bool      `json:"email_verified"`
	UserVerified  bool      `json:"user_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (q *Queries) AddAndVerifyUser(ctx context.Context, arg AddAndVerifyUserParams) (AddAndVerifyUserRow, error) {
	row := q.queryRow(ctx, q.addAndVerifyUserStmt, addAndVerifyUser, arg.Name, arg.Email, arg.Password)
	var i AddAndVerifyUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Role,
		&i.IsBlocked,
		&i.EmailVerified,
		&i.UserVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const addSeller = `-- name: AddSeller :one
INSERT INTO users
(name, email, phone, password, role, gst_no, about)
VALUES ($1, $2, $3, $4, 'seller', $5, $6)
RETURNING id, name, email, phone, role, is_blocked, email_verified, user_verified, gst_no, about, created_at, updated_at
`

type AddSellerParams struct {
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Phone    sql.NullInt64  `json:"phone"`
	Password string         `json:"password"`
	GstNo    sql.NullString `json:"gst_no"`
	About    sql.NullString `json:"about"`
}

type AddSellerRow struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         sql.NullInt64  `json:"phone"`
	Role          string         `json:"role"`
	IsBlocked     bool           `json:"is_blocked"`
	EmailVerified bool           `json:"email_verified"`
	UserVerified  bool           `json:"user_verified"`
	GstNo         sql.NullString `json:"gst_no"`
	About         sql.NullString `json:"about"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (q *Queries) AddSeller(ctx context.Context, arg AddSellerParams) (AddSellerRow, error) {
	row := q.queryRow(ctx, q.addSellerStmt, addSeller,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.Password,
		arg.GstNo,
		arg.About,
	)
	var i AddSellerRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.IsBlocked,
		&i.EmailVerified,
		&i.UserVerified,
		&i.GstNo,
		&i.About,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const addUser = `-- name: AddUser :one
INSERT INTO users
(name, email, phone, password, role)
VALUES ($1, $2, $3, $4, 'user')
RETURNING id, name, email, phone, role, is_blocked, email_verified, user_verified, created_at, updated_at
`

type AddUserParams struct {
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Phone    sql.NullInt64 `json:"phone"`
	Password string        `json:"password"`
}

type AddUserRow struct {
	ID            uuid.UUID     `json:"id"`
	Name          string        `json:"name"`
	Email         string        `json:"email"`
	Phone         sql.NullInt64 `json:"phone"`
	Role          string        `json:"role"`
	IsBlocked     bool          `json:"is_blocked"`
	EmailVerified bool          `json:"email_verified"`
	UserVerified  bool          `json:"user_verified"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) (AddUserRow, error) {
	row := q.queryRow(ctx, q.addUserStmt, addUser,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.Password,
	)
	var i AddUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.IsBlocked,
		&i.EmailVerified,
		&i.UserVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const blockUserByID = `-- name: BlockUserByID :one
UPDATE users
SET is_blocked = true, updated_at = current_timestamp
WHERE id = $1
RETURNING id, name, email, phone, role, is_blocked, email_verified, user_verified, gst_no, about, created_at, updated_at
`

type BlockUserByIDRow struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         sql.NullInt64  `json:"phone"`
	Role          string         `json:"role"`
	IsBlocked     bool           `json:"is_blocked"`
	EmailVerified bool           `json:"email_verified"`
	UserVerified  bool           `json:"user_verified"`
	GstNo         sql.NullString `json:"gst_no"`
	About         sql.NullString `json:"about"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (q *Queries) BlockUserByID(ctx context.Context, id uuid.UUID) (BlockUserByIDRow, error) {
	row := q.queryRow(ctx, q.blockUserByIDStmt, blockUserByID, id)
	var i BlockUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.IsBlocked,
		&i.EmailVerified,
		&i.UserVerified,
		&i.GstNo,
		&i.About,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, name, email, phone, role, is_blocked, email_verified, user_verified, gst_no, about, created_at, updated_at FROM users
`

type GetAllUsersRow struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         sql.NullInt64  `json:"phone"`
	Role          string         `json:"role"`
	IsBlocked     bool           `json:"is_blocked"`
	EmailVerified bool           `json:"email_verified"`
	UserVerified  bool           `json:"user_verified"`
	GstNo         sql.NullString `json:"gst_no"`
	About         sql.NullString `json:"about"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (q *Queries) GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error) {
	rows, err := q.query(ctx, q.getAllUsersStmt, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllUsersRow{}
	for rows.Next() {
		var i GetAllUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.Role,
			&i.IsBlocked,
			&i.EmailVerified,
			&i.UserVerified,
			&i.GstNo,
			&i.About,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllUsersByRole = `-- name: GetAllUsersByRole :many
SELECT id, name, email, phone, role, is_blocked, email_verified, user_verified, gst_no, about, created_at, updated_at FROM users
WHERE role = $1
`

type GetAllUsersByRoleRow struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         sql.NullInt64  `json:"phone"`
	Role          string         `json:"role"`
	IsBlocked     bool           `json:"is_blocked"`
	EmailVerified bool           `json:"email_verified"`
	UserVerified  bool           `json:"user_verified"`
	GstNo         sql.NullString `json:"gst_no"`
	About         sql.NullString `json:"about"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (q *Queries) GetAllUsersByRole(ctx context.Context, role string) ([]GetAllUsersByRoleRow, error) {
	rows, err := q.query(ctx, q.getAllUsersByRoleStmt, getAllUsersByRole, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllUsersByRoleRow{}
	for rows.Next() {
		var i GetAllUsersByRoleRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.Role,
			&i.IsBlocked,
			&i.EmailVerified,
			&i.UserVerified,
			&i.GstNo,
			&i.About,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCurrentTimestamp = `-- name: GetCurrentTimestamp :one
select current_timestamp
`

func (q *Queries) GetCurrentTimestamp(ctx context.Context) (interface{}, error) {
	row := q.queryRow(ctx, q.getCurrentTimestampStmt, getCurrentTimestamp)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const getSellerByProductID = `-- name: GetSellerByProductID :one
SELECT u.id, u.name, u.email, u.phone, u.role, u.is_blocked, u.email_verified, u.user_verified, u.gst_no, u.about, u.created_at, u.updated_at
FROM  products p
INNER JOIN  users u
on p.seller_id = u.id and u.role = 'seller' and p.is_deleted = false
where p.id = $1
`

type GetSellerByProductIDRow struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         sql.NullInt64  `json:"phone"`
	Role          string         `json:"role"`
	IsBlocked     bool           `json:"is_blocked"`
	EmailVerified bool           `json:"email_verified"`
	UserVerified  bool           `json:"user_verified"`
	GstNo         sql.NullString `json:"gst_no"`
	About         sql.NullString `json:"about"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (q *Queries) GetSellerByProductID(ctx context.Context, id uuid.UUID) (GetSellerByProductIDRow, error) {
	row := q.queryRow(ctx, q.getSellerByProductIDStmt, getSellerByProductID, id)
	var i GetSellerByProductIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.IsBlocked,
		&i.EmailVerified,
		&i.UserVerified,
		&i.GstNo,
		&i.About,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, phone, role, is_blocked, email_verified, user_verified, gst_no, about, created_at, updated_at FROM users
WHERE email = $1
`

type GetUserByEmailRow struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         sql.NullInt64  `json:"phone"`
	Role          string         `json:"role"`
	IsBlocked     bool           `json:"is_blocked"`
	EmailVerified bool           `json:"email_verified"`
	UserVerified  bool           `json:"user_verified"`
	GstNo         sql.NullString `json:"gst_no"`
	About         sql.NullString `json:"about"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.IsBlocked,
		&i.EmailVerified,
		&i.UserVerified,
		&i.GstNo,
		&i.About,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, email, phone, role, is_blocked, email_verified, user_verified, gst_no, about, created_at, updated_at FROM users
WHERE id = $1
`

type GetUserByIdRow struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         sql.NullInt64  `json:"phone"`
	Role          string         `json:"role"`
	IsBlocked     bool           `json:"is_blocked"`
	EmailVerified bool           `json:"email_verified"`
	UserVerified  bool           `json:"user_verified"`
	GstNo         sql.NullString `json:"gst_no"`
	About         sql.NullString `json:"about"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (GetUserByIdRow, error) {
	row := q.queryRow(ctx, q.getUserByIdStmt, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.IsBlocked,
		&i.EmailVerified,
		&i.UserVerified,
		&i.GstNo,
		&i.About,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserWithPasswordByEmail = `-- name: GetUserWithPasswordByEmail :one
SELECT id, name, email, phone, password, role, email_verified, user_verified, is_blocked, gst_no, about, created_at, updated_at FROM users
WHERE email = $1
`

func (q *Queries) GetUserWithPasswordByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserWithPasswordByEmailStmt, getUserWithPasswordByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.Role,
		&i.EmailVerified,
		&i.UserVerified,
		&i.IsBlocked,
		&i.GstNo,
		&i.About,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsersByRole = `-- name: GetUsersByRole :many
SELECT id, name, email, phone, role, is_blocked, email_verified, user_verified, gst_no, about, created_at, updated_at FROM users
WHERE role = $1
`

type GetUsersByRoleRow struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         sql.NullInt64  `json:"phone"`
	Role          string         `json:"role"`
	IsBlocked     bool           `json:"is_blocked"`
	EmailVerified bool           `json:"email_verified"`
	UserVerified  bool           `json:"user_verified"`
	GstNo         sql.NullString `json:"gst_no"`
	About         sql.NullString `json:"about"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (q *Queries) GetUsersByRole(ctx context.Context, role string) ([]GetUsersByRoleRow, error) {
	rows, err := q.query(ctx, q.getUsersByRoleStmt, getUsersByRole, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUsersByRoleRow{}
	for rows.Next() {
		var i GetUsersByRoleRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.Role,
			&i.IsBlocked,
			&i.EmailVerified,
			&i.UserVerified,
			&i.GstNo,
			&i.About,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unblockUserByID = `-- name: UnblockUserByID :one
UPDATE users
SET is_blocked = false, updated_at = current_timestamp
WHERE id = $1
RETURNING id, name, email, phone, role, is_blocked, email_verified, user_verified, gst_no, about, created_at, updated_at
`

type UnblockUserByIDRow struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         sql.NullInt64  `json:"phone"`
	Role          string         `json:"role"`
	IsBlocked     bool           `json:"is_blocked"`
	EmailVerified bool           `json:"email_verified"`
	UserVerified  bool           `json:"user_verified"`
	GstNo         sql.NullString `json:"gst_no"`
	About         sql.NullString `json:"about"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (q *Queries) UnblockUserByID(ctx context.Context, id uuid.UUID) (UnblockUserByIDRow, error) {
	row := q.queryRow(ctx, q.unblockUserByIDStmt, unblockUserByID, id)
	var i UnblockUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.IsBlocked,
		&i.EmailVerified,
		&i.UserVerified,
		&i.GstNo,
		&i.About,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const verifySellerByID = `-- name: VerifySellerByID :one
update users
set user_verified = true, updated_at = current_timestamp
where id = $1
returning id, name, email, phone, role, is_blocked, email_verified, user_verified, created_at, updated_at
`

type VerifySellerByIDRow struct {
	ID            uuid.UUID     `json:"id"`
	Name          string        `json:"name"`
	Email         string        `json:"email"`
	Phone         sql.NullInt64 `json:"phone"`
	Role          string        `json:"role"`
	IsBlocked     bool          `json:"is_blocked"`
	EmailVerified bool          `json:"email_verified"`
	UserVerified  bool          `json:"user_verified"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

func (q *Queries) VerifySellerByID(ctx context.Context, id uuid.UUID) (VerifySellerByIDRow, error) {
	row := q.queryRow(ctx, q.verifySellerByIDStmt, verifySellerByID, id)
	var i VerifySellerByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.IsBlocked,
		&i.EmailVerified,
		&i.UserVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const verifySellerEmailByID = `-- name: VerifySellerEmailByID :one
UPDATE users
SET email_verified = true, updated_at = current_timestamp
WHERE id = $1
RETURNING id, name, email, phone, role, is_blocked, email_verified, user_verified, created_at, updated_at
`

type VerifySellerEmailByIDRow struct {
	ID            uuid.UUID     `json:"id"`
	Name          string        `json:"name"`
	Email         string        `json:"email"`
	Phone         sql.NullInt64 `json:"phone"`
	Role          string        `json:"role"`
	IsBlocked     bool          `json:"is_blocked"`
	EmailVerified bool          `json:"email_verified"`
	UserVerified  bool          `json:"user_verified"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

func (q *Queries) VerifySellerEmailByID(ctx context.Context, id uuid.UUID) (VerifySellerEmailByIDRow, error) {
	row := q.queryRow(ctx, q.verifySellerEmailByIDStmt, verifySellerEmailByID, id)
	var i VerifySellerEmailByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.IsBlocked,
		&i.EmailVerified,
		&i.UserVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const verifyUserByID = `-- name: VerifyUserByID :one
UPDATE users
SET email_verified = true, user_verified = true, updated_at = current_timestamp
WHERE id = $1
RETURNING id, name, email, phone, role, is_blocked, email_verified, user_verified, created_at, updated_at
`

type VerifyUserByIDRow struct {
	ID            uuid.UUID     `json:"id"`
	Name          string        `json:"name"`
	Email         string        `json:"email"`
	Phone         sql.NullInt64 `json:"phone"`
	Role          string        `json:"role"`
	IsBlocked     bool          `json:"is_blocked"`
	EmailVerified bool          `json:"email_verified"`
	UserVerified  bool          `json:"user_verified"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

func (q *Queries) VerifyUserByID(ctx context.Context, id uuid.UUID) (VerifyUserByIDRow, error) {
	row := q.queryRow(ctx, q.verifyUserByIDStmt, verifyUserByID, id)
	var i VerifyUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.IsBlocked,
		&i.EmailVerified,
		&i.UserVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
