package abstract

import (
	"factory/simple"
)

type FruitAbstractFactory interface {
	CreateApple() simple.Fruit
	CreateBanner() simple.Fruit
}

type AbstractFruitFactory struct{}

func NewSimpleFruitFactory() FruitAbstractFactory {
	return &AbstractFruitFactory{}
}

func (s *AbstractFruitFactory) CreateApple() simple.Fruit {
	return &simple.Apple{}
}

func (s *AbstractFruitFactory) CreateBanner() simple.Fruit {
	return &simple.Banana{}
}
