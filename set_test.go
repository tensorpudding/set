package set

import (
	"fmt"
	"rand"
	"testing"
)

const (
	ELEM_TEST_TRIES = 10
	UNION_TEST_TRIES = 10
)

func randomizeSet(n int) *IntTreap {
	s := Init([]int{})
	for i := 0; i < n; i++ {
		p := rand.Int()
		s = s.Insert(p)
	}
	return s
}

func TestSanity (t *testing.T) {
	fmt.Printf("this works!\n")
}

func TestAddElem (test *testing.T) {
	fmt.Printf("Testing TestElem...\n")
	s := Init([]int{})
	for i := 0; i < ELEM_TEST_TRIES; i++ {
		p := rand.Int()
		s = s.Insert(p)
		if !s.Elem(p) {
			test.Fail()
		}
	}
}

func TestUnion (test *testing.T) {
	s := randomizeSet(UNION_TEST_TRIES)
	t := randomizeSet(UNION_TEST_TRIES)
	u := s.Union(t)
	u.Display()
}		

func TestParUnion (test *testing.T) {
	fmt.Printf("Testing parallelized union...")
	s := randomizeSet(UNION_TEST_TRIES)
	t := randomizeSet(UNION_TEST_TRIES)
	u := s.UnionPar(t)
	u.Display()
}

func BenchmarkUnion (b *testing.B) {
	s := randomizeSet(b.N)
	t := randomizeSet(b.N)
	s.Union(t)
}

func BenchmarkParUnion (b *testing.B) {
	s := randomizeSet(b.N)
	t := randomizeSet(b.N)
	s.UnionPar(t)
}	
		
