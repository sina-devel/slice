package slice

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
