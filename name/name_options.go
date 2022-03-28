package name

type NameOption func(*nameOption)

type nameOption struct {
	length uint
}

func WithLength(length uint) NameOption {
	return func(no *nameOption) {
		no.length = length
	}
}
