package slice

import (
	"sort"
)

type sliceFn[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (sf sliceFn[T]) Len() int           { return len(sf.s) }
func (sf sliceFn[T]) Swap(i, j int)      { sf.s[i], sf.s[j] = sf.s[j], sf.s[i] }
func (sf sliceFn[T]) Less(i, j int) bool { return sf.cmp(sf.s[i], sf.s[j]) }

// Sort sorts s in order as determined by the cmp function.
// The sort is not guaranteed to be stable.
func Sort[T any](s []T, cmp func(T, T) bool) {
	sort.Sort(sliceFn[T]{s, cmp})
}
