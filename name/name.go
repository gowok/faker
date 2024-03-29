package name

import (
	"fmt"

	"github.com/gowok/faker/locales"
	"github.com/gowok/faker/random"
)

type Fakename struct {
}

func New() *Fakename {
	return &Fakename{}
}

func (f *Fakename) getRandomIndex(list []string) uint {
	r := random.New()
	return r.Uint(0, uint(len(list)-1))
}

func (f *Fakename) buildOpts(fallback *nameOption, opts ...NameOption) *nameOption {
	if fallback == nil {
		fallback = &nameOption{
			length: 2,
			locale: "id",
		}
	}
	for _, o := range opts {
		o(fallback)
	}

	return fallback
}

func (f *Fakename) fromFirstnamesAndLastnames(firstnames, lastnames []string, opt *nameOption) string {
	var name string
	for i := uint(0); i < opt.length; i++ {
		if i != opt.length-1 {
			nameIndex := f.getRandomIndex(firstnames)
			if i == 0 {
				name = firstnames[nameIndex]
			} else {
				name = fmt.Sprintf("%s %s", name, firstnames[nameIndex])
			}
		} else {
			nameIndex := f.getRandomIndex(lastnames)
			name = fmt.Sprintf("%s %s", name, lastnames[nameIndex])
		}
	}

	return name
}

func (f *Fakename) Male(opts ...NameOption) string {
	opt := f.buildOpts(nil, opts...)
	l, _ := locales.Get(opt.locale)
	return f.fromFirstnamesAndLastnames(l.Name().MaleFirstnames, l.Name().MaleLastnames, opt)
}

func (f *Fakename) Female(opts ...NameOption) string {
	opt := f.buildOpts(nil, opts...)
	l, _ := locales.Get(opt.locale)
	return f.fromFirstnamesAndLastnames(l.Name().FemaleFirstnames, l.Name().FemaleLastnames, opt)
}
