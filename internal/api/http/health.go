package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *DiscoverHandler) Health(c *gin.Context) {
	err := h.healthUsecase.HealthCheck(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "unhealthy"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
