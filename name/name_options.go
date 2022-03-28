package name

type NameOption func(*nameOption)

type nameOption struct {
	length uint
	locale string
}

func WithLength(length uint) NameOption {
	return func(no *nameOption) {
		no.length = length
	}
}

func WithLocale(name string) NameOption {
	return func(no *nameOption) {
		no.locale = name
	}
}
