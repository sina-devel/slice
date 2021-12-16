package slice

import "fmt"

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func Equal[T comparable](s1, s2 []T) bool {
	return EqualFunc(s1, s2, func(a, b T) bool {
		return a == b
	})
}

// EqualFunc reports whether two slices are equal using a comparison
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// index order, and the comparison stops at the first index for which
// eq returns false.
func EqualFunc[T1, T2 any](s1 []T1, s2 []T2, eq func(T1, T2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if !eq(s1[i], s2[i]) {
			return false
		}
	}
	return true
}

// Index returns the index of the first occurrence of v in s, or -1 if not present.
func Index[T comparable](s []T, v T) int {
	return IndexFunc(s, func(x T) bool {
		return v == x
	})
}

// IndexFunc returns the index into s of the first element
// satisfying f(c), or -1 if none do.
func IndexFunc[T any](s []T, f func(T) bool) int {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

// Contains reports whether v is present in s.
func Contains[T comparable](s []T, v T) bool {
	return Index(s, v) != -1
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[T any](s []T) []T {
	if s == nil {
		return nil
	}

	cloned := make([]T, len(s))
	for i := 0; i < len(s); i++ {
		cloned[i] = s[i]
	}
	return cloned
}

// Insert inserts the values v... into s at index i, returning the modified slice.
// In the returned slice r, r[i] == the first v.  Insert panics if i is out of range.
// This function is O(len(s) + len(v)).
func Insert[T any](s []T, i int, v ...T) []T {
	if i < 0 || i > len(s) {
		panic(fmt.Errorf("runtime error: index out of range [%d] with length %d", i, len(s)))
	}

	if n := len(s) + len(v); n <= cap(s) {
		s2 := s[:n]
		copy(s2[i+len(v):], s[i:])
		copy(s2[i:], v)
		return s2
	}

	s2 := make([]T, len(s)+len(v))
	copy(s2, s[:i])
	copy(s2[i:], v)
	copy(s2[i+len(v):], s[i:])
	return s2
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if s[i:j] is not a valid slice of s.
// Delete modifies the contents of the slice s; it does not create a new slice.
// Delete is O(len(s)), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
func Delete[T any](s []T, i, j int) []T {
	if i < 0 || j < i || j > len(s) {
		panic(fmt.Errorf("runtime error: slice bounds out of range [%d:%d] with length %d", i, j, len(s)))
	}

	copy(s[i:], s[j:])
	var zero T
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = zero
	}

	return s[:len(s)-j+i]
}
