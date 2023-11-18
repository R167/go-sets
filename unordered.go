package sets

import "fmt"

type empty = struct{}

// Unordered is a set of comparable elements of type E. Note since Unordered is backed by a map,
// the zero value of Unordered is not an empty set. Use New[E]() to create an empty set.
//
// Unordered is not thread-safe.
type Unordered[E comparable] map[E]empty

// New returns a new set containing the given elements.
func New[E comparable](elements ...E) Unordered[E] {
	s := make(Unordered[E], len(elements))
	for _, e := range elements {
		s[e] = empty{}
	}
	return s
}

// FromMap returns a new Unordered set containing the keys of the given map.
func FromMap[E comparable, T any, M ~map[E]T](m M) Unordered[E] {
	s := make(Unordered[E], len(m))
	for e := range m {
		s[e] = empty{}
	}
	return s
}

// Insert adds the given element to s.
func (s Unordered[E]) Insert(e ...E) {
	for _, e := range e {
		s[e] = empty{}
	}
}

// Delete removes the given element from s. Returns true if the element was in s.
func (s Unordered[E]) Delete(e E) bool {
	if _, ok := s[e]; ok {
		delete(s, e)
		return true
	}
	return false
}

// Has returns true if e is in s.
func (s Unordered[E]) Has(e E) bool {
	_, ok := s[e]
	return ok
}

// Union returns a new set containing all the elements in s and other.
//
//	s := New[int](1, 2, 3)
//	other := New[int](3, 4, 5)
//	union := s.Union(other)
//	// union contains 1, 2, 3, 4, 5
func (s Unordered[E]) Union(other Unordered[E]) Unordered[E] {
	if len(s) > len(other) {
		s, other = other, s
	}
	union := FromMap(other)
	for e := range s {
		union.Insert(e)
	}
	return union
}

// Intersection returns a new set containing the elements in both s and other.
//
//	s := New[int](1, 2, 3)
//	other := New[int](3, 4, 5)
//	intersection := s.Intersection(other)
//	// intersection contains 3
func (s Unordered[E]) Intersection(other Unordered[E]) Unordered[E] {
	if len(s) > len(other) {
		s, other = other, s
	}
	intersection := make(Unordered[E])
	for e := range s {
		if other.Has(e) {
			intersection.Insert(e)
		}
	}
	return intersection
}

// Difference returns a new set containing the elements in s that are not in other.
//
//	s := New[int](1, 2, 3)
//	other := New[int](3, 4, 5)
//	sMinusOther := s.Difference(other)
//	// contains 1, 2
//	otherMinusS := other.Difference(s)
//	// contains 4, 5
func (s Unordered[E]) Difference(other Unordered[E]) Unordered[E] {
	difference := New[E]()
	for e := range s {
		if !other.Has(e) {
			difference.Insert(e)
		}
	}
	return difference
}

// Equal returns true if s and other contain all the same elements.
func (s Unordered[E]) Equal(other Unordered[E]) bool {
	if len(s) != len(other) {
		return false
	}
	for e := range s {
		if !other.Has(e) {
			return false
		}
	}
	return true
}

// Slice returns a slice of all the elements in s.
//
// Note: The order of the output elements is undefined.
func (s Unordered[E]) Slice() []E {
	slice := make([]E, 0, len(s))
	for e := range s {
		slice = append(slice, e)
	}
	return slice
}

// String returns a string representation of s.
func (s Unordered[E]) String() string {
	return fmt.Sprintf("Set%v", s.Slice())
}

func (s Unordered[E]) Clone() Unordered[E] {
	return FromMap(s)
}

func (s Unordered[E]) subset(other Unordered[E]) bool {
	if len(s) > len(other) {
		return false
	}
	for e := range s {
		if !other.Has(e) {
			return false
		}
	}
	return true
}

// Subset returns true if s is a subset of other.
func (s Unordered[E]) Subset(other Unordered[E]) bool {
	return s.subset(other)
}

// Superset returns true if s is a superset of other.
func (s Unordered[E]) Superset(other Unordered[E]) bool {
	return other.Subset(s)
}
