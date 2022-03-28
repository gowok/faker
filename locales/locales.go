package locales

import (
	"github.com/gowok/faker/locales/id"
	"github.com/gowok/faker/locales/id/name"
)

type locale interface {
	Name() name.Name
}

var locales = map[string]locale{
	"id": id.New(),
}

func Get(name string) (locale, bool) {
	l, ok := locales[name]
	return l, ok
}
