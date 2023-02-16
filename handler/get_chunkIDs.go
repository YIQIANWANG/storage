package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/app"
)

func GetChunkIDs(c *gin.Context) {
	chunkService := app.Default.GetChunkService()
	data, err := chunkService.GetChunkIDs()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.Data(http.StatusOK, "text/plain", data)
}
