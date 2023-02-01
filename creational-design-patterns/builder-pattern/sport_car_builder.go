package main

type sportCarBuilder struct {
	car
}

func newSportCarBuilder() *sportCarBuilder {
	return &sportCarBuilder{}
}

func (b *sportCarBuilder) setSeats() {
	b.seats = 2
}

func (b *sportCarBuilder) setStructure() {
	b.structure = "carbon fiber"
}

func (b *sportCarBuilder) setWheels() {
	b.wheels = 4
}

func (b *sportCarBuilder) getCar() car {
	return b.car
}
