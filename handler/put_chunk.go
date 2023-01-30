package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"storage/app"
)

func PutChunk(c *gin.Context) {
	chunkID := c.Query("chunkID")
	// chunkID := c.PostForm("chunkID")
	// chunkID := c.Request.Form.Get("chunkID")
	chunkData, _ := ioutil.ReadAll(c.Request.Body)

	chunkService := app.Default.GetChunkService()
	err := chunkService.PutChunk(chunkID, chunkData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success."})
}
