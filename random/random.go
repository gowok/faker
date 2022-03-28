package random

import (
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	ErrMinMoreThanMax = errors.New("min should less than max")
)

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numerals = "0123456789"
	Ascii    = Alphabet + Numerals + "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
)

func New() *randomizer {
	rand.Seed(time.Now().UnixNano())
	return &randomizer{}
}

type randomizer struct {
}

func (r randomizer) Uint(min, max uint) (v uint) {
	if min > max {
		log.Fatal(ErrMinMoreThanMax)
	}
	return uint(rand.Intn(int(max-min))) + min
}

func (r randomizer) String(length uint, opts ...RandomStringOption) (v string) {
	opt := &stringOpts{}
	for _, o := range opts {
		o(opt)
	}

	max := uint(len(Alphabet) - 1)
	for i := uint(0); i < length; i++ {
		v += string(Alphabet[r.Uint(0, max)])
	}

	if opt.cases == lowerCase {
		return strings.ToLower(v)
	}

	if opt.cases == upperCase {
		return strings.ToUpper(v)
	}

	return
}

func (r randomizer) Ascii(length uint) (v string) {
	max := uint(len(Ascii) - 1)
	for i := uint(0); i < length; i++ {
		v += string(Ascii[r.Uint(0, max)])
	}
	return
}
