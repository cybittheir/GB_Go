package main

import (
	"fmt"
	"time"
)

const cacheLife = 2 // time.Second
const delay = 3     // time.Second

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}

var _ Cache = (*cacheImpl)(nil)

// Доработает конструктор и методы кеша, так чтобы они соответствовали интерфейсу Cache
func newCacheImpl() *cacheImpl {
	return &cacheImpl{}
}

type cacheImpl struct {
	data map[string]string
	exp  time.Time
}

func (c *cacheImpl) Get(k string) (string, bool) {
	if c.data[k] != "" && c.exp.Add(cacheLife*time.Second).Compare(time.Now()) > 0 { // checking cache lifetime
		return c.data[k], true
	} else if c.data[k] != "" {
		delete(c.data, k) // delete cache record if lifetime expired
	}
	return "", false
}

func (c *cacheImpl) Set(k, v string) {
	c.data[k] = v
	c.exp = time.Now()
}

// аналог Set для dbs
func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "works", "best": "practic", "sweet": "dream"}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	if ok {
		d.cache.Set(k, v) // adding to cache
	}
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
	c := newCacheImpl()

	db := newDbImpl(c)
	c.data = map[string]string{}

	// cache warming
	for k, v := range db.dbs {
		if k != "best" {
			c.Set(k, v)
			time.Sleep(delay * time.Second)
		}
	}

	fmt.Println(db.Get("test"))

	fmt.Println(db.Get("cool"))

	fmt.Println(db.Get("best"))

	fmt.Println(db.Get("test"))

	fmt.Println(db.Get("best"))

	fmt.Println(db.Get("sweet"))

	fmt.Println(db.Get("hello"))

}
