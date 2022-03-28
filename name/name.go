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

func (f *fakename) Male(opts ...NameOption) string {
	opt := &nameOption{
		length: 2,
	}
	for _, o := range opts {
		o(opt)
	}

	r := random.New()
	var name string
	for i := uint(0); i < opt.length; i++ {
		var nameIndex uint
		if i == 0 {
			nameIndex = r.Uint(0, uint(len(id.MaleFirstnames)-1))
			name = id.MaleFirstnames[nameIndex]
		} else {
			nameIndex = r.Uint(0, uint(len(id.MaleLastnames)-1))
			name = fmt.Sprintf("%s %s", name, id.MaleLastnames[nameIndex])
		}

	}

	return name
}
