package faker

import (
	"testing"

	"github.com/golang-must/must"
)

func Test_Random(t *testing.T) {
	must.New(t).NotNil(Random())
}

func Test_Name(t *testing.T) {
	must.New(t).NotNil(Name())
}
