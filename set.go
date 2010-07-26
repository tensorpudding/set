// The set package implements efficient integer sets
package set

import (
	"fmt"
	"rand"
)

// IntSet provides an interface for general integer set operations.
type IntSet interface {
	// Initialize the set with elements from an array slice.
	// Duplicates are ignored, of course.
 	Init(a []int) *IntSet

	// Insert a new element into the set.
	// Non-destructive at the moment.
 	Insert(x int) *IntSet

	// Tests whether the given element is contained in the set.
 	Elem(x int) bool

	// Takes the union of the two sets, non-destructive.
 	Union(y IntSet) *IntSet

	// Union (parallel version!)
	UnionPar(y IntSet) *IntSet
	
	// Takes the intersection of the two sets, non-destructive.
	Intersection(y *IntSet) *IntSet

	// Takes the set difference of the two sets.
	Diff(y *IntSet) *IntSet

	// remove me
	Display()
}

type IntTreap struct {
	x int           // value
	p int           // randomized priority
	left *IntTreap  // left subtreap
	right *IntTreap // right subtreap
}
func Init(a []int) *IntTreap {
	var s *IntTreap = nil
	for i := 0; i < len(a); i++ {
		s = s.Insert(a[i])
	}
	return s
}

func (s *IntTreap) Insert(x int) *IntTreap {
	if (s == nil) {
		s := new(IntTreap)
		s.x = x
		s.p = rand.Int()
		s.left = nil
		s.right = nil
		return s
	}	
	if (s.x == x) {
		return s
	} else if (s.x > x) {
		if s.left != nil {
			s.left = s.left.Insert(x)
		} else {
			s.left = NewIntTreap()
			s.left.x = x
			s.left.p = rand.Int()
			s.left.left = nil
			s.left.right = nil
		}
		if s.left.p > s.p {
			return RightSwap(s.left, s)
		} else {
			return s
		}
	} else if (s.x < x) {
		if s.right != nil {
			s.right = s.right.Insert(x)
		} else {
			s.right = NewIntTreap()
			s.right.x = x
			s.right.p = rand.Int()
			s.right.left = nil
			s.right.right = nil
		}
		if s.right.p > s.p {
			return LeftSwap(s.right, s)
		} else {
			return s
		}
	}
	return nil
}

func (s *IntTreap) Elem(x int) bool {
	if (s == nil) {
		return false
	}
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

/* RightSwap
 * When the pivot is to the left of the root, do a left tree rotation
 */
func RightSwap(pivot *IntTreap, root *IntTreap) *IntTreap {
	root.left = pivot.right
	pivot.right = root
	return pivot
}

/* LeftSwap
 * When the pivot is to the right of the root, do a right tree rotation
 */
func LeftSwap(pivot *IntTreap, root *IntTreap) *IntTreap {
	root.right = pivot.left
	pivot.left = root
	return pivot
}

/* Split
 * Split the treap by a key k, the left tree is all keys < k,
 * the right tree is all keys > k
 */
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
		Tr.p = s.p
		Tr.right = s.right
		return Tl, m, Tr
	}
	if (k > s.x) {
		T2, m, Tr := s.right.Split(k)
		Tl := NewIntTreap()
		Tl.left = s.left
		Tl.x = s.x
		Tl.p = s.p
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
		tnew.p = s.p
		tnew.right = tr
		return tnew
	} else {
		tl := s.Join(t.left)
		tnew := NewIntTreap()
		tnew.left = tl
		tnew.x = t.x
		tnew.p = t.p
		tnew.right = t.right
		return tnew
	}
	return nil
}

/* Union
 * Takes the union of s and t
 * Destructive?
 */
func (s *IntTreap) Union(t *IntTreap) *IntTreap {
	u := new(IntTreap)
	if (s == nil) { return t }
	if (t == nil) { return s }

	if (s.p < t.p) {
		temp := t
		t = s
		s = temp
	}

	Tl, _, Tr := t.Split(s.x)
	u.x = s.x
	u.p = s.p
	u.left = s.left.Union(Tl)
	u.right = s.right.Union(Tr)
	return u
}

// broken goroutine version

func (s *IntTreap) UnionPar(t *IntTreap) *IntTreap {
	ch := make(chan *IntTreap)
	if (s == nil) { return t }
	if (t == nil) { return s }
	go s.UnionHelper(t, ch)
	u := <- ch
	return u
}

func (s *IntTreap) UnionHelper(t *IntTreap, ch chan *IntTreap) {
	u := new(IntTreap)
	if (s == nil) {
		ch <- t
	} else if (t == nil) {
		ch <- s
	} else if (s.p >= t.p) {
		Tl, _, Tr := t.Split(s.x)
		u.x = s.x
		u.p = s.p
//		chl := make(chan *IntTreap)
//		chr := make(chan *IntTreap)
//		go s.left.UnionHelper(Tl, chl)
//		go s.right.UnionHelper(Tr, chr)
//		u.left = <-chl
//		u.right = <-chr
		u.left = s.left.Union(Tl)
		u.right = s.right.Union(Tr)
		ch <- u
	} else {
		Sl, _, Sr := s.Split(t.x)
		u.x = t.x
		u.p = t.p
// 		chl := make(chan *IntTreap)
// 		chr := make(chan *IntTreap)
// 		go t.left.UnionHelper(Sl, chl)
// 		go t.right.UnionHelper(Sr, chr)
// 		u.left = <-chl
// 		u.right = <-chr
		u.left = s.left.Union(Sl)
		u.right = s.right.Union(Sr)
		ch <- u
	}
}


func NewIntTreap() *IntTreap {
	return new(IntTreap)
}

func NewIntLeaf(x int) *IntTreap {
	l := new(IntTreap)
	l.x = x
	l.p = rand.Int()
	l.left = nil
	l.right = nil
	return l
}

// debugging, remove after
func (s *IntTreap) Display() {
	if s != nil {
		s.left.Display()
		fmt.Printf("%d ", s.x)
		s.right.Display()
	}
}
