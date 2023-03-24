package data

import (
	"os"
	"sync"
)

// OpCount 操作数，并发不安全需要加锁
var OpCount int64
var OpCountLock sync.Mutex

// OpLog 操作日志文件，追加写并发安全
var OpLog *os.File
