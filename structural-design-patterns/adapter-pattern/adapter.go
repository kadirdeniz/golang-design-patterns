package main

import "fmt"

type adapter struct {
	*windows
}

func (a *adapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	a.insertIntoUSBPort()
}
