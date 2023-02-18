package main

import "fmt"

type jean struct {
	user user
}

func (j jean) wear() {
	fmt.Println("i am wearing jean")
	j.user.wear()
}
