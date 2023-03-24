package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/app"
)

func DelChunk(c *gin.Context) {
	chunkID := c.Query("chunkID")

	chunkService := app.Default.GetChunkService()
	err := chunkService.DelChunk(chunkID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success."})
}
