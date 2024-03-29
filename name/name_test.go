package name

import (
	"strings"
	"testing"

	"github.com/golang-must/must"
)

var Name = New()

func TestName(t *testing.T) {
	t.Run("can fake male name with length", func(t *testing.T) {
		must := must.New(t)
		expected := uint(5)
		result := Name.Male(WithLength(expected))
		actual := uint(len(strings.Split(result, " ")))
		must.Equal(expected, actual)
	})

	t.Run("can fake female name with length", func(t *testing.T) {
		must := must.New(t)
		expected := uint(5)
		result := Name.Female(WithLength(expected))
		actual := uint(len(strings.Split(result, " ")))
		must.Equal(expected, actual)
	})

	t.Run("can fake male name with locale", func(t *testing.T) {
		expected := "id"
		result := &nameOption{}

		WithLocale(expected)(result)

		must := must.New(t)
		must.Equal(expected, result.locale)
	})

}
