package handler

import (
	"Interface_droch_3/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @Summary Registering a new user
// @Description Registers a new user and returns his ID
// @Accept json
// @Tags auth
// @Produce json
// @Param input body model.UserExampleRegistration true "New user's data"
// @Success 200 {string} string  "A user with the specified ID has been successfully created"
// @Failure 400 {object} errorResponse "Validation error or incorrect data"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /auth/sign-up [post]
func (h *Handler) signUP(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.AuthUser.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type signInInput struct {
	Username string `json:"username" binding:"required" example:"ZAK"`
	Password string `json:"password" binding:"required" example:"qwerty"`
}

// @Summary User authentication
// @Description Authenticates the user and returns the access token
// @Accept json
// @Tags auth
// @Produce json
// @Param input body signInInput true "User authentication data"
// @Success 200 {string} string "Successful authentication, returns the access token"
// @Failure 400 {object} errorResponse "Validation error or incorrect data"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /auth/sign-in [post]
func (h *Handler) signIN(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.AuthUser.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
