package keys

type ContextKeys int

const SessionKey ContextKeys = 0

func (key ContextKeys) String() string {
	switch key {
	case 0:
		return "session"
	default:
		return "N/A"
	}
}
