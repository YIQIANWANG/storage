package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/app"
)

func Synchronize(c *gin.Context) {
	dstAddress := c.Query("dstAddress")
	mode := c.Query("mode")

	syncService := app.Default.GetSyncService()
	var err error
	if mode == "chunk" {
		chunkID := c.Query("chunkID")
		err = syncService.SyncChunk(dstAddress, chunkID)
	} else if mode == "all" {
		err = syncService.SyncAll(dstAddress)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success."})
}
