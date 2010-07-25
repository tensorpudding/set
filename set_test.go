package set

import (
	"rand"
	"container/set"
	"testing"
)

const (
	ELEM_TEST_TRIES 1000
)

func TestSetElem (t *testing.T) {
	s := set.Init()
	for i := 0; i < ELEM_TEST_TRIES; i++ {
		p := rand.Int()
		s.Add(p)
		if !s.Elem(p) { Fail() }
	}
}
