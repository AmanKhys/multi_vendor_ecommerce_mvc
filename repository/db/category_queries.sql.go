// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: category_queries.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const addCateogry = `-- name: AddCateogry :one
insert into categories
(name) values ($1)
returning id, name, is_deleted, created_at, updated_at
`

func (q *Queries) AddCateogry(ctx context.Context, name string) (Category, error) {
	row := q.queryRow(ctx, q.addCateogryStmt, addCateogry, name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCategoryByName = `-- name: DeleteCategoryByName :one
update categories
set is_deleted = true, updated_at = current_timestamp
where name = $1
returning id, name, is_deleted, created_at, updated_at
`

func (q *Queries) DeleteCategoryByName(ctx context.Context, name string) (Category, error) {
	row := q.queryRow(ctx, q.deleteCategoryByNameStmt, deleteCategoryByName, name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const editCategoryNameByID = `-- name: EditCategoryNameByID :one
update categories
set name = $2, updated_at = current_timestamp
where id = $1
returning id, name, is_deleted, created_at, updated_at
`

type EditCategoryNameByIDParams struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (q *Queries) EditCategoryNameByID(ctx context.Context, arg EditCategoryNameByIDParams) (Category, error) {
	row := q.queryRow(ctx, q.editCategoryNameByIDStmt, editCategoryNameByID, arg.ID, arg.Name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllCategories = `-- name: GetAllCategories :many
select id, name, is_deleted, created_at, updated_at from categories
where is_deleted = false
`

func (q *Queries) GetAllCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.query(ctx, q.getAllCategoriesStmt, getAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.IsDeleted,
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

const getAllCategoriesForAdmin = `-- name: GetAllCategoriesForAdmin :many
select id, name, is_deleted, created_at, updated_at from categories
`

func (q *Queries) GetAllCategoriesForAdmin(ctx context.Context) ([]Category, error) {
	rows, err := q.query(ctx, q.getAllCategoriesForAdminStmt, getAllCategoriesForAdmin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.IsDeleted,
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

const getCategoryByID = `-- name: GetCategoryByID :one
select id, name, is_deleted, created_at, updated_at from categories
where id = $1 and is_deleted = false
`

func (q *Queries) GetCategoryByID(ctx context.Context, id uuid.UUID) (Category, error) {
	row := q.queryRow(ctx, q.getCategoryByIDStmt, getCategoryByID, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCategoryByName = `-- name: GetCategoryByName :one
select id, name, is_deleted, created_at, updated_at from categories
where name = $1 and is_deleted = false
`

func (q *Queries) GetCategoryByName(ctx context.Context, name string) (Category, error) {
	row := q.queryRow(ctx, q.getCategoryByNameStmt, getCategoryByName, name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
