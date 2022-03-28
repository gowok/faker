package random

import (
	"testing"
	"unicode"

	"github.com/golang-must/must"
)

var Random = New()

func TestRandom(t *testing.T) {
	t.Run("can random string", func(t *testing.T) {
		must := must.New(t)
		expected := 5
		result := Random.String(uint(expected))
		actual := len(result)
		must.Equal(expected, actual)
	})

	t.Run("can random string with lowercase", func(t *testing.T) {
		must := must.New(t)
		expected := 5
		result := Random.String(uint(expected), WithLowerCase())
		actual := len(result)
		must.Equal(expected, actual)
		for _, r := range result {
			must.True(unicode.IsLower(r))
		}
	})

	t.Run("can random string with uppercase", func(t *testing.T) {
		must := must.New(t)
		expected := 5
		result := Random.String(uint(expected), WithUpperCase())
		actual := len(result)
		must.Equal(expected, actual)
		for _, r := range result {
			must.True(unicode.IsUpper(r))
		}
	})

	t.Run("can random ascii", func(t *testing.T) {
		must := must.New(t)
		expected := 5
		result := Random.Ascii(uint(expected))
		actual := len(result)
		must.Equal(expected, actual)
	})

}
