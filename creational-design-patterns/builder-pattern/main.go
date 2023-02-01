package main

import "fmt"

func main() {
	sportCarBuilder := newCarBuilder("sports")
	director := newDirector(sportCarBuilder)

	sportCar := director.getCar()
	fmt.Println(sportCar)

	luxuryCarBuilder := newCarBuilder("luxury")
	director.setBuilder(luxuryCarBuilder)

	luxuryCar := director.getCar()
	fmt.Println(luxuryCar)
}
