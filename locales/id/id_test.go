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

}
