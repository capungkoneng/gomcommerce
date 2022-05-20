package api

import (
	db "github.com/capungkoneng/gomcommerce.git/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server services HTTP request for our gomcommerce/bank services
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// New Server creates a new http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/akun", server.CreateAkun)
	router.GET("/akun/:id", server.GetOneAkun)
	router.GET("/akun/", server.GetListAkun)

	server.router = router
	return server
}

// Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
