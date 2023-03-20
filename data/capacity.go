package data

import "sync"

// AvailableCap 可用容量，并发不安全需要加锁
var AvailableCap int
var AvailableCapLock sync.Mutex
