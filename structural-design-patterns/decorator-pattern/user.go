package main

import "fmt"

type user interface {
	wear()
}

type userStruct struct {
}

func (u userStruct) wear() {
	fmt.Println("user is wearing these:")
}
