package strategy

import (
	"testing"
)

func Test_demo(t *testing.T) {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Add("c", "3")

	cache.SetEvictionAlgo(&Fifo{})

	cache.Add("d", "4")

	cache.SetEvictionAlgo(&Lru{})
	cache.Add("e", "5")

}
