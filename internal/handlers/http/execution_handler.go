package http

import (
	"net/http"
	"relay-engine/internal/models"
	"relay-engine/internal/services"

	"github.com/gin-gonic/gin"
)

type ExecutionHandler struct {
	service services.ExecutionService
}

func NewExecutionHandler(service services.ExecutionService) *ExecutionHandler {
	return &ExecutionHandler{service: service}
}

func (h *ExecutionHandler) Start(c *gin.Context) {
	var req models.WorkflowExecution
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Start(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
