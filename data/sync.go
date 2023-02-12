package data

import (
	"os"
	"sync"
)

// OperationCount 操作数，并发不安全需要加锁
var OperationCount int64
var Lock sync.RWMutex

// OpLogFile 操作日志，追加写并发安全
var OpLogFile *os.File
