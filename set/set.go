package set

type set[K comparable] map[K]struct{}

// Non-thread safe set
func NewSet[K comparable]() set[K] {
	return make(set[K])
}

func (s set[K]) Has(key K) bool {
	_, ok := s[key]
	return ok
}

func (s set[K]) Put(key K) {
	s[key] = struct{}{}
}

func (s set[K]) Delete(key K) {
	delete(s, key)
}
