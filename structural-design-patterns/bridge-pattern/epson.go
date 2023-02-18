package main

import "fmt"

type epson struct{}

func (e *epson) printFile() {
	fmt.Println("Printing by epson printer")
}
