package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/ChaitanyaSaiV/simple-bank/token"

	db "github.com/ChaitanyaSaiV/simple-bank/internal/db/methods"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
)

// PostgreSQL error codes for specific constraint violations
const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

// Request payload for creating an account
type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

// CreateAccount handles the creation of a new account
func (server *Server) CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(autorizationPayloadKey).(*token.Payload)

	arg := db.CreateAccountParams{
		Owner:    authPayload.UserName,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		errCode := ErrorCode(err)
		if errCode == ForeignKeyViolation || errCode == UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// Request parameters for getting an account by ID
type getAccountParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// GetAccount handles fetching an account by ID
func (server *Server) GetAccount(ctx *gin.Context) {
	var arg getAccountParams
	if err := ctx.ShouldBindUri(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, arg.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(autorizationPayloadKey).(*token.Payload)

	if account.Owner != authPayload.UserName {
		err := errors.New("account does not belongs to the user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// Request parameters for listing accounts with pagination
type listAccountParams struct {
	Owner    string
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

// ListAccounts handles listing accounts with pagination
func (server *Server) ListAccounts(ctx *gin.Context) {
	var req listAccountParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(autorizationPayloadKey).(*token.Payload)

	arg := db.ListAccountsParams{
		Owner:  authPayload.UserName,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

// ErrorCode extracts the PostgreSQL error code from an error
func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		log.Println("Error Code")
		return pgErr.Code
	}
	return ""
}
