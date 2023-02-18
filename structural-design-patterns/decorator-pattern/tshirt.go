package main

import "fmt"

type tshirt struct {
	user user
}

func (t tshirt) wear() {
	fmt.Println("i am wearing tshirt")
	t.user.wear()
}
