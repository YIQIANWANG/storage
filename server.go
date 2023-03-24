package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"storage/app"
	"storage/conf"
	"storage/data"
	"storage/service"
	"time"
)

func init() {
	initLog()
	initChunk()
	app.InitDefault()
	initReporter()
}

func initLog() {
	if service.PathExists(conf.LogFilePath) {
		_ = os.RemoveAll(conf.LogFilePath)
	}
	err := os.MkdirAll(conf.LogFilePath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Log
	logFileName := fmt.Sprintf("%s/%s-%s-%s", conf.LogFilePath, time.Now().Format("2006"), time.Now().Format("01"), time.Now().Format("02"))
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm) // 创建、追加、读写，777（所有权限）
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Operation Log
	opLogFileName := conf.LogFilePath + "/" + conf.OpLogName
	data.OpLog, err = os.OpenFile(opLogFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func initChunk() {
	if service.PathExists(conf.ChunkFilePath) {
		_ = os.RemoveAll(conf.ChunkFilePath)
	}
	err := os.MkdirAll(conf.ChunkFilePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func initReporter() {
	heartbeatService := app.Default.GetHeartbeatService()
	err := heartbeatService.InitReport()
	if err != nil {
		panic(err)
	}
	heartbeatService.StartReport()
}
