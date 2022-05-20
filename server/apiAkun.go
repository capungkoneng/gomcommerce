package api

import (
	"database/sql"
	"net/http"

	db "github.com/capungkoneng/gomcommerce/db/sqlc"
	"github.com/gin-gonic/gin"
)

/****************************************************************/
/************************** CREATE AKUN ONE ************************/
/****************************************************************/
type createAkunRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD IDN"`
}

// Create akun one
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

/****************************************************************/
/************************** GET AKUN ONE ************************/
/****************************************************************/
type getAkunRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

//Get akun one
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

/****************************************************************/
/************************** GET AKUN LIST ************************/
/****************************************************************/
type getListAkunRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

//Get akun list
func (server *Server) GetListAkun(ctx *gin.Context) {
	var req getListAkunRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListAuthorsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	akun, err := server.store.ListAuthors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, akun)

}

/****************************************************************/
/************************** ERROR FUNC ************************/
/****************************************************************/
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
