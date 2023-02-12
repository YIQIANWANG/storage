package main

import (
	"github.com/gin-gonic/gin"
	"storage/conf"
	"storage/handler"
)

func main() {
	router := gin.Default()
	router.POST("/chunk", handler.PutChunk)
	router.GET("/chunk", handler.GetChunk)
	router.DELETE("/chunk", handler.DelChunk)
	router.GET("chunkIDs", handler.GetChunkIDs)
	router.GET("/PING", handler.PING)
	router.GET("/RESET", handler.RESET)

	err := router.Run(":" + conf.PORT)
	if err != nil {
		panic(err)
	}
}
