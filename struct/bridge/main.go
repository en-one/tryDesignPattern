package main

import (
	"fmt"
	"tryDesignPattern/v2/struct/bridge/epson"
	"tryDesignPattern/v2/struct/bridge/hp"
	"tryDesignPattern/v2/struct/bridge/mac"
	"tryDesignPattern/v2/struct/bridge/windows"
)

func main() {

	// 打印机
	hpPrinter := &hp.Hp{}
	epsonPrinter := &epson.Epson{}

	// 电脑
	macComputer := &mac.Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println("----")

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println("----")

	windowsComputer := &windows.Windows{}
	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()
	fmt.Println("----")

	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
	fmt.Println("----")

}
