// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.blockUserByIDStmt, err = db.PrepareContext(ctx, blockUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query BlockUserByID: %w", err)
	}
	if q.deleteProductByIDStmt, err = db.PrepareContext(ctx, deleteProductByID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProductByID: %w", err)
	}
	if q.deleteProductsBySellerIDStmt, err = db.PrepareContext(ctx, deleteProductsBySellerID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProductsBySellerID: %w", err)
	}
	if q.getAllProductsStmt, err = db.PrepareContext(ctx, getAllProducts); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllProducts: %w", err)
	}
	if q.getAllUsersStmt, err = db.PrepareContext(ctx, getAllUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllUsers: %w", err)
	}
	if q.getOTPByUserIDStmt, err = db.PrepareContext(ctx, getOTPByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query GetOTPByUserID: %w", err)
	}
	if q.getProductsByCategoryIDStmt, err = db.PrepareContext(ctx, getProductsByCategoryID); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductsByCategoryID: %w", err)
	}
	if q.getProductsBySellerIDStmt, err = db.PrepareContext(ctx, getProductsBySellerID); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductsBySellerID: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.getUsersByRoleStmt, err = db.PrepareContext(ctx, getUsersByRole); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsersByRole: %w", err)
	}
	if q.insertProductStmt, err = db.PrepareContext(ctx, insertProduct); err != nil {
		return nil, fmt.Errorf("error preparing query InsertProduct: %w", err)
	}
	if q.insertUserStmt, err = db.PrepareContext(ctx, insertUser); err != nil {
		return nil, fmt.Errorf("error preparing query InsertUser: %w", err)
	}
	if q.unblockUserByIDStmt, err = db.PrepareContext(ctx, unblockUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query UnblockUserByID: %w", err)
	}
	if q.updateProductByIDStmt, err = db.PrepareContext(ctx, updateProductByID); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProductByID: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.blockUserByIDStmt != nil {
		if cerr := q.blockUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing blockUserByIDStmt: %w", cerr)
		}
	}
	if q.deleteProductByIDStmt != nil {
		if cerr := q.deleteProductByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProductByIDStmt: %w", cerr)
		}
	}
	if q.deleteProductsBySellerIDStmt != nil {
		if cerr := q.deleteProductsBySellerIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProductsBySellerIDStmt: %w", cerr)
		}
	}
	if q.getAllProductsStmt != nil {
		if cerr := q.getAllProductsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllProductsStmt: %w", cerr)
		}
	}
	if q.getAllUsersStmt != nil {
		if cerr := q.getAllUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUsersStmt: %w", cerr)
		}
	}
	if q.getOTPByUserIDStmt != nil {
		if cerr := q.getOTPByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOTPByUserIDStmt: %w", cerr)
		}
	}
	if q.getProductsByCategoryIDStmt != nil {
		if cerr := q.getProductsByCategoryIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductsByCategoryIDStmt: %w", cerr)
		}
	}
	if q.getProductsBySellerIDStmt != nil {
		if cerr := q.getProductsBySellerIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductsBySellerIDStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.getUsersByRoleStmt != nil {
		if cerr := q.getUsersByRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersByRoleStmt: %w", cerr)
		}
	}
	if q.insertProductStmt != nil {
		if cerr := q.insertProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertProductStmt: %w", cerr)
		}
	}
	if q.insertUserStmt != nil {
		if cerr := q.insertUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertUserStmt: %w", cerr)
		}
	}
	if q.unblockUserByIDStmt != nil {
		if cerr := q.unblockUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing unblockUserByIDStmt: %w", cerr)
		}
	}
	if q.updateProductByIDStmt != nil {
		if cerr := q.updateProductByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProductByIDStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                           DBTX
	tx                           *sql.Tx
	blockUserByIDStmt            *sql.Stmt
	deleteProductByIDStmt        *sql.Stmt
	deleteProductsBySellerIDStmt *sql.Stmt
	getAllProductsStmt           *sql.Stmt
	getAllUsersStmt              *sql.Stmt
	getOTPByUserIDStmt           *sql.Stmt
	getProductsByCategoryIDStmt  *sql.Stmt
	getProductsBySellerIDStmt    *sql.Stmt
	getUserByEmailStmt           *sql.Stmt
	getUserByIdStmt              *sql.Stmt
	getUsersByRoleStmt           *sql.Stmt
	insertProductStmt            *sql.Stmt
	insertUserStmt               *sql.Stmt
	unblockUserByIDStmt          *sql.Stmt
	updateProductByIDStmt        *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                           tx,
		tx:                           tx,
		blockUserByIDStmt:            q.blockUserByIDStmt,
		deleteProductByIDStmt:        q.deleteProductByIDStmt,
		deleteProductsBySellerIDStmt: q.deleteProductsBySellerIDStmt,
		getAllProductsStmt:           q.getAllProductsStmt,
		getAllUsersStmt:              q.getAllUsersStmt,
		getOTPByUserIDStmt:           q.getOTPByUserIDStmt,
		getProductsByCategoryIDStmt:  q.getProductsByCategoryIDStmt,
		getProductsBySellerIDStmt:    q.getProductsBySellerIDStmt,
		getUserByEmailStmt:           q.getUserByEmailStmt,
		getUserByIdStmt:              q.getUserByIdStmt,
		getUsersByRoleStmt:           q.getUsersByRoleStmt,
		insertProductStmt:            q.insertProductStmt,
		insertUserStmt:               q.insertUserStmt,
		unblockUserByIDStmt:          q.unblockUserByIDStmt,
		updateProductByIDStmt:        q.updateProductByIDStmt,
	}
}
