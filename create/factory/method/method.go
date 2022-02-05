package method

import (
	"factory/simple"
)

type FruitFactory interface {
	CreateIntroduce() simple.Fruit
}

type AppleFactory struct{}

func (a AppleFactory) CreateIntroduce() simple.Fruit {
	return simple.Apple{}
}

type BananaFactory struct{}

func (BananaFactory) CreateIntroduce() simple.Fruit {
	return simple.Banana{}
}

func NewFruitFactory(t string) FruitFactory {
	switch t {
	case "apple":
		return AppleFactory{}
	case "banana":
		return BananaFactory{}
	}
	return nil
}
