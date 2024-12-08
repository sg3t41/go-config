package enum

type RunMode int

const (
	Default RunMode = iota
	Development
	Staging
	Production
)

func (r RunMode) String() string {
	switch r {
	case Development:
		return "development"
	case Staging:
		return "staging"
	case Production:
		return "production"
	default:
		return "default"
	}
}
