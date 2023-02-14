package id

import (
	"testing"

	"github.com/golang-must/must"
)

func Test_New(t *testing.T) {
	id := New()

	t.Run("not nil", func(t *testing.T) {
		must := must.New(t)
		must.NotNil(id)
	})

	t.Run("has name", func(t *testing.T) {
		name := id.Name()

		must := must.New(t)
		must.NotNil(name)
		must.NotNil(name.FemaleFirstnames)
		must.NotNil(name.FemaleLastnames)
		must.NotNil(name.MaleFirstnames)
		must.NotNil(name.MaleLastnames)
	})
}
