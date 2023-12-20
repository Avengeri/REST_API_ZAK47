package handler

import (
	"Interface_droch_3/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided JSON data
// @Tags user
// @Param user body model.User true "User data in JSON format"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/user [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Set(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "Пользователь успешно создан"})
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user with the provided JSON data
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID" format(int64)
// @Success 200 {string} string "User get successfully"
// @Failure 500 {object} errorResponse
// @Router /api/user/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)

	user, err := h.service.Get(userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// CheckUser godoc
// @Summary Check if a user exists
// @Description Check if a user with the provided ID exists
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID" format(int64)
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/user/check/{id} [get]
func (h *Handler) CheckUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("Проверка пользователя с ID %d", userID)

	exists, err := h.service.Check(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Пользователь найден: %v", exists)

	if exists {
		c.JSON(http.StatusOK, statusResponse{Status: "Пользователь успешно найден"})
	} else {
		c.JSON(http.StatusNotFound, errorResponse{Message: "Пользователь не найден"})
	}
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user with the provided ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID" format(int64)
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/user/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {

	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Delete(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "Пользователь успешно удален"})
}

// GetAllUsers godoc
// @Summary Get a list of all users
// @Description Get a list of all users with their IDs
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} int "List of user IDs"
// @Failure 400 {object} errorResponse
// @Router /api/user/get_all [get]
func (h *Handler) GetAllUsers(c *gin.Context) {

	userIDs, err := h.service.GetAllId()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, userIDs)
}
