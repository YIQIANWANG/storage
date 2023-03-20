package conf

// 日志文件存放路径
const (
	LogFilePath = "logs"
	OpLogName   = "oplog"
)

// ChunkFilePath 数据文件存放路径
const (
	ChunkFilePath = "chunks"
)

// 服务配置
const (
	IP                = "localhost"
	PORT              = "8800"
	GROUP             = "group0"
	Capacity          = 10000000000 // 最大容量10GB
	HeartbeatInternal = 2           // 心跳间隔为2s
)

// MongoDB配置
const (
	PROTOCOL       = "mongodb"
	USERNAME       = "mongouser"
	PASSWORD       = "YqMTE*5873QpUJ"
	ADDRESS        = "9.134.32.73:27017,9.134.38.231:27017,9.134.47.32:27017"
	AUTHENTICATION = "somedb?authSource=admin"
	DATABASE       = "localhost" // 本地测试
	// DATABASE       = "stress" // 压力测试
	// DATABASE       = "cos"    // 生产环境
)
