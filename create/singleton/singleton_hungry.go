package singleton

type TheBestFruitSingleton struct{}

var peach *TheBestFruitSingleton

func init() {
	peach = &TheBestFruitSingleton{}
}

func GetHungryInstance() *TheBestFruitSingleton {
	return peach
}
