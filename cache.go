package gi

import (
	"time"
)

// status:
// 0: prepare
// 1: doing
// 2: failed
// 3: done

type CacheDetail struct {
	Status       int
	Msg          string
	ConnectCount int
}

type Cache struct {
	Map   map[string]*CacheDetail
	Total int
	_Chan chan string
}

func NewCache() *Cache {
	return &Cache{Map: map[string]*CacheDetail{}, Total: 0}
}

func (c *Cache) Add(url string) {

	if _, exists := c.Map[url]; exists {
		return
	}
	c.Map[url] = &CacheDetail{Status: 0, ConnectCount: 0}
	c.Total++
}

func (c *Cache) Doing(url string) {
	c.Map[url].Status = 1
}

func (c *Cache) Failed(url string, err error) {
	c.Map[url].Status = 2
	c.Map[url].Msg = err.Error()
}

func (c *Cache) Done(url string) {
	if c.Map[url].Status == 0 {
		return
	}
	c.Map[url].Status = 3
}

func (c *Cache) Refresh(url string) {
	c.Add(url)

	c.Map[url].Status = 0
}

func (c *Cache) Go() {

	for u, d := range c.Map {
		if d.Status == 0 {
			c._Chan <- u
		}
	}

	time.AfterFunc(5*time.Second, c.Go)
}

func (c *Cache) InfoLoop(s int) {
}
