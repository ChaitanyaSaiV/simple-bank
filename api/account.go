package api

import (
	"fmt"
	"net/http"

	db "github.com/ChaitanyaSaiV/simple-bank/internal/db/methods"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)
}

type getAccountParams struct {
	ID int64 `uri:"id", binding:"required, min=1"`
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var arg getAccountParams
	if err := ctx.ShouldBindUri(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	fmt.Println(arg.ID)
	account, err := server.store.GetAccount(ctx, arg.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listaccountParams struct {
	PageID   int32 `form:"page_id", binding:"required, min=1"`
	PageSize int32 `form:"page_size", binding:"required, min=1, max=10"`
}

func (s *Server) ListAccounts(ctx *gin.Context) {
	var req listaccountParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	fmt.Println(arg)
	Accounts, err := s.store.ListAccounts(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, Accounts)
}
