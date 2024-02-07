package server

import "github.com/gin-gonic/gin"

func (server *Server) InitRouter() {
	v1 := server.router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
	}
}
