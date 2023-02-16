package conf

// 文件存放路径
const (
	LogFilePath   = "logs"
	OpLogName     = "oplog"
	ChunkFilePath = "chunks"
	//HashFilePath = "hashes"
)

// 服务配置
const (
	IP    = "localhost"
	PORT  = "8800"
	GROUP = "group0"
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
