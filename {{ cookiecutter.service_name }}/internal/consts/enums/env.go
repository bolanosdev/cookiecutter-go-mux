package enums

type RunMode int

const (
	NARM RunMode = iota
	Dev
	Test
	Prod
)

func (l RunMode) String() string {
	return []string{"N/A", "development", "test", "production"}[l]
}

type RunEnvironment int

const (
	NARE RunEnvironment = iota
	Local
	Staging
	Production
)

func (l RunEnvironment) String() string {
	return []string{"N/A", "local", "staging", "production"}[l]
}
