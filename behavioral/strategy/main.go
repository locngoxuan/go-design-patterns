package main

import "log"

type EvictAlgorithm interface {
	Evict(c *Cache)
}

type Lru struct{}

func (l *Lru) Evict(c *Cache) {
	log.Println("Evicting by Lru strategy")
}

type Lfu struct{}

func (l *Lfu) Evict(c *Cache) {
	log.Println("Evicting by Lfu strategy")
}

type Fifo struct{}

func (l *Fifo) Evict(c *Cache) {
	log.Println("Evicting by Fifo strategy")
}

type Cache struct {
	storage     map[string]string
	evictAlgo   EvictAlgorithm
	capacity    int
	maxCapacity int
}

func initCache(e EvictAlgorithm) *Cache {
	return &Cache{
		storage:     make(map[string]string),
		evictAlgo:   e,
		capacity:    0,
		maxCapacity: 2,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictAlgorithm) {
	c.evictAlgo = e
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.Evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) {
	delete(c.storage, key)
}

func (c *Cache) Evict() {
	c.evictAlgo.Evict(c)
	c.capacity--
}

func main() {
	c := initCache(&Fifo{})
	c.Add("key1", "value1")
	c.Add("key2", "value2")
	c.Add("key3", "value3")

	c.SetEvictionAlgo(&Lru{})
	c.Add("key4", "value4")

	c.SetEvictionAlgo(&Lfu{})
	c.Add("key5", "value5")
}
