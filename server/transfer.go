package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/capungkoneng/gomcommerce/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

/****************************************************************/
/************************** CREATE TRANSFER ************************/
/****************************************************************/
type transferRequest struct {
	FromAkun int64  `json:"from_akun" binding:"required,min=1"`
	ToAkun   int64  `json:"to_akun" binding:"required,min=1 "`
	Amount   int64  `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency" binding:"required,oneof=USD IDN"`
}

// Create akun one
func (server *Server) CreateTransfer(ctx *gin.Context) {
	var req transferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAkun(ctx, req.FromAkun, req.Currency) {
		return
	}
	if !server.validAkun(ctx, req.ToAkun, req.Currency) {
		return
	}
	arg := db.TransferTxParams{
		FromAkun: req.FromAkun,
		ToAkun:   req.ToAkun,
		Amount:   req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAkun(ctx *gin.Context, akunId int64, currency string) bool {
	akun, err := server.store.GetAuthor(ctx, akunId)
	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}
	if akun.Currency != currency {
		err := fmt.Errorf("akun [%d] currency mismatch: %s vs %s", akun.ID, akun.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return true
}
