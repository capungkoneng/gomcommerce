package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/capungkoneng/gomcommerce/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

//Get Mobil list
type GetMobilJoinParams struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

func (server *Server) GetListMobil(ctx *gin.Context) {
	var req GetMobilJoinParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.GetMobilJoinManyParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	mobil, err := server.store.GetMobilJoinMany(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println(mobil)
	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{
		"result": mobil,
	},
		"code":   http.StatusOK,
		"status": "success",
	})
}

type createMobilRequest struct {
	Nama       string         `json:"nama"`
	Deskripsi  sql.NullString `json:"deskripsi"`
	KategoriID int64          `json:"kategori_id"`
	UserID     string         `json:"user_id"`
	Gambar     sql.NullString `json:"gambar"`
	Trf6jam    int64          `json:"trf_6jam"`
	Trf12jam   int64          `json:"trf_12jam"`
	Trf24jam   int64          `json:"trf_24jam"`
	Seat       sql.NullInt64  `json:"seat"`
	TopSpeed   sql.NullInt64  `json:"top_speed"`
	MaxPower   sql.NullInt64  `json:"max_power"`
	Pintu      sql.NullInt64  `json:"pintu"`
	Gigi       sql.NullString `json:"gigi"`
}

// Create User one
func (server *Server) CreateMobil(ctx *gin.Context) {
	var req createMobilRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMobilParams{
		Nama:       req.Nama,
		Deskripsi:  req.Deskripsi,
		KategoriID: req.KategoriID,
		UserID:     req.UserID,
		Gambar:     req.Gambar,
		Trf6jam:    req.Trf6jam,
		Trf12jam:   req.Trf12jam,
		Trf24jam:   req.Trf24jam,
		Seat:       req.Seat,
		TopSpeed:   req.TopSpeed,
		MaxPower:   req.MaxPower,
		Pintu:      req.Pintu,
		Gigi:       req.Gigi,
	}

	mobil, err := server.store.CreateMobil(ctx, arg)
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
	ctx.JSON(http.StatusOK, mobil)
}

//Get akun one
type getMobileOneRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetMobilOne(ctx *gin.Context) {
	var req getMobileOneRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	mobil, err := server.store.GetMobilJoinOne(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, mobil)

}
