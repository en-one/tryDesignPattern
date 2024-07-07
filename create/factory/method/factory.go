package simple

/*

如何实现简单工厂模式
	1、 Go 本身是没有构造函数的，一般而言我们采用 NewName  的方式创建对象/接口，当它返回的是接口的时候，其实就是简单工厂模式

如何实现工厂方法模式：
	1、构建一个接口IFruit
	2、构建一个结构体Fruit，实现接口IFruit
	3、构建工厂的结构体Apple，Banana。结构体中包含第二步中的结构体Fruit；同时每个工厂结构体可以重写接口中方法
	4、写一个new函数，返回IFruit。根据传入的参数，返回不同的结构体
*/

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
