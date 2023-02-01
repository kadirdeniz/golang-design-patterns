package main

import "fmt"

func main() {
	basketballPlayer := newHumanFactory("basketballer")
	footballPlayer := newHumanFactory("footballer")

	fmt.Println(basketballPlayer)
	fmt.Println(footballPlayer)
}
