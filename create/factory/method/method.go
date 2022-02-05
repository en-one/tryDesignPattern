package method

import (
	"factory/simple"
)

type FruitFactoryMethod interface {
	CreateIntroduce() simple.FruitFactory
}

type AppleFactory struct{}

func (a AppleFactory) CreateIntroduce() simple.FruitFactory {
	return simple.Apple{}
}

type BananaFactory struct{}

func (BananaFactory) CreateIntroduce() simple.FruitFactory {
	return simple.Banana{}
}

func NewFruitFactoryMethod(t string) FruitFactoryMethod {
	switch t {
	case "apple":
		return AppleFactory{}
	case "banana":
		return BananaFactory{}
	}
	return nil
}
