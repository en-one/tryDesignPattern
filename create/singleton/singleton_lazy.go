package singleton

import (
	"sync"
)

var (
	peach2 *TheBestFruitSingleton
	once   = &sync.Once{}
)

func GetLazyInstance() *TheBestFruitSingleton {
	if peach2 != nil {
		once.Do(func() {
			peach2 = &TheBestFruitSingleton{}
		})
	}
	return peach2
}
