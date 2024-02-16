package server

import (
	"dcard-2024-backend-intern-assignment/controllers"
	"dcard-2024-backend-intern-assignment/repositories"
	"dcard-2024-backend-intern-assignment/services"
)

func (server *Server) InitRouter() {
	v1 := server.router.Group("/api/v1")
	{
		adRouter := v1.Group("/ad")
		{
			adRepository := repositories.NewAdRepository(server.db.GetPool())
			adService := services.NewAdService()
			adController := controllers.NewAdController(*adRepository, *adService)
			adRouter.POST("/", adController.CreateAd)
		}
	}
}
