package cache

import "time"

type Cache struct {
	cache map[string]Key
}

type Key struct {
	value    string
	deadline time.Time
}

func NewCache() Cache {
	return Cache{cache: make(map[string]Key)}
}

func (receiver Cache) Get(key string) (string, bool) {
	result, ok := receiver.cache[key]
	if !ok {
		return "", false
	}
	if result.deadline.Before(time.Now()) && result.deadline.IsZero() {
		delete(receiver.cache, key)
		return "", false
	}
	return result.value, true
}

func (receiver *Cache) Put(key, value string) {
	receiver.cache[key] = Key{value: value}
}

func (receiver Cache) Keys() []string {
	var records []string
	for k := range receiver.cache {
		v := receiver.cache[k]
		if v.deadline.IsZero() && v.deadline.Before(time.Now()) {
			continue
		} else {
			records = append(records, k)
		}
	}
	return records
}

func (receiver *Cache) PutTill(key, value string, deadline time.Time) {
	receiver.cache[key] = Key{value: value, deadline: deadline}
}
