package types

// Config 配置文件
type Config struct {
	// AppName 应用名称
	AppName string

	// RunMode 运行模式，可配置值：
	// debug    : 调试，    对应常量 env.RunModeDebug
	// test     : 测试，    对应常量 env.RunModeTest
	// release  : 线上发布， 对应常量 env.RunModeRelease
	RunMode string

	Env string

	Port int // 开放端口

	// HTTPServer http server的 配置
	HTTPServer HTTPServerConfig

	// logger配置
	Logger LoggerSetting

	// HTTPClient 配置
	HTTPClient HTTPClientSetting

	// GrpcSetting 配置
	GrpcSetting GrpcClientSetting

	// mysql 配置
	MysqlSetting MysqlClientSetting
}

// HTTPServerConfig Config http server 的配置内容
type HTTPServerConfig struct {
	ReadTimeout  int // 单位 ms
	WriteTimeout int // 单位 ms
	IdleTimeout  int // ms
}

// LoggerSetting config log 的配置内容
type LoggerSetting struct {
	LogDir             string
	LogFileName        string
	GinWebLogName      string
	CronjobLogFileName string
	RecoverLogFileName string
}

type HTTPClientSetting struct {
	MaxIdleConnsPerHost int
	Timeoutms           int // 单位 ms
}

type GrpcClientSetting struct {
	Timeout              int64
	KeepaliveTime        int
	KeepaliveTimeout     int
	MaxIdle              int
	MaxActive            int
	MaxConcurrentStreams int
	Gcms                 string
	Ums                  string
}

type KafkaSetting struct {
	Server                 []string
	EnableTLS              bool
	TopicRealTimeKnowledge string
	TopicLongKnowledge     string
	TopicShortKnowledge    string
	TopicPush              string
	PemFilePath            string
	KeyFilePath            string
	CAFilePath             string
}

type MysqlClientSetting struct {
	MessageDB string
}
