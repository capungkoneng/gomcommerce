package api

import (
	"database/sql"
	"net/http"

	db "github.com/capungkoneng/gomcommerce.git/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAkunRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required, oneof=USD IDN"`
}

// Create akun
func (server *Server) CreateAkun(ctx *gin.Context) {
	var req createAkunRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAuthorParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	akun, err := server.store.CreateAuthor(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, akun)
}

// Error func
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

type getAkunRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

//Get akun
func (server *Server) GetOneAkun(ctx *gin.Context) {
	var req getAkunRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	akun, err := server.store.GetAuthor(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, akun)

}
