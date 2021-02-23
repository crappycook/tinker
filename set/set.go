package set

type set map[interface{}]struct{}

// Non-thread safe set
func NewSet() set {
	return make(set)
}

func (s set) Has(key interface{}) bool {
	_, ok := s[key]
	return ok
}

func (s set) Put(key interface{}) {
	s[key] = struct{}{}
}

func (s set) Delete(key interface{}) {
	delete(s, key)
}
