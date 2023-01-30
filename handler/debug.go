package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"storage/conf"
)

func PING(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PONG!"})
}

func RESET(c *gin.Context) {
	_ = os.RemoveAll(conf.ChunkFilePath)
	_ = os.RemoveAll(conf.LogFilePath)
	log.Println("RESET")
	c.JSON(http.StatusOK, gin.H{"message": "Success."})
}
