package handler

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Message string
}

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message string) {
	h.Logger.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
