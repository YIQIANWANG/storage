package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PING(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PONG!"})
}
