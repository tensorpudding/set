package set

import (
	"fmt"
	"rand"
	"testing"
)

const (
	ELEM_TEST_TRIES = 1000
)

func TestSanity (t *testing.T) {
	fmt.Printf("this works!\n")
}

func TestSuit (test *testing.T) {
	s := Init([]int{})
	t := Init([]int{})
	fmt.Printf("\n")
	for i := 0; i < 4; i++ {
		x := rand.Int()
		y := rand.Int()
		s = s.Add(x)
		fmt.Printf("s: ")
		s.Display()
		fmt.Printf("\n")
		t = t.Add(y)
		fmt.Printf("t: ")
		t.Display()
		fmt.Printf("\n")
	}
	u := s.Union(t)
	fmt.Printf("u: ")
	u.Display()
	fmt.Printf("\n")
}
