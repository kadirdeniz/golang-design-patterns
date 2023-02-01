package main

type ihuman interface {
	setName(name string)
	getName() string
}

type human struct {
	name string
}

func (h *human) setName(name string) {
	h.name = name
}

func (h *human) getName() string {
	return h.name
}
