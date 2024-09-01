package api

import (
	db "github.com/Sanpeta/rinha-backend-2023-q3-golang/db/sqlc"
	"github.com/Sanpeta/rinha-backend-2023-q3-golang/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our service
type Server struct {
	config util.Config
	store  *db.SQLStore
	router *gin.Engine
}

func CORSConfig() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}
		context.Next()
	}
}

func NewServer(config util.Config, store *db.SQLStore) (*Server, error) {
	server := &Server{
		store:  store,
		config: config,
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORSConfig())

	server.router = router

	//Rotas
	router.POST("/pessoas", server.createPessoa)
	router.GET("/pessoas/:id", server.getPessoa)
	router.GET("/pessoas", server.listPessoas)
	router.GET("/contagem-pessoas", server.contagemPessoas)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has error:": err.Error()}
}
