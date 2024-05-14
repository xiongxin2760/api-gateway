package static

// LogPrefix 是该服务生成的logID固定前缀
var LogPrefix string

// context global
const (
	GlobalLogger      string = "Global-Logger"
	GlobalUserInfo    string = "Global-User-Info"
	GlobalLogID       string = "Global-logId"
	GlobalAdapterData string = "Global-Adapter-Data"
	GlobalUserAgent   string = "Global-User-Agent"
	GlobalScene       string = "Global-Scene"
	GlobalQueryMsgID  string = "Global-Query-MsgID"
	GlobalSessionID   string = "Global-SessionId"
	GlobalSource      string = "Global-Source"
	GlobalGcAuthToken string = "Global-GcAuthToken"
)

// 鉴权字段
const (
	// GC字段
	GcLogID       string = "X-Logid"
	GcAdapterData string = "Adapter-Data"
	GcUserAgent   string = "User-Agent"
	GcPerfID      string = "X-Perfid"
	GcBDUSS       string = "BDUSS"
	GcHisign      string = "hisign"
	GcAuthToken   string = "x-gc-auth-token"

	// token鉴权
	AppAuthID        string = "X-Appid"
	AppAuthToken     string = "X-Token"
	AppAuthTimeStamp string = "X-Timestamp"
	AppAuthUserID    string = "X-UuapId"
)
