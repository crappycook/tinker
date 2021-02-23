package lru

import "container/list"

// Non-thread safe LRU cache
type LRU struct {
	capacity  int
	hash      map[string]*list.Element
	evictList *list.List
}

type entry struct {
	key   string
	value interface{}
}

// New constructs LRU cache
func New(capacity int) *LRU {
	return &LRU{
		capacity:  capacity,
		hash:      make(map[string]*list.Element),
		evictList: list.New(),
	}
}

// Get the key from LRU cache
func (lru *LRU) Get(key string) (interface{}, bool) {
	if e, ok := lru.hash[key]; ok {
		lru.evictList.MoveToFront(e)
		if e.Value.(*entry) == nil {
			return nil, false
		}
		return e.Value.(*entry).value, true
	}

	return nil, false
}

// Put key and value to LRU cache
func (lru *LRU) Put(key string, val interface{}) {
	// element already exists
	if e, ok := lru.hash[key]; ok {
		lru.evictList.MoveToFront(e)
		e.Value.(*entry).value = val
		return
	}
	// add new entry
	en := &entry{key: key, value: val}
	element := lru.evictList.PushFront(en)
	lru.hash[key] = element

	if lru.Len() > lru.capacity {
		// evict
		lru.removeOldest()
	}
	return
}

// Len returns the lru cache current length
func (lru *LRU) Len() int {
	return lru.evictList.Len()
}

// Peek finds the value without moving the element
func (lru *LRU) Peek(key string) (interface{}, bool) {
	if ele, ok := lru.hash[key]; ok {
		val := ele.Value.(*entry).value
		return val, ok
	}
	return nil, false
}

// Contains checks if a key is in the cache, without updating the recent-ness
func (lru *LRU) Contains(key string) (ok bool) {
	_, ok = lru.hash[key]
	return ok
}

// GetOldest returns the oldest element in the cache
func (lru *LRU) GetOldest(key string) (interface{}, bool) {
	tail := lru.evictList.Back()
	if tail != nil {
		return tail.Value.(*entry).value, true
	}
	return nil, false
}

// Resize the LRU cache capacity
func (lru *LRU) Resize(capacity int) int {
	diff := lru.capacity - capacity
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff; i++ {
		// scale down
		lru.removeOldest()
	}
	lru.capacity = capacity
	return diff
}

// Delete removes the provided key from the cache
func (lru *LRU) Delete(key string) {
	if ent, ok := lru.hash[key]; ok {
		lru.removeElement(ent)
	}
}

func (lru *LRU) removeOldest() {
	tail := lru.evictList.Back()
	if tail != nil {
		lru.removeElement(tail)
	}
}

func (lru *LRU) removeElement(e *list.Element) {
	lru.evictList.Remove(e)
	entry := e.Value.(*entry)
	delete(lru.hash, entry.key)
}
