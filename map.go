package jgunify

// KeysOfMap 会接收一个 map 和一个空接口的 slice，然后将 map 的所有键复制到 slice 中。
func KeysOfMap[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
