package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/app"
)

func DelChunk(c *gin.Context) {
	chunkID := c.Query("chunkID")

	chunkService := app.Default.GetChunkService()
	size, err := chunkService.DelChunk(chunkID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err, "size": -1})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success.", "size": size})
}
