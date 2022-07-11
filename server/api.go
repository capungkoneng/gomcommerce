package api

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/capungkoneng/gomcommerce/db/sqlc"
	"github.com/capungkoneng/gomcommerce/token"
	"github.com/capungkoneng/gomcommerce/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server services HTTP request for our gomcommerce/bank services
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// New Server creates a new http server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "OPTIONS", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.loginUser)
	router.GET("/users/", server.GetListUser)

	router.GET("/mobil/:id", server.GetMobilOne)
	router.GET("/mobil/", server.GetListMobil)
	router.POST("/mobil", server.CreateMobil)

	router.POST("/kategori", server.CreateKateg)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/akun", server.CreateAkun)
	authRoutes.GET("/akun/:id", server.GetOneAkun)
	authRoutes.GET("/akun/", server.GetListAkun)

	authRoutes.POST("/transfer", server.CreateTransfer)

	// authRoutes.POST("/kategori", server.CreateKateg)

	// authRoutes.GET("/mobil/", server.GetListMobil)

	server.router = router

}

// Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
