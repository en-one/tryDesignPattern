package simple

import (
	"fmt"
	"testing"
)

func TestNewFruitFactory(t *testing.T) {
	banana := NewFruit("Banana")
	fmt.Println(banana.getName())

	apple := NewFruit("Apple")
	fmt.Println(apple.getName())
}
