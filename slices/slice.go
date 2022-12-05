package slices

// Unique returns true if slice does not contain any duplicates.
func Unique[T comparable](slice []T) bool {
	exists := make(map[T]struct{})

	for _, elem := range slice {
		if _, ok := exists[elem]; ok {
			return false
		}

		exists[elem] = struct{}{}
	}

	return true
}

// Exists returns true if elem is found in slice.
func Exists[T comparable](slice []T, elem T) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}

	return false
}

// Index returns the index of the first occurrence of elem in slice.
// If elem is not found, then -1 is returned.
func Index[T comparable](slice []T, elem T) int {
	for i, e := range slice {
		if e == elem {
			return i
		}
	}

	return -1
}
