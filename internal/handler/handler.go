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
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUP)
		auth.POST("/sign-in", h.signIN)
	}
	api := router.Group("/api", h.userIdentity)
	{
		user := api.Group("/user")
		{
			user.POST("/", h.CreateUser)
			user.GET("/:id", h.GetUser)
			user.GET("/check/:id", h.CheckUser)
			user.DELETE("/:id", h.DeleteUser)
			user.GET("/get_all", h.GetAllUsers)
		}
	}
	return router
}
