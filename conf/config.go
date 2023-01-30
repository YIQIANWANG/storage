package conf

// 文件路径
const (
	LogFilePath   = "logs"
	ChunkFilePath = "chunks"
	//HashFilePath = "hashes"
)

// 操作日志文件
const OpLogName = "oplog"

// 服务配置
const (
	IP    = "localhost"
	PORT  = "8800"
	GROUP = "group0"
)

// MongoDB
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
