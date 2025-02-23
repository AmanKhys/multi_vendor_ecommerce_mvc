// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: order_queries.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addOrder = `-- name: AddOrder :one
insert into orders
(user_id)
values ($1)
returning id, user_id, created_at, updated_at
`

func (q *Queries) AddOrder(ctx context.Context, userID uuid.UUID) (Order, error) {
	row := q.queryRow(ctx, q.addOrderStmt, addOrder, userID)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const addOrderITem = `-- name: AddOrderITem :one
insert into order_items
(order_id, product_id,price, quantity, total_amount)
values
($1, $2, $3, $4, $5)
returning id, order_id, product_id, price, quantity, total_amount, status, created_at, updated_at
`

type AddOrderITemParams struct {
	OrderID     uuid.UUID `json:"order_id"`
	ProductID   uuid.UUID `json:"product_id"`
	Price       float64   `json:"price"`
	Quantity    int32     `json:"quantity"`
	TotalAmount float64   `json:"total_amount"`
}

func (q *Queries) AddOrderITem(ctx context.Context, arg AddOrderITemParams) (OrderItem, error) {
	row := q.queryRow(ctx, q.addOrderITemStmt, addOrderITem,
		arg.OrderID,
		arg.ProductID,
		arg.Price,
		arg.Quantity,
		arg.TotalAmount,
	)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ProductID,
		&i.Price,
		&i.Quantity,
		&i.TotalAmount,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const addPayment = `-- name: AddPayment :one
insert into payments
(order_id, method, status, total_amount)
values
($1, $2, $3, $4)
returning id, order_id, method, status, total_amount, transaction_id, created_at, updated_at
`

type AddPaymentParams struct {
	OrderID     uuid.UUID `json:"order_id"`
	Method      string    `json:"method"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
}

func (q *Queries) AddPayment(ctx context.Context, arg AddPaymentParams) (Payment, error) {
	row := q.queryRow(ctx, q.addPaymentStmt, addPayment,
		arg.OrderID,
		arg.Method,
		arg.Status,
		arg.TotalAmount,
	)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.Method,
		&i.Status,
		&i.TotalAmount,
		&i.TransactionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const addShippingAddress = `-- name: AddShippingAddress :one
insert into shipping_address
(order_id, house_name, street_name, town, district, state, pincode)
values
($1, $2, $3, $4, $5,$6, $7)
returning id, house_name, street_name, town, district, state, pincode
`

type AddShippingAddressParams struct {
	OrderID    uuid.UUID `json:"order_id"`
	HouseName  string    `json:"house_name"`
	StreetName string    `json:"street_name"`
	Town       string    `json:"town"`
	District   string    `json:"district"`
	State      string    `json:"state"`
	Pincode    int32     `json:"pincode"`
}

type AddShippingAddressRow struct {
	ID         uuid.UUID `json:"id"`
	HouseName  string    `json:"house_name"`
	StreetName string    `json:"street_name"`
	Town       string    `json:"town"`
	District   string    `json:"district"`
	State      string    `json:"state"`
	Pincode    int32     `json:"pincode"`
}

func (q *Queries) AddShippingAddress(ctx context.Context, arg AddShippingAddressParams) (AddShippingAddressRow, error) {
	row := q.queryRow(ctx, q.addShippingAddressStmt, addShippingAddress,
		arg.OrderID,
		arg.HouseName,
		arg.StreetName,
		arg.Town,
		arg.District,
		arg.State,
		arg.Pincode,
	)
	var i AddShippingAddressRow
	err := row.Scan(
		&i.ID,
		&i.HouseName,
		&i.StreetName,
		&i.Town,
		&i.District,
		&i.State,
		&i.Pincode,
	)
	return i, err
}

const cancelOrderByID = `-- name: CancelOrderByID :exec
update order_items
set status = "cancelled", updated_at = current_timestamp
where order_id = $1
`

func (q *Queries) CancelOrderByID(ctx context.Context, orderID uuid.UUID) error {
	_, err := q.exec(ctx, q.cancelOrderByIDStmt, cancelOrderByID, orderID)
	return err
}

const cancelPaymentByOrderID = `-- name: CancelPaymentByOrderID :exec
update payments
set status = "returned", total_amount = 0, updated_at = current_timestamp
where order_id = $1
`

func (q *Queries) CancelPaymentByOrderID(ctx context.Context, orderID uuid.UUID) error {
	_, err := q.exec(ctx, q.cancelPaymentByOrderIDStmt, cancelPaymentByOrderID, orderID)
	return err
}

const decPaymentAmountByOrderItemID = `-- name: DecPaymentAmountByOrderItemID :one
WITH cte AS (
  SELECT oi.order_id, oi.total_amount
  FROM order_items oi
  WHERE oi.id = $1
)
UPDATE payments
SET total_amount = payments.total_amount - cte.total_amount
FROM cte
WHERE payments.order_id = cte.order_id
RETURNING payments.id, payments.order_id, payments.method, payments.status, payments.total_amount, payments.transaction_id, payments.created_at, payments.updated_at
`

func (q *Queries) DecPaymentAmountByOrderItemID(ctx context.Context, id uuid.UUID) (Payment, error) {
	row := q.queryRow(ctx, q.decPaymentAmountByOrderItemIDStmt, decPaymentAmountByOrderItemID, id)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.Method,
		&i.Status,
		&i.TotalAmount,
		&i.TransactionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOrderByID = `-- name: DeleteOrderByID :exec
delete from orders
where id = $1
`

func (q *Queries) DeleteOrderByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteOrderByIDStmt, deleteOrderByID, id)
	return err
}

const editOrderItemStatusByID = `-- name: EditOrderItemStatusByID :one
update order_items
set status = $2, updated_at = current_timestamp
where id = $1
returning id, order_id, product_id, price, quantity, total_amount, status, created_at, updated_at
`

type EditOrderItemStatusByIDParams struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

func (q *Queries) EditOrderItemStatusByID(ctx context.Context, arg EditOrderItemStatusByIDParams) (OrderItem, error) {
	row := q.queryRow(ctx, q.editOrderItemStatusByIDStmt, editOrderItemStatusByID, arg.ID, arg.Status)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ProductID,
		&i.Price,
		&i.Quantity,
		&i.TotalAmount,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const editPaymentStatusByID = `-- name: EditPaymentStatusByID :one
update payments
set status = $2, updated_at = current_timestamp
where id = $1
returning id, order_id, method, status, total_amount, transaction_id, created_at, updated_at
`

type EditPaymentStatusByIDParams struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

func (q *Queries) EditPaymentStatusByID(ctx context.Context, arg EditPaymentStatusByIDParams) (Payment, error) {
	row := q.queryRow(ctx, q.editPaymentStatusByIDStmt, editPaymentStatusByID, arg.ID, arg.Status)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.Method,
		&i.Status,
		&i.TotalAmount,
		&i.TransactionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const editPaymentStatusByOrderID = `-- name: EditPaymentStatusByOrderID :one
update payments
set status = $2, updated_at = current_timestamp
where order_id = $1
returning id, order_id, method, status, total_amount, transaction_id, created_at, updated_at
`

type EditPaymentStatusByOrderIDParams struct {
	OrderID uuid.UUID `json:"order_id"`
	Status  string    `json:"status"`
}

func (q *Queries) EditPaymentStatusByOrderID(ctx context.Context, arg EditPaymentStatusByOrderIDParams) (Payment, error) {
	row := q.queryRow(ctx, q.editPaymentStatusByOrderIDStmt, editPaymentStatusByOrderID, arg.OrderID, arg.Status)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.Method,
		&i.Status,
		&i.TotalAmount,
		&i.TransactionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrderByID = `-- name: GetOrderByID :one
select id, user_id, created_at, updated_at from orders
where id = $1
`

func (q *Queries) GetOrderByID(ctx context.Context, id uuid.UUID) (Order, error) {
	row := q.queryRow(ctx, q.getOrderByIDStmt, getOrderByID, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrderItemByID = `-- name: GetOrderItemByID :one
select id, order_id, product_id, price, quantity, total_amount, status, created_at, updated_at from order_items
where id = $1
`

func (q *Queries) GetOrderItemByID(ctx context.Context, id uuid.UUID) (OrderItem, error) {
	row := q.queryRow(ctx, q.getOrderItemByIDStmt, getOrderItemByID, id)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ProductID,
		&i.Price,
		&i.Quantity,
		&i.TotalAmount,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrderItemsByOrderID = `-- name: GetOrderItemsByOrderID :many
select oi.id, oi.order_id, oi.product_id, oi.price, oi.quantity, oi.total_amount, oi.status, oi.created_at, oi.updated_at, p.name as product_name
from order_items oi
inner join products p
on oi.product_id = p.id
where oi.order_id = $1
`

type GetOrderItemsByOrderIDRow struct {
	ID          uuid.UUID `json:"id"`
	OrderID     uuid.UUID `json:"order_id"`
	ProductID   uuid.UUID `json:"product_id"`
	Price       float64   `json:"price"`
	Quantity    int32     `json:"quantity"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ProductName string    `json:"product_name"`
}

func (q *Queries) GetOrderItemsByOrderID(ctx context.Context, orderID uuid.UUID) ([]GetOrderItemsByOrderIDRow, error) {
	rows, err := q.query(ctx, q.getOrderItemsByOrderIDStmt, getOrderItemsByOrderID, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrderItemsByOrderIDRow{}
	for rows.Next() {
		var i GetOrderItemsByOrderIDRow
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.ProductID,
			&i.Price,
			&i.Quantity,
			&i.TotalAmount,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProductName,
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

const getOrderItemsByUserID = `-- name: GetOrderItemsByUserID :many
select oi.id, oi.order_id, oi.product_id, oi.price, oi.quantity, oi.total_amount, oi.status, oi.created_at, oi.updated_at from order_items oi
inner join orders o
on oi.order_id = o.id
inner join users u
on o.user_id = u.id
where u.id = $1
`

func (q *Queries) GetOrderItemsByUserID(ctx context.Context, id uuid.UUID) ([]OrderItem, error) {
	rows, err := q.query(ctx, q.getOrderItemsByUserIDStmt, getOrderItemsByUserID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderItem{}
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.ProductID,
			&i.Price,
			&i.Quantity,
			&i.TotalAmount,
			&i.Status,
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

const getOrdersByUserID = `-- name: GetOrdersByUserID :many
select id, user_id, created_at, updated_at from orders
where user_id = $1
`

func (q *Queries) GetOrdersByUserID(ctx context.Context, userID uuid.UUID) ([]Order, error) {
	rows, err := q.query(ctx, q.getOrdersByUserIDStmt, getOrdersByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
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

const getUserIDFromOrderItemID = `-- name: GetUserIDFromOrderItemID :one
select u.id from order_items oi
inner join orders o
on oi.order_id = o.id
inner join users u
on o.user_id = u.id
where oi.id = $1
`

func (q *Queries) GetUserIDFromOrderItemID(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.getUserIDFromOrderItemIDStmt, getUserIDFromOrderItemID, id)
	err := row.Scan(&id)
	return id, err
}
