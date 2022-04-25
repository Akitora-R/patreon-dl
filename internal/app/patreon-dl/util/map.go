package util

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	r := map[K]V{}
	for k, v := range m {
		r[k] = v
	}
	return r
}
