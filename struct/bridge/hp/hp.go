package hp

import "fmt"

type Hp struct {
}

func (h *Hp) PrintFile() {
	fmt.Println("Printing by HP printer")
}
