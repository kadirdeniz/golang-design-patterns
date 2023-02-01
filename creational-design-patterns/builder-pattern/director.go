package main

type Director struct {
	builder icarbuilder
}

func newDirector(b icarbuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b icarbuilder) {
	d.builder = b
}

func (d *Director) getCar() car {
	d.builder.setSeats()
	d.builder.setStructure()
	d.builder.setWheels()
	return d.builder.getCar()
}
