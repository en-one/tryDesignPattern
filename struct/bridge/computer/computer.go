package main

import "tryDesignPattern/v2/struct/bridge/printer"

type Computer interface {
	Print()
	SetPrinter(printer.Printer)
}
