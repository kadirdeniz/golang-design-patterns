package main

func main() {
	client := &client{}

	mac := &mac{}
	client.insertLightningConnectorIntoComputer(mac)

	windows := &windows{}
	adapter := &adapter{windows}
	client.insertLightningConnectorIntoComputer(adapter)
}
