package array

func Chunk[T any](arr []T, chunkSize int) [][]T {
	var result [][]T
	for i := 0; i < len(arr); i += chunkSize {
		end := i + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}
	return result
}

func In[T comparable](arr []T, target T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func Map[T, E any](arr []T, fn func(T) E) []E {
	result := make([]E, len(arr))
	for _, v := range arr {
		result = append(result, fn(v))
	}
	return result
}

func MapExtractKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapGetOrDefault[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	if val, ok := m[key]; ok {
		return val
	} else {
		return defaultVal
	}
}
