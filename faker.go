package faker

import (
	"github.com/gowok/faker/name"
	"github.com/gowok/faker/random"
)

var defaultRandom = random.New()
var defaultName = name.New()

func Random() *random.Randomizer {
	return defaultRandom
}

func Name() *name.Fakename {
	return defaultName
}
