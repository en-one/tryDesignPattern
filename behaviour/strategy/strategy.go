package strategy

/*

	下面是一个·内存缓存的情景，由于内存大小存在限制。所以达到上限后一些条目就必须背移除以留出空间：
		- 最少最近使用LRU：移除最近最少使用的一条标签
		- 先进先出FIFO：移除最先进入的一条标签
		- 最小使用Lfu：移除最少使用的一条标签

	策略模式场景发挥作用。创建一系列算法，每个算法都实现相同的接口。假设为evictionAlgo
	限制，
*/

type EvictionAlgo interface {
	evict(c *Cache)
}

// 先进先出
type Fifo struct{}

func (f *Fifo) evict(c *Cache) {
	println("Evicting by fifo strategy")
}

// 最少最近使用LRU
type Lru struct{}

func (l *Lru) evict(c *Cache) {
	println("Evicting by lru strategy")
}

// 最小使用Lfu
type Lfu struct{}

func (l *Lfu) evict(c *Cache) {
	println("Evicting by lfu strategy")
}
