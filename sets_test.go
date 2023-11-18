package sets_test

import (
	"reflect"
	"testing"

	"github.com/R167/go-sets"
)

func assertEqual[E any, A any](t *testing.T, expected E, actual A) {
	t.Helper()
	if v, ok := any(expected).(interface{ Equal(A) bool }); ok && v.Equal(actual) {
		return
	}
	if reflect.DeepEqual(expected, actual) {
		return
	}
	t.Errorf("expected: %v, actual: %v", expected, actual)
}

func TestNew(t *testing.T) {
	s := sets.New[int](1, 2, 3)
	assertEqual(t, 3, len(s))
	assertEqual(t, true, s.Contains(1))
	assertEqual(t, true, s.Contains(2))
	assertEqual(t, true, s.Contains(3))
}

func TestFromMap(t *testing.T) {
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	s := sets.FromMap(m)
	assertEqual(t, 3, len(s))

	ref := sets.New[int](1, 2, 3)
	assertEqual(t, ref, s)
}

func TestSet_Add(t *testing.T) {
	s := sets.New[int](1, 2, 3)
	s.Add(4)
	assertEqual(t, 4, len(s))
	assertEqual(t, true, s.Contains(4))
	s.Add(4)
	assertEqual(t, 4, len(s))
}

func TestSet_Remove(t *testing.T) {
	s := sets.New[int](1, 2, 3)
	assertEqual(t, true, s.Remove(2))
	assertEqual(t, 2, len(s))
	assertEqual(t, false, s.Remove(2))
	assertEqual(t, 2, len(s))
	assertEqual(t, false, s.Contains(2))
}

func TestSet_Contains(t *testing.T) {
	s := sets.New[int](1, 2, 3)
	assertEqual(t, true, s.Contains(2))
	assertEqual(t, false, s.Contains(4))
}

func TestSet_Union(t *testing.T) {
	s := sets.New[int](1, 2, 3)
	other := sets.New[int](3, 4, 5, 6)
	assertEqual(t, sets.New[int](1, 2, 3, 4, 5, 6), s.Union(other))
	assertEqual(t, sets.New[int](1, 2, 3, 4, 5, 6), other.Union(s))
}

func TestSet_Intersect(t *testing.T) {
	s := sets.New[int](1, 2, 3)
	other := sets.New[int](3, 4, 5, 6)
	assertEqual(t, sets.New[int](3), s.Intersect(other))
	assertEqual(t, sets.New[int](3), other.Intersect(s))
}

func TestSet_Equal(t *testing.T) {
	s := sets.New[int](1, 2, 3)
	other := sets.New[int](3, 4, 5)
	assertEqual(t, false, s.Equal(other))

	s2 := s.Clone()
	assertEqual(t, true, s.Equal(s2))
	s2.Add(9)
	assertEqual(t, false, s.Equal(s2))
}

func TestSet_Difference(t *testing.T) {
	s := sets.New[int](1, 2, 3)
	other := sets.New[int](3, 4, 5)
	assertEqual(t, sets.New[int](1, 2), s.Difference(other))
	assertEqual(t, sets.New[int](4, 5), other.Difference(s))
}
