package util

import (
	"testing"
)

func TestRandomSeed(t *testing.T) {
	s, err := RandomSeed()
	t.Log(s, err)
}
