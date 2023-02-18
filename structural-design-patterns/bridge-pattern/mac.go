package main

import "fmt"

type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Printing by mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}
