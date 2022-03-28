package faker

type stringCase string
type RandomStringOption func(*stringOpts)

const (
	lowerCase stringCase = "lowerCase"
	upperCase stringCase = "upperCase"
)

type stringOpts struct {
	cases stringCase
}

func WithLowerCase() func(*stringOpts) {
	return func(so *stringOpts) {
		so.cases = lowerCase
	}
}

func WithUpperCase() func(*stringOpts) {
	return func(so *stringOpts) {
		so.cases = upperCase
	}
}
