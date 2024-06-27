package singleton

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

// HungrySingleton 饿汉式单例模式，类加载时就创建一个单例实例
type single struct{}

var singleInstance *single

func getSingleInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		// 双重检查锁定，通过两次检查单例实例是否创建加，减少不必要的同步开销
		if singleInstance == nil {
			singleInstance = &single{}
			fmt.Println("Creating single instance now.")
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}
