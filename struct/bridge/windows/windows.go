package windows

import (
	"fmt"

	"tryDesignPattern/v2/struct/bridge/printer"
)

type Windows struct {
	printer printer.Printer
}

func (m *Windows) Print() {
	fmt.Println("Print request for windows")
	m.printer.PrintFile()
}

func (m *Windows) SetPrinter(p printer.Printer) {
	m.printer = p
}
