package main

import "fmt"

type skirt struct {
	user user
}

func (s skirt) wear() {
	fmt.Println("i am wearing skirt")
	s.user.wear()
}
