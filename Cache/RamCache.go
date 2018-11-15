package Cache

import (
	"errors"
	"sync"
)

func GetNewRamCache() RamCache {
	data := make(map[string]string)
	return RamCache{data: data}
}

type RamCache struct {
	sync.Mutex
	data map[string]string
}

func (c *RamCache) Contains(key string) bool {
	_, ok := c.data[key]
	return ok
}

func (c *RamCache) Put(key string, value string) error {
	c.Lock()
	c.data[key] = value
	c.Unlock()
	return nil
}

func (c *RamCache) Get(key string) (string, error) {
	val, ok := c.data[key]
	if ok {
		return val, nil
	}
	return "", errors.New("value not exist")
}

func (c *RamCache) Delete(key string) error {
	_, ok := c.data[key]
	if ok {
		delete(c.data, key)
		return nil
	}
	return errors.New("value not exist")
}
