package main

type luxuryCarBuilder struct {
	car
}

func newLuxuryCarBuilder() *luxuryCarBuilder {
	return &luxuryCarBuilder{}
}

func (l *luxuryCarBuilder) setSeats() {
	l.seats = 4
}

func (l *luxuryCarBuilder) setStructure() {
	l.structure = "aluminum"
}

func (l *luxuryCarBuilder) setWheels() {
	l.wheels = 4
}

func (l *luxuryCarBuilder) getCar() car {
	return l.car
}
