package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHealthStatus returns health state of the riddles service
func GetHealthStatus(c *gin.Context) {

	c.JSON(http.StatusOK, "OK")
}
