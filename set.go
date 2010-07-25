package set

import (
	"rand"
)

type IntSet interface {
 	Elem(x int) bool
 	Add(x int) *IntSet
 	Union(y IntSet) *IntSet
 	Intersection(y *IntSet) *IntSet
	Diff(y *IntSet) *IntSet
 	Init(a []int) *IntSet
}

type IntTreap struct {
	x int // value
	p int // randomized priority
	left *IntTreap  // left subtreap
	right *IntTreap // right subtreap
}
func Init(a []int) *IntTreap {
	s := NewIntTreap()
	for i := 0; i < len(a); i++ {
		s.Add(a[i])
	}
	return s
}

func (s *IntTreap) Add(x int) *IntTreap {
	return s
// 	if (s.x == x) {
// 		return s
// 	} else if (s.x > x) {
// 		if (s.left == nil) {
// 			s.left = NewIntTreap()
// 			s.left.x = x
// 			s.left.p = rand.Int()
// 			if (
}

func (s *IntTreap) Split(k int) (*IntTreap, int, *IntTreap) {
	if (s == nil) {
		return nil, 0, nil
	}
	if (k == s.x) {
		return s.left, s.x, s.right
	}
	if (k < s.x) {
		Tl, m, T2 := s.left.Split(k)
		Tr := NewIntTreap()
		Tr.left = T2
		Tr.x = s.x
		Tr.right = s.right
		return Tl, m, Tr
	}
	if (k > s.x) {
		T2, m, Tr := s.right.Split(k)
		Tl := NewIntTreap()
		Tl.left = s.left
		Tl.x = s.x
		Tl.right = T2
		return Tl, m, Tr
	}
	return nil, 0, nil
}

func (s *IntTreap) Join(t *IntTreap) *IntTreap {
	if (s == nil) {
		return t
	} else if (t == nil) {
		return s
	} else if (s.p > t.p) {
		tr := s.right.Join(t)
		tnew := NewIntTreap()
		tnew.left = s.left
		tnew.x = s.x
		tnew.right = tr
		return tnew
	} else {
		tl := s.Join(t.left)
		tnew := NewIntTreap()
		tnew.left = tl
		tnew.x = t.x
		tnew.right = t.right
		return tnew
	}
	return nil
}

// func (s *IntTreap) Union(t *IntTreap) *IntTreap {
// 	u := NewIntTreap()
// 	if (s == nil) { return t }
// 	if (t == nil) { return s }
	
// 	if (s.p < t.p) Swap(s, t)
// 	less, x, gtr := t.Split(s.x)
// 	u.x = s.x
// 	u.p = s.p
// 	u.left = s.left.Union(less)
// 	u.right = s.right.Union(gtr)
// 	return u
// }

// func (s *IntTreap) Intersection(t *IntTreap) *IntTreap {
// 	i := NewIntTreap()
// 	if (s == nil) { return nil }
// 	if (t == nil) { return nil }

// 	if (s.p < t.p) Swap(s, t)
// 	less, x, gtr := t.Split(s.x)

// }

// func Swap(s *IntTreap, t *IntTreap) *IntTreap {
// 	pivot := t
// 	root 
// }

func NewIntTreap() *IntTreap {
	t:= new(IntTreap)
	t.x = 0
	t.p = rand.Int()
	t.left = nil
	t.right = nil
	return t
}

func (s *IntTreap) Elem(x int) bool {
	if (s.x == x) {
		return true
	} else if (s.x > x) {
		if (s.left == nil) {
			return false
		}
		return s.left.Elem(x)
	} else {
		if (s.right == nil) {
			return false
		}
		return s.right.Elem(x)
	}
	return false
}
