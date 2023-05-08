// Set
package set

import (
	"errors"
)

// Append takes a slice and an element of an orderable type and turns a new
// slice of that must contain the given element exaclty once.
// Returns `error` if the slice contains elements that are not unique, to avoid
// destructive changes to the slice. Does not error if the element to be added
// already exists (expected behavior).
// Use `EnsureUnique` or the `Ss` constructor to ensure the slice only contains
// unique elements.
// Order of elements is not preserved!
func Append[T comparable](s []T, e T) ([]T, error) {
	m := make(map[T]struct{}, len(s))

	for _, el := range s {
		if _, exists := m[el]; exists {
			return nil, errors.New("slice contains non-unique elements")
		}

		m[el] = struct{}{}
	}

	m[e] = struct{}{}
	result := make([]T, len(m), len(m))
	for el := range m {
		result = append(result, el)
	}
	return result, nil
}

// Contains takes a slice and an element of the same comparable type and
// returns true if the slice contains the element AT LEAST once. Does not
// error, if the elements in the slice are not unique, because the operation
// will not be destructive on the slice.
func Contains[T comparable](s []T, e T) bool {
	for _, el := range s {
		if el == e {
			return true
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Intersection of two sets returns all elements of s1 that also belong to s2.
// Does not error if the given slices contain non-unique elements, but the
// resulting slice will only contain each element as most once.
func Intersection[T comparable](s1, s2 []T) []T {
	m1 := make(map[T]struct{}, len(s1))
	for _, e := range s1 {
		m1[e] = struct{}{}
	}
	m2 := make(map[T]struct{}, len(s2))
	for _, e := range s2 {
		m2[e] = struct{}{}
	}

	result := make([]T, 0, max(len(m1), len(m2)))
	for e := range m1 {
		if _, exists := m2[e]; exists {
			result = append(result, e)
		}
	}

	return result
}
