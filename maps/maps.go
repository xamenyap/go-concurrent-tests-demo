package maps

// Keys returns all keys in m.
func Keys[K comparable, V any](m map[K]V) []K {
	ks := make([]K, 0)
	for k := range m {
		ks = append(ks, k)
	}

	return ks
}
