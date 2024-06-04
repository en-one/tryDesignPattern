package mac

import (
	"fmt"
	"tryDesignPattern/v2/struct/bridge/printer"
)

type Mac struct {
	printer printer.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p printer.Printer) {
	m.printer = p
}
