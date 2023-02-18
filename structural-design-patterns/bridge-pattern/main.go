package main

type printer interface {
	printFile()
}

type computer interface {
	print()
	setPrinter(printer)
}

func main() {

	epson := &epson{}
	hp := &hp{}

	mac := &mac{
		printer: hp,
	}

	mac.print()

	mac.setPrinter(epson)
	mac.print()

	windows := &windows{
		printer: hp,
	}

	windows.print()

	windows.setPrinter(epson)
	windows.print()

}
