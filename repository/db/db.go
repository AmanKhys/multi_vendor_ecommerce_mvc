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
	if q.addAndVerifyUserStmt, err = db.PrepareContext(ctx, addAndVerifyUser); err != nil {
		return nil, fmt.Errorf("error preparing query AddAndVerifyUser: %w", err)
	}
	if q.addCateogryStmt, err = db.PrepareContext(ctx, addCateogry); err != nil {
		return nil, fmt.Errorf("error preparing query AddCateogry: %w", err)
	}
	if q.addOTPStmt, err = db.PrepareContext(ctx, addOTP); err != nil {
		return nil, fmt.Errorf("error preparing query AddOTP: %w", err)
	}
	if q.addProductStmt, err = db.PrepareContext(ctx, addProduct); err != nil {
		return nil, fmt.Errorf("error preparing query AddProduct: %w", err)
	}
	if q.addProductToCategoryByCategoryNameStmt, err = db.PrepareContext(ctx, addProductToCategoryByCategoryName); err != nil {
		return nil, fmt.Errorf("error preparing query AddProductToCategoryByCategoryName: %w", err)
	}
	if q.addProductToCategoryByIDStmt, err = db.PrepareContext(ctx, addProductToCategoryByID); err != nil {
		return nil, fmt.Errorf("error preparing query AddProductToCategoryByID: %w", err)
	}
	if q.addSellerStmt, err = db.PrepareContext(ctx, addSeller); err != nil {
		return nil, fmt.Errorf("error preparing query AddSeller: %w", err)
	}
	if q.addSessionStmt, err = db.PrepareContext(ctx, addSession); err != nil {
		return nil, fmt.Errorf("error preparing query AddSession: %w", err)
	}
	if q.addUserStmt, err = db.PrepareContext(ctx, addUser); err != nil {
		return nil, fmt.Errorf("error preparing query AddUser: %w", err)
	}
	if q.blockUserByIDStmt, err = db.PrepareContext(ctx, blockUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query BlockUserByID: %w", err)
	}
	if q.deleteCategoryByNameStmt, err = db.PrepareContext(ctx, deleteCategoryByName); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCategoryByName: %w", err)
	}
	if q.deleteOTPByEmailStmt, err = db.PrepareContext(ctx, deleteOTPByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteOTPByEmail: %w", err)
	}
	if q.deleteProductByIDStmt, err = db.PrepareContext(ctx, deleteProductByID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProductByID: %w", err)
	}
	if q.deleteProductsBySellerIDStmt, err = db.PrepareContext(ctx, deleteProductsBySellerID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProductsBySellerID: %w", err)
	}
	if q.deleteSessionByIDStmt, err = db.PrepareContext(ctx, deleteSessionByID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSessionByID: %w", err)
	}
	if q.deleteSessionsByuserIDStmt, err = db.PrepareContext(ctx, deleteSessionsByuserID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSessionsByuserID: %w", err)
	}
	if q.editCategoryNameByIDStmt, err = db.PrepareContext(ctx, editCategoryNameByID); err != nil {
		return nil, fmt.Errorf("error preparing query EditCategoryNameByID: %w", err)
	}
	if q.editProductByIDStmt, err = db.PrepareContext(ctx, editProductByID); err != nil {
		return nil, fmt.Errorf("error preparing query EditProductByID: %w", err)
	}
	if q.getAllCategoriesStmt, err = db.PrepareContext(ctx, getAllCategories); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllCategories: %w", err)
	}
	if q.getAllCategoriesForAdminStmt, err = db.PrepareContext(ctx, getAllCategoriesForAdmin); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllCategoriesForAdmin: %w", err)
	}
	if q.getAllProductsStmt, err = db.PrepareContext(ctx, getAllProducts); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllProducts: %w", err)
	}
	if q.getAllProductsForAdminStmt, err = db.PrepareContext(ctx, getAllProductsForAdmin); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllProductsForAdmin: %w", err)
	}
	if q.getAllSessionsByUserIDStmt, err = db.PrepareContext(ctx, getAllSessionsByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllSessionsByUserID: %w", err)
	}
	if q.getAllUsersStmt, err = db.PrepareContext(ctx, getAllUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllUsers: %w", err)
	}
	if q.getAllUsersByRoleStmt, err = db.PrepareContext(ctx, getAllUsersByRole); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllUsersByRole: %w", err)
	}
	if q.getCategoryByIDStmt, err = db.PrepareContext(ctx, getCategoryByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetCategoryByID: %w", err)
	}
	if q.getCategoryByNameStmt, err = db.PrepareContext(ctx, getCategoryByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetCategoryByName: %w", err)
	}
	if q.getCurrentTimestampStmt, err = db.PrepareContext(ctx, getCurrentTimestamp); err != nil {
		return nil, fmt.Errorf("error preparing query GetCurrentTimestamp: %w", err)
	}
	if q.getProductAndCategoryNameByIDStmt, err = db.PrepareContext(ctx, getProductAndCategoryNameByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductAndCategoryNameByID: %w", err)
	}
	if q.getProductByIDStmt, err = db.PrepareContext(ctx, getProductByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductByID: %w", err)
	}
	if q.getProductsByCategoryNameStmt, err = db.PrepareContext(ctx, getProductsByCategoryName); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductsByCategoryName: %w", err)
	}
	if q.getProductsBySellerIDStmt, err = db.PrepareContext(ctx, getProductsBySellerID); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductsBySellerID: %w", err)
	}
	if q.getSellerByProductIDStmt, err = db.PrepareContext(ctx, getSellerByProductID); err != nil {
		return nil, fmt.Errorf("error preparing query GetSellerByProductID: %w", err)
	}
	if q.getSessionDetailsByIDStmt, err = db.PrepareContext(ctx, getSessionDetailsByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetSessionDetailsByID: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.getUserBySessionIDStmt, err = db.PrepareContext(ctx, getUserBySessionID); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserBySessionID: %w", err)
	}
	if q.getUserWithPasswordByEmailStmt, err = db.PrepareContext(ctx, getUserWithPasswordByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserWithPasswordByEmail: %w", err)
	}
	if q.getUsersByRoleStmt, err = db.PrepareContext(ctx, getUsersByRole); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsersByRole: %w", err)
	}
	if q.getValidOTPByUserIDStmt, err = db.PrepareContext(ctx, getValidOTPByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query GetValidOTPByUserID: %w", err)
	}
	if q.unblockUserByIDStmt, err = db.PrepareContext(ctx, unblockUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query UnblockUserByID: %w", err)
	}
	if q.verifySellerByIDStmt, err = db.PrepareContext(ctx, verifySellerByID); err != nil {
		return nil, fmt.Errorf("error preparing query VerifySellerByID: %w", err)
	}
	if q.verifySellerEmailByIDStmt, err = db.PrepareContext(ctx, verifySellerEmailByID); err != nil {
		return nil, fmt.Errorf("error preparing query VerifySellerEmailByID: %w", err)
	}
	if q.verifyUserByIDStmt, err = db.PrepareContext(ctx, verifyUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query VerifyUserByID: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addAndVerifyUserStmt != nil {
		if cerr := q.addAndVerifyUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addAndVerifyUserStmt: %w", cerr)
		}
	}
	if q.addCateogryStmt != nil {
		if cerr := q.addCateogryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addCateogryStmt: %w", cerr)
		}
	}
	if q.addOTPStmt != nil {
		if cerr := q.addOTPStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addOTPStmt: %w", cerr)
		}
	}
	if q.addProductStmt != nil {
		if cerr := q.addProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addProductStmt: %w", cerr)
		}
	}
	if q.addProductToCategoryByCategoryNameStmt != nil {
		if cerr := q.addProductToCategoryByCategoryNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addProductToCategoryByCategoryNameStmt: %w", cerr)
		}
	}
	if q.addProductToCategoryByIDStmt != nil {
		if cerr := q.addProductToCategoryByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addProductToCategoryByIDStmt: %w", cerr)
		}
	}
	if q.addSellerStmt != nil {
		if cerr := q.addSellerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addSellerStmt: %w", cerr)
		}
	}
	if q.addSessionStmt != nil {
		if cerr := q.addSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addSessionStmt: %w", cerr)
		}
	}
	if q.addUserStmt != nil {
		if cerr := q.addUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUserStmt: %w", cerr)
		}
	}
	if q.blockUserByIDStmt != nil {
		if cerr := q.blockUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing blockUserByIDStmt: %w", cerr)
		}
	}
	if q.deleteCategoryByNameStmt != nil {
		if cerr := q.deleteCategoryByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCategoryByNameStmt: %w", cerr)
		}
	}
	if q.deleteOTPByEmailStmt != nil {
		if cerr := q.deleteOTPByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteOTPByEmailStmt: %w", cerr)
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
	if q.deleteSessionByIDStmt != nil {
		if cerr := q.deleteSessionByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSessionByIDStmt: %w", cerr)
		}
	}
	if q.deleteSessionsByuserIDStmt != nil {
		if cerr := q.deleteSessionsByuserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSessionsByuserIDStmt: %w", cerr)
		}
	}
	if q.editCategoryNameByIDStmt != nil {
		if cerr := q.editCategoryNameByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing editCategoryNameByIDStmt: %w", cerr)
		}
	}
	if q.editProductByIDStmt != nil {
		if cerr := q.editProductByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing editProductByIDStmt: %w", cerr)
		}
	}
	if q.getAllCategoriesStmt != nil {
		if cerr := q.getAllCategoriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllCategoriesStmt: %w", cerr)
		}
	}
	if q.getAllCategoriesForAdminStmt != nil {
		if cerr := q.getAllCategoriesForAdminStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllCategoriesForAdminStmt: %w", cerr)
		}
	}
	if q.getAllProductsStmt != nil {
		if cerr := q.getAllProductsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllProductsStmt: %w", cerr)
		}
	}
	if q.getAllProductsForAdminStmt != nil {
		if cerr := q.getAllProductsForAdminStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllProductsForAdminStmt: %w", cerr)
		}
	}
	if q.getAllSessionsByUserIDStmt != nil {
		if cerr := q.getAllSessionsByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllSessionsByUserIDStmt: %w", cerr)
		}
	}
	if q.getAllUsersStmt != nil {
		if cerr := q.getAllUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUsersStmt: %w", cerr)
		}
	}
	if q.getAllUsersByRoleStmt != nil {
		if cerr := q.getAllUsersByRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUsersByRoleStmt: %w", cerr)
		}
	}
	if q.getCategoryByIDStmt != nil {
		if cerr := q.getCategoryByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCategoryByIDStmt: %w", cerr)
		}
	}
	if q.getCategoryByNameStmt != nil {
		if cerr := q.getCategoryByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCategoryByNameStmt: %w", cerr)
		}
	}
	if q.getCurrentTimestampStmt != nil {
		if cerr := q.getCurrentTimestampStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCurrentTimestampStmt: %w", cerr)
		}
	}
	if q.getProductAndCategoryNameByIDStmt != nil {
		if cerr := q.getProductAndCategoryNameByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductAndCategoryNameByIDStmt: %w", cerr)
		}
	}
	if q.getProductByIDStmt != nil {
		if cerr := q.getProductByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductByIDStmt: %w", cerr)
		}
	}
	if q.getProductsByCategoryNameStmt != nil {
		if cerr := q.getProductsByCategoryNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductsByCategoryNameStmt: %w", cerr)
		}
	}
	if q.getProductsBySellerIDStmt != nil {
		if cerr := q.getProductsBySellerIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductsBySellerIDStmt: %w", cerr)
		}
	}
	if q.getSellerByProductIDStmt != nil {
		if cerr := q.getSellerByProductIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSellerByProductIDStmt: %w", cerr)
		}
	}
	if q.getSessionDetailsByIDStmt != nil {
		if cerr := q.getSessionDetailsByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSessionDetailsByIDStmt: %w", cerr)
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
	if q.getUserBySessionIDStmt != nil {
		if cerr := q.getUserBySessionIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserBySessionIDStmt: %w", cerr)
		}
	}
	if q.getUserWithPasswordByEmailStmt != nil {
		if cerr := q.getUserWithPasswordByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserWithPasswordByEmailStmt: %w", cerr)
		}
	}
	if q.getUsersByRoleStmt != nil {
		if cerr := q.getUsersByRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersByRoleStmt: %w", cerr)
		}
	}
	if q.getValidOTPByUserIDStmt != nil {
		if cerr := q.getValidOTPByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getValidOTPByUserIDStmt: %w", cerr)
		}
	}
	if q.unblockUserByIDStmt != nil {
		if cerr := q.unblockUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing unblockUserByIDStmt: %w", cerr)
		}
	}
	if q.verifySellerByIDStmt != nil {
		if cerr := q.verifySellerByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing verifySellerByIDStmt: %w", cerr)
		}
	}
	if q.verifySellerEmailByIDStmt != nil {
		if cerr := q.verifySellerEmailByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing verifySellerEmailByIDStmt: %w", cerr)
		}
	}
	if q.verifyUserByIDStmt != nil {
		if cerr := q.verifyUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing verifyUserByIDStmt: %w", cerr)
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
	db                                     DBTX
	tx                                     *sql.Tx
	addAndVerifyUserStmt                   *sql.Stmt
	addCateogryStmt                        *sql.Stmt
	addOTPStmt                             *sql.Stmt
	addProductStmt                         *sql.Stmt
	addProductToCategoryByCategoryNameStmt *sql.Stmt
	addProductToCategoryByIDStmt           *sql.Stmt
	addSellerStmt                          *sql.Stmt
	addSessionStmt                         *sql.Stmt
	addUserStmt                            *sql.Stmt
	blockUserByIDStmt                      *sql.Stmt
	deleteCategoryByNameStmt               *sql.Stmt
	deleteOTPByEmailStmt                   *sql.Stmt
	deleteProductByIDStmt                  *sql.Stmt
	deleteProductsBySellerIDStmt           *sql.Stmt
	deleteSessionByIDStmt                  *sql.Stmt
	deleteSessionsByuserIDStmt             *sql.Stmt
	editCategoryNameByIDStmt               *sql.Stmt
	editProductByIDStmt                    *sql.Stmt
	getAllCategoriesStmt                   *sql.Stmt
	getAllCategoriesForAdminStmt           *sql.Stmt
	getAllProductsStmt                     *sql.Stmt
	getAllProductsForAdminStmt             *sql.Stmt
	getAllSessionsByUserIDStmt             *sql.Stmt
	getAllUsersStmt                        *sql.Stmt
	getAllUsersByRoleStmt                  *sql.Stmt
	getCategoryByIDStmt                    *sql.Stmt
	getCategoryByNameStmt                  *sql.Stmt
	getCurrentTimestampStmt                *sql.Stmt
	getProductAndCategoryNameByIDStmt      *sql.Stmt
	getProductByIDStmt                     *sql.Stmt
	getProductsByCategoryNameStmt          *sql.Stmt
	getProductsBySellerIDStmt              *sql.Stmt
	getSellerByProductIDStmt               *sql.Stmt
	getSessionDetailsByIDStmt              *sql.Stmt
	getUserByEmailStmt                     *sql.Stmt
	getUserByIdStmt                        *sql.Stmt
	getUserBySessionIDStmt                 *sql.Stmt
	getUserWithPasswordByEmailStmt         *sql.Stmt
	getUsersByRoleStmt                     *sql.Stmt
	getValidOTPByUserIDStmt                *sql.Stmt
	unblockUserByIDStmt                    *sql.Stmt
	verifySellerByIDStmt                   *sql.Stmt
	verifySellerEmailByIDStmt              *sql.Stmt
	verifyUserByIDStmt                     *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                     tx,
		tx:                                     tx,
		addAndVerifyUserStmt:                   q.addAndVerifyUserStmt,
		addCateogryStmt:                        q.addCateogryStmt,
		addOTPStmt:                             q.addOTPStmt,
		addProductStmt:                         q.addProductStmt,
		addProductToCategoryByCategoryNameStmt: q.addProductToCategoryByCategoryNameStmt,
		addProductToCategoryByIDStmt:           q.addProductToCategoryByIDStmt,
		addSellerStmt:                          q.addSellerStmt,
		addSessionStmt:                         q.addSessionStmt,
		addUserStmt:                            q.addUserStmt,
		blockUserByIDStmt:                      q.blockUserByIDStmt,
		deleteCategoryByNameStmt:               q.deleteCategoryByNameStmt,
		deleteOTPByEmailStmt:                   q.deleteOTPByEmailStmt,
		deleteProductByIDStmt:                  q.deleteProductByIDStmt,
		deleteProductsBySellerIDStmt:           q.deleteProductsBySellerIDStmt,
		deleteSessionByIDStmt:                  q.deleteSessionByIDStmt,
		deleteSessionsByuserIDStmt:             q.deleteSessionsByuserIDStmt,
		editCategoryNameByIDStmt:               q.editCategoryNameByIDStmt,
		editProductByIDStmt:                    q.editProductByIDStmt,
		getAllCategoriesStmt:                   q.getAllCategoriesStmt,
		getAllCategoriesForAdminStmt:           q.getAllCategoriesForAdminStmt,
		getAllProductsStmt:                     q.getAllProductsStmt,
		getAllProductsForAdminStmt:             q.getAllProductsForAdminStmt,
		getAllSessionsByUserIDStmt:             q.getAllSessionsByUserIDStmt,
		getAllUsersStmt:                        q.getAllUsersStmt,
		getAllUsersByRoleStmt:                  q.getAllUsersByRoleStmt,
		getCategoryByIDStmt:                    q.getCategoryByIDStmt,
		getCategoryByNameStmt:                  q.getCategoryByNameStmt,
		getCurrentTimestampStmt:                q.getCurrentTimestampStmt,
		getProductAndCategoryNameByIDStmt:      q.getProductAndCategoryNameByIDStmt,
		getProductByIDStmt:                     q.getProductByIDStmt,
		getProductsByCategoryNameStmt:          q.getProductsByCategoryNameStmt,
		getProductsBySellerIDStmt:              q.getProductsBySellerIDStmt,
		getSellerByProductIDStmt:               q.getSellerByProductIDStmt,
		getSessionDetailsByIDStmt:              q.getSessionDetailsByIDStmt,
		getUserByEmailStmt:                     q.getUserByEmailStmt,
		getUserByIdStmt:                        q.getUserByIdStmt,
		getUserBySessionIDStmt:                 q.getUserBySessionIDStmt,
		getUserWithPasswordByEmailStmt:         q.getUserWithPasswordByEmailStmt,
		getUsersByRoleStmt:                     q.getUsersByRoleStmt,
		getValidOTPByUserIDStmt:                q.getValidOTPByUserIDStmt,
		unblockUserByIDStmt:                    q.unblockUserByIDStmt,
		verifySellerByIDStmt:                   q.verifySellerByIDStmt,
		verifySellerEmailByIDStmt:              q.verifySellerEmailByIDStmt,
		verifyUserByIDStmt:                     q.verifyUserByIDStmt,
	}
}
