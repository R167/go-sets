package sets

import "fmt"

// It's perfectly fine to modify Set directly, so use a type alias for convenience
// rather than a wrapper type.
type empty = struct{}

// Set is a set of comparable elements of type E. Note since Set is backed by a map,
// the zero value of Set is not an empty set. Use New[E]() to create an empty set.
//
// Set is not thread-safe.
type Set[E comparable] map[E]empty

// New returns a new set containing the given elements.
func New[E comparable](elements ...E) Set[E] {
	s := make(Set[E], len(elements))
	for _, e := range elements {
		s[e] = empty{}
	}
	return s
}

// FromMap returns a new Unordered set containing the keys of the given map.
func FromMap[E comparable, T any, M ~map[E]T](m M) Set[E] {
	s := make(Set[E], len(m))
	for e := range m {
		s[e] = empty{}
	}
	return s
}

// Add the given element to s and returns itself.
func (s Set[E]) Add(e ...E) Set[E] {
	for _, e := range e {
		s[e] = empty{}
	}
	return s
}

// Delete the given element from s. Returns true if the element was in s.
// If you need to remove multiple elements at once, use Subtract.
func (s Set[E]) Delete(e E) bool {
	if _, ok := s[e]; ok {
		delete(s, e)
		return true
	}
	return false
}

// Subtract removes multipe elements from s and returns itself
func (s Set[E]) Subtract(e ...E) Set[E] {
	for _, e := range e {
		delete(s, e)
	}
	return s
}

// Has returns true if e is in s.
func (s Set[E]) Has(e E) bool {
	_, ok := s[e]
	return ok
}

// Union returns a new set containing all the elements in s and other.
//
//	s := New[int](1, 2, 3)
//	other := New[int](3, 4, 5)
//	union := s.Union(other)
//	// union contains 1, 2, 3, 4, 5
func (s Set[E]) Union(other Set[E]) Set[E] {
	if len(s) > len(other) {
		s, other = other, s
	}
	union := FromMap(other)
	for e := range s {
		union.Add(e)
	}
	return union
}

// Intersection returns a new set containing the elements in both s and other.
//
//	s := New[int](1, 2, 3)
//	other := New[int](3, 4, 5)
//	intersection := s.Intersection(other)
//	// intersection contains 3
func (s Set[E]) Intersection(other Set[E]) Set[E] {
	if len(s) > len(other) {
		s, other = other, s
	}
	intersection := make(Set[E])
	for e := range s {
		if other.Has(e) {
			intersection.Add(e)
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
func (s Set[E]) Difference(other Set[E]) Set[E] {
	difference := New[E]()
	for e := range s {
		if !other.Has(e) {
			difference.Add(e)
		}
	}
	return difference
}

// Equal returns true if s and other contain all the same elements.
func (s Set[E]) Equal(other Set[E]) bool {
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
func (s Set[E]) Slice() []E {
	slice := make([]E, 0, len(s))
	for e := range s {
		slice = append(slice, e)
	}
	return slice
}

// String returns a string representation of s.
func (s Set[E]) String() string {
	return fmt.Sprintf("Set%v", s.Slice())
}

// Clone returns a shallow copy of s.
func (s Set[E]) Clone() Set[E] {
	return FromMap(s)
}

func (s Set[E]) subset(other Set[E]) bool {
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
func (s Set[E]) Subset(other Set[E]) bool {
	return s.subset(other)
}

// Superset returns true if s is a superset of other.
func (s Set[E]) Superset(other Set[E]) bool {
	return other.Subset(s)
}
