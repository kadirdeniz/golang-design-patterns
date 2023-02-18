package main

type icarbuilder interface {
	setWheels()
	setSeats()
	setStructure()
	getCar() car
}

// factory pattern
func newCarBuilder(carType string) icarbuilder {
	switch carType {
	case "sports":
		return newSportCarBuilder()
	case "luxury":
		return newLuxuryCarBuilder()
	default:
		return nil
	}
}
