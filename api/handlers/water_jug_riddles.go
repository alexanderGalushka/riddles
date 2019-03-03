package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetWaterJugRiddlesSolution returns step by step solution to water jug riddle
func GetWaterJugRiddlesSolution(c *gin.Context) {
	itemID, err := uc.GetURIParam(c, lc.ItemIDKey)

	c.JSON(http.StatusOK, result)
}

