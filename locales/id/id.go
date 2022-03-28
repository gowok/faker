package id

import "github.com/gowok/faker/locales/id/name"

type Locale struct {
	name name.Name
}

func New() Locale {
	return Locale{
		name.New(),
	}
}

func (l Locale) Name() name.Name {
	return l.name
}
