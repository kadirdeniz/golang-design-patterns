package main

import "fmt"

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Printing by windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}
