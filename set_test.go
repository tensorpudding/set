package set

import (
	"fmt"
//	"rand"
	"container/set"
	"testing"
)

const (
	ELEM_TEST_TRIES = 1000
)

func TestSanity (t *testing.T) {
	fmt.Printf("this works!\n")
}

func TestEasy (t *testing.T) {
	var s *set.IntTreap = new(set.IntTreap)
//	s := set.Init([]int{})
	s.Display()
}
