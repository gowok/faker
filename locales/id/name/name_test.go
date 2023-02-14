package name

import (
	"testing"

	"github.com/golang-must/must"
)

func Test_New(t *testing.T) {
	name := New()

	must := must.New(t)
	must.NotNil(name)
	must.NotNil(name.FemaleFirstnames)
	must.NotNil(name.FemaleLastnames)
	must.NotNil(name.MaleFirstnames)
	must.NotNil(name.MaleLastnames)
}
