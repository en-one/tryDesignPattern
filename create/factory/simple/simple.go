package simple

// 水果接口
type IFruit interface {
	setName(name string)
	getName() string
}

// 具体水果类
type Fruit struct {
	name string
}

func (f *Fruit) setName(name string) {
	f.name = name
}
func (f *Fruit) getName() string {
	return f.name
}

// 具体水果-----苹果
type Apple struct{ Fruit }

func newApple() IFruit {
	return &Apple{
		Fruit: Fruit{name: "apple"},
	}
}

// 具体水果-----香蕉
type Banana struct{ Fruit }

func newBanana() IFruit {
	return &Banana{
		Fruit: Fruit{name: "banana"},
	}
}

// 重写香蕉的名字
func (b *Banana) getName() string {
	return b.name + "rewrite"
}

// NewFruitFactory 当我们以NewXXX的方式创建对象/接口的时候，当返回为接口的时候，就是简单工厂模式
func NewFruit(t string) IFruit {
	switch t {
	case "Apple":
		return newApple()
	case "Banana":
		return newBanana()
	}
	return nil
}
