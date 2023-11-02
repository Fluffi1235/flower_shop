package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/web1/internal/models"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	err := c.BindJSON(&input)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.CreateUser(input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type singInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input singInInput

	err := c.BindJSON(&input)
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.Service.GenerateToken(input.Login, input.Password)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
