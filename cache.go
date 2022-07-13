package cache

import "time"

type Cache struct {
	cache map[string]Key
}

type Key struct {
	value    string
	deadline *time.Time
}

func NewCache() Cache {
	result := make(map[string]Key)
	return Cache{cache: result}
}

func (receiver Cache) Get(key string) (string, bool) {
	result, ok := receiver.cache[key]
	if !ok {
		return "", false
	}
	if result.deadline.Before(time.Now()) {
		delete(receiver.cache, key)
		return "", false
	}
	return result.value, true
}

func (receiver) Put(key, value string) {

}

func (receiver) Keys() []string {
}

func (receiver) PutTill(key, value string, deadline time.Time) {
}
