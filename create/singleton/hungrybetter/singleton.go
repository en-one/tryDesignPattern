package hungrybetter

import (
	"fmt"
	"sync"
)

/*
1\在 Go 语言中，使用 sync.Once 可以简洁地实现线程安全的单例模式。
2、sync.Once 可以确保单例模式是线程安全的，并且只会执行一次
*/
type singleton struct {
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	// 第一次检查
	if instance == nil {
		once.Do(func() {
			instance = &singleton{}
			fmt.Println("Creating single instance now.")
		})
	} else {
		fmt.Println("Single instance already created.")
	}
	return instance
}
