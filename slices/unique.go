package slices

func UniqueBy[T any, K comparable](xs []T, keyFn func(T) K) []T {
	seen := make(map[K]bool)
	j := 0
	for _, item := range xs {
		key := keyFn(item)
		if seen[key] {
			continue
		}
		seen[key] = true
		xs[j] = item
		j++
	}
	return xs[:j]
}
