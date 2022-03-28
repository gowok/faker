package name

import (
	"fmt"

	"github.com/gowok/faker/name/id"
	"github.com/gowok/faker/random"
)

type fakename struct {
}

func New() *fakename {
	return &fakename{}
}

func (f *fakename) getRandomIndex(list []string) uint {
	r := random.New()
	return r.Uint(0, uint(len(list)-1))
}

func (f *fakename) FromFirstnamesAndLastnames(firstnames, lastnames []string, opts ...NameOption) string {
	opt := &nameOption{
		length: 2,
	}
	for _, o := range opts {
		o(opt)
	}

	var name string
	for i := uint(0); i < opt.length; i++ {
		if i == 0 {
			nameIndex := f.getRandomIndex(firstnames)
			name = firstnames[nameIndex]
		} else {
			nameIndex := f.getRandomIndex(lastnames)
			name = fmt.Sprintf("%s %s", name, lastnames[nameIndex])
		}
	}

	return name
}

func (f *fakename) Male(opts ...NameOption) string {
	return f.FromFirstnamesAndLastnames(id.MaleFirstnames, id.MaleLastnames, opts...)
}

func (f *fakename) Female(opts ...NameOption) string {
	return f.FromFirstnamesAndLastnames(id.FemaleFirstnames, id.FemaleLastnames, opts...)
}
