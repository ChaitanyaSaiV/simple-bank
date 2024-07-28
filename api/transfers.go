package api

import (
	"database/sql"
	"net/http"

	db "github.com/ChaitanyaSaiV/simple-bank/internal/db/methods"
	"github.com/gin-gonic/gin"
)

type CreateTransferAccountParams struct {
	FromAccountID int64  `json:"from_account_id", binding:"required, min=1"`
	ToAccountID   int64  `json:"to_account_id", binding:"required, min=1`
	Amount        int64  `json:"amount", binding:"required, min=1`
	Currency      string `json: "currency", binding:"required, oneof=USD EUR CAD"`
}

func (s *Server) CreateTransfer(ctx *gin.Context) {
	var req CreateTransferAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	if !s.ValidateAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}
	if !s.ValidateAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	res, err := s.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *Server) ValidateAccount(ctx *gin.Context, id int64, currency string) bool {
	account, err := s.store.GetAccount(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}
	if currency == account.Currency {
		return true
	}
	return false
}
