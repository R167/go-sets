package sets_test

import (
	"fmt"
	"testing"

	"github.com/R167/go-sets"
)

var splits = []struct {
	a int
	b int
}{
	{10, 10},
	{10, 50},
	{10, 100},
}

func makeSet(n int) sets.Set[int] {
	s := make(sets.Set[int], n)
	for i := 0; i < n; i++ {
		s[i] = struct{}{}
	}
	return s
}

func BenchmarkIntersection(b *testing.B) {
	for _, split := range splits {
		s1 := makeSet(split.a)
		s2 := makeSet(split.b)
		b.Run(fmt.Sprintf("%d-%d", split.a, split.b), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s1, s2 = s2, s1
				s1.Intersection(s2)
			}
		})
	}
}

func BenchmarkUnion(b *testing.B) {
	for _, split := range splits {
		s1 := makeSet(split.a)
		s2 := makeSet(split.b)
		b.Run(fmt.Sprintf("%d-%d", split.a, split.b), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s1, s2 = s2, s1
				s1.Union(s2)
			}
		})
	}
}
