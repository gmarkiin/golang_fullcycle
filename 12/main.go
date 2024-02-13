package main

import "fmt"

type Address struct {
	street       string
	neighborhood string
	number       int
}

type Client struct {
	name    string
	age     int
	active  bool
	Address // Composition in struct
}

func main() {
	marcos := Client{
		name:   "Marcos",
		age:    24,
		active: true,
	}

	marcos.Address.street = "Rua Joaquina de Galapagos"
	marcos.Address.neighborhood = "Renato Gaucho"
	marcos.Address.number = 69

	fmt.Printf(
		"Client name: %v \nClient age: %v \nClient status: %v \n", marcos.name, marcos.age, marcos.active)
	print("----------------------------------------------------\n")

	fmt.Printf(
		"Address street: %v \nAddress neighborhood: %v \nAddress number: %v \n", marcos.Address.street, marcos.Address.neighborhood, marcos.Address.number)
}
