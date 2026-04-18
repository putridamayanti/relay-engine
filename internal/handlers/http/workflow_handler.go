package http

import (
	"net/http"
	"relay-engine/internal/models"
	"relay-engine/internal/services"

	"github.com/gin-gonic/gin"
)

type WorkflowHandler struct {
	service services.WorkflowService
}

func NewWorkflowHandler(service services.WorkflowService) *WorkflowHandler {
	return &WorkflowHandler{service: service}
}

func (h *WorkflowHandler) Create(c *gin.Context) {
	var req models.WorkflowDefinition
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
