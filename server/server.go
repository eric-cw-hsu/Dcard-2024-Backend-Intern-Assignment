package server

import (
	"dcard-2024-backend-intern-assignment/databases"
	"dcard-2024-backend-intern-assignment/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config utils.Config
	db     databases.BaseDatabase
}

func NewServer(config utils.Config, database databases.BaseDatabase) *Server {
	return &Server{
		router: gin.Default(),
		config: config,
		db:     database,
	}
}

func (server *Server) StartDatabaseConnection() {
	server.db.Connect()
}

func (server *Server) CloseDatabaseConnection() {
	server.db.Close()
}

func (server *Server) Serve() error {
	return server.router.Run(fmt.Sprintf(":%s", server.config.PORT))
}
