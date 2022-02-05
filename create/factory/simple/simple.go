package simple

import "fmt"

type FruitFactory interface {
	Introduce(name string) string
}

type Apple struct{}

func (Apple) Introduce(name string) string {
	return fmt.Sprintf("i was an %s", name)
}

type Banana struct{}

func (Banana) Introduce(name string) string {
	return fmt.Sprintf("i was a %s", name)
}

// NewFruitFactory 当我们以NewXXX的方式创建对象/接口的时候，当返回为接口的时候，就是简单工厂模式
func NewFruitFactory(t string) FruitFactory {
	switch t {
	case "Apple":
		return Apple{}
	case "Banana":
		return Banana{}
	}
	return nil
}
