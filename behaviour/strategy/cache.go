package strategy

type Cache struct {
	storage      map[string]string
	EvictionAlgo EvictionAlgo
	capcity      int
	maxCapcity   int
}

func initCache(algo EvictionAlgo) *Cache {
	return &Cache{
		storage:      make(map[string]string),
		EvictionAlgo: algo,
		capcity:      0,
		maxCapcity:   2,
	}
}

func (c *Cache) Add(key, value string) {
	if c.capcity == c.maxCapcity {
		c.evict()
	}
	c.storage[key] = value
	c.capcity++
}

func (c *Cache) get(key string) string {
	return c.storage[key]
}

// 执行逐出策略
func (c *Cache) evict() {
	c.EvictionAlgo.evict(c)
	c.capcity--
}

// 切换策略
func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.EvictionAlgo = e
}
