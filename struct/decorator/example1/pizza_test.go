package pizza

import (
	"fmt"
	"testing"
)

func TestColorSquare_Draw(t *testing.T) {
	pizza := &VeggieMania{}
	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		Pizza: pizza,
	}

	// Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.Pizza.GetPrice())
}
