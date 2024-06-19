package pizza

type IPizza interface {
	GetPrice() int
}

// 素食主义者
type VeggieMania struct{}

func (p *VeggieMania) GetPrice() int {
	return 15
}

type TomatoTopping struct {
	Pizza IPizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 8
}

type CheeseTopping struct {
	Pizza IPizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}
