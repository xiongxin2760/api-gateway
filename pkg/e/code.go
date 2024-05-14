package e

type ErrorCode int

const (
	ZeroCode        ErrorCode = 0
	Success         ErrorCode = 200
	PartSuccess     ErrorCode = 206
	Error           ErrorCode = 500
	InvalidParams   ErrorCode = 400
	TokenExpired    ErrorCode = 401
	StatusForbidden ErrorCode = 403
	NotFound        ErrorCode = 404
	RateLimit       ErrorCode = 429
	VoiceErrorCode  ErrorCode = 30001
)
