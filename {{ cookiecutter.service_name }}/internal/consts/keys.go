package consts

type ContextKey int

const (
	SessionKey      ContextKey = 0
	LoggerKey       ContextKey = 1
	ErrorMessageKey ContextKey = 2
	ErrorCodeKey    ContextKey = 3
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationBearerKey  = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func (key ContextKey) String() string {
	switch key {
	case SessionKey:
		return "session"

	case LoggerKey:
		return "logger"

	case ErrorMessageKey:
		return "error_message"

	case ErrorCodeKey:
		return "error_code"

	default:
		return "N/A"
	}
}
