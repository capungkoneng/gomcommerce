package api

import (
	"database/sql"
	"net/http"

	db "github.com/capungkoneng/gomcommerce/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

/****************************************************************/
/************************** CREATE AKUN ONE ************************/
/****************************************************************/
type createKategRequest struct {
	NamaKategori string         `json:"nama_kategori" binding:"required"`
	Deskripsi    sql.NullString `json:"deskripsi" binding:"required,max=100"`
}

// Create akun one
func (server *Server) CreateKateg(ctx *gin.Context) {
	var req createKategRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateKategoriParams{
		NamaKategori: req.NamaKategori,
		Deskripsi:    req.Deskripsi,
	}

	kateg, err := server.store.CreateKategori(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, kateg)
}
