package flyweight

type FlyWeight struct {
	Name string
}

func NewFlyWeight(name string) *FlyWeight {
	return &FlyWeight{Name: name}
}

// FlyWeightFactory 享元工厂类
type FlyWeightFactory struct {
	pool map[string]*FlyWeight
}

func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{pool: make(map[string]*FlyWeight)}
}

func (f *FlyWeightFactory) GetFlyWeight(name string) *FlyWeight {
	weight, ok := f.pool[name]
	if !ok {
		weight = NewFlyWeight(name)
		f.pool[name] = weight
	}
	return weight
}
