package locales

import (
	"testing"

	"github.com/golang-must/must"
)

func Test_Get(t *testing.T) {
	locale, ok := Get("id")

	must := must.New(t)
	must.True(ok)
	must.NotNil(locale)
}
