package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/app"
)

func GetChunk(c *gin.Context) {
	chunkID := c.Query("chunkID")

	chunkService := app.Default.GetChunkService()
	data, err := chunkService.GetChunk(chunkID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.Data(http.StatusOK, "text/plain", data)
}
