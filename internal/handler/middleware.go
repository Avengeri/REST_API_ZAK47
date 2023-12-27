package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerPorts := strings.Split(header, " ")
	if len(headerPorts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.service.AuthUser.ParseToken(headerPorts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}
