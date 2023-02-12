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
	LogInit()
	app.InitDefault()
	ReporterInit()
}

func LogInit() {
	if !service.PathExists(conf.LogFilePath) {
		err := os.MkdirAll(conf.LogFilePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	// error log
	dataLogFileName := fmt.Sprintf("%s/%s-%s-%s", conf.LogFilePath, time.Now().Format("2006"), time.Now().Format("01"), time.Now().Format("02"))
	dataLogFile, err := os.OpenFile(dataLogFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm) // 创建、追加、读写，777（所有权限）
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(os.Stdout, dataLogFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// operation Log
	opLogFileName := conf.LogFilePath + "/" + conf.OpLogName
	data.OpLogFile, err = os.OpenFile(opLogFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func ReporterInit() {
	heartbeatService := app.Default.GetHeartbeatService()
	err := heartbeatService.InitReport()
	if err != nil {
		panic(err)
	}
	heartbeatService.StartReport()
}
