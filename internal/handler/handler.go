package handler

import (
	_ "Interface_droch_3/internal/model"
	"Interface_droch_3/internal/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		auth := api.Group("/user")
		{
			auth.POST("/", h.CreateUser)
			auth.GET("/:id", h.GetUser)
			auth.GET("/check/:id", h.CheckUser)
			auth.DELETE("/:id", h.DeleteUser)
			auth.GET("/get_all", h.GetAllUsers)
		}
	}
	return router
}
