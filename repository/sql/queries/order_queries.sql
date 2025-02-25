-- name: GetAllOrderForAdmin :many
select * from order_items
order by created_at desc;

-- name: GetOrdersByUserID :many
select * from orders
where user_id = $1
order by created_at desc;

-- name: GetOrderByID :one
select * from orders
where id = $1;

-- name: GetOrderItemsBySellerID :many
select oi.* from order_items oi
inner join products p
on oi.product_id = p.id
where p.seller_id = $1
order by oi.created_at desc;

-- name: GetSellerIDFromOrderItemID :one
select p.seller_id from order_items oi
inner join products p
on oi.product_id = p.id
where oi.id = $1;


-- name: AddOrder :one
insert into orders
(user_id)
values (@user_id)
returning *;

-- name: DeleteOrderByID :exec
delete from orders
where id = $1;

-- name: AddOrderITem :one
insert into order_items
(order_id, product_id,price, quantity, total_amount)
values
($1, $2, $3, $4, $5)
returning *;

-- name: GetOrderItemsByUserID :many
select oi.* from order_items oi
inner join orders o
on oi.order_id = o.id
inner join users u
on o.user_id = u.id
where u.id = $1
order by oi.created_at desc;

-- name: GetOrderItemByID :one
select * from order_items
where id = $1;

-- name: GetUserIDFromOrderItemID :one
select u.id from order_items oi
inner join orders o
on oi.order_id = o.id
inner join users u
on o.user_id = u.id
where oi.id = $1;

-- name: GetOrderItemsByOrderID :many
select oi.*, p.name as product_name
from order_items oi
inner join products p
on oi.product_id = p.id
where oi.order_id = $1;

-- name: AddShippingAddress :one
insert into shipping_address
(order_id, house_name, street_name, town, district, state, pincode)
values
($1, $2, $3, $4, $5,$6, $7)
returning id, house_name, street_name, town, district, state, pincode;

-- name: ChangeOrderItemStatusByID :one
update order_items oi
set status =  $2
where id = $1
returning *;

-- name: EditOrderItemStatusByID :one
update order_items
set status = $2, updated_at = current_timestamp
where id = $1
returning *;

-- name: CancelOrderByID :exec
update order_items
set status = 'cancelled', updated_at = current_timestamp
where order_id = $1;
