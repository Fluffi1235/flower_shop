package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/web1/internal/services"
	"go.uber.org/zap"
)

type Handler struct {
	Service *services.Service
	Logger  *zap.Logger
}

func NewHandler(service *services.Service, logger *zap.Logger) *Handler {
	return &Handler{
		Service: service,
		Logger:  logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		shop := api.Group("/shop")
		{
			shop.GET("/basket", h.getItems)
		}
	}
	return router
}
