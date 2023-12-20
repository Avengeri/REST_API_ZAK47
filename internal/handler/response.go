package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

type statusResponse struct {
	Status string `json:"status"`
}
type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
