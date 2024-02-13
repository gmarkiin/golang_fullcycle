package main

import "fmt"

type Address struct {
	street       string
	neighborhood string
	number       int
}

type Client struct {
	name   string
	age    int
	active bool
	Address
}

func (client Client) Disable() { // Methods in struct
	client.active = false
	fmt.Printf("The client %s was disabled", client.name)
}

func main() {
	marcos := Client{
		name:   "Marcos",
		age:    24,
		active: true,
	}

	marcos.Disable()
}
