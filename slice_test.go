package slice_test

import (
	"testing"

	. "github.com/sina-devel/slice"
)

func TestEqual(t *testing.T) {
	testCases := []struct {
		name string
		s1   []int
		s2   []int
		want bool
	}{
		{
			name: "different lengths",
			s1:   []int{1, 2, 3},
			s2:   []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "not equal",
			s1:   []int{1, 3, 4},
			s2:   []int{1, 2, 4},
			want: false,
		},
		{
			name: "equal",
			s1:   []int{1, 2, 4},
			s2:   []int{1, 2, 4},
			want: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Equal(tc.s1, tc.s2); got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestEqualFunc(t *testing.T) {
	type person struct {
		age uint
	}

	testCases := []struct {
		name string
		s1   []person
		s2   []person
		eq   func(a, b person) bool
		want bool
	}{
		{
			name: "different lengths",
			s1:   []person{{1}, {2}, {3}},
			s2:   []person{{1}, {2}, {3}, {4}},
			eq:   func(a, b person) bool { return a.age == b.age },
			want: false,
		},
		{
			name: "not equal",
			s1:   []person{{1}, {2}, {3}},
			s2:   []person{{1}, {2}, {4}},
			eq:   func(a, b person) bool { return a.age == b.age },
			want: false,
		},
		{
			name: "equal",
			s1:   []person{{1}, {2}, {3}},
			s2:   []person{{1}, {2}, {3}},
			eq:   func(a, b person) bool { return a.age == b.age },
			want: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := EqualFunc(tc.s1, tc.s2, tc.eq); got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
		v    int
		want int
	}{
		{
			name: "nil slice",
			s:    nil,
			v:    4,
			want: -1,
		},
		{
			name: "zero length",
			s:    []int{},
			v:    4,
			want: -1,
		},
		{
			name: "not exist in slice",
			s:    []int{1, 3, 4},
			v:    5,
			want: -1,
		},
		{
			name: "exist",
			s:    []int{1, 3, 4},
			v:    3,
			want: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Index(tc.s, tc.v); got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
		v    int
		want bool
	}{
		{
			name: "nil slice",
			s:    nil,
			v:    4,
			want: false,
		},
		{
			name: "zero length",
			s:    []int{},
			v:    4,
			want: false,
		},
		{
			name: "not exist in slice",
			s:    []int{1, 3, 4},
			v:    5,
			want: false,
		},
		{
			name: "exist",
			s:    []int{1, 3, 4},
			v:    3,
			want: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Contains(tc.s, tc.v); got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestClone(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "nil slice",
			s:    nil,
			want: nil,
		},
		{
			name: "zero length",
			s:    []int{},
			want: []int{},
		},
		{
			name: "clone",
			s:    []int{1, 3, 4},
			want: []int{1, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Clone(tc.s); !Equal(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
		cmp  func(a, b int) bool
		want []int
	}{
		{
			name: "asc",
			s:    []int{4, 3, 5},
			cmp:  func(a, b int) bool { return a < b },
			want: []int{3, 4, 5},
		},
		{
			name: "dsc",
			s:    []int{4, 3, 5},
			cmp:  func(a, b int) bool { return a > b },
			want: []int{5, 4, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if Sort(tc.s, tc.cmp); !Equal(tc.s, tc.want) {
				t.Errorf("got %v, want %v", tc.s, tc.want)
			}
		})
	}
}
