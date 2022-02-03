package singleton

type TheBestFruitSingleton struct{}

var peach *TheBestFruitSingleton

func init() {
	peach = &TheBestFruitSingleton{}
}

// GetHungryInstance 饿汉式
// 类加载时候， instance 就已经创建好了
func GetHungryInstance() *TheBestFruitSingleton {
	return peach
}
