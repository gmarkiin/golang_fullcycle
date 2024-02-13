package main

import "fmt"

type Address struct {
	street       string
	neighborhood string
	number       int
}

type Person interface { // I don't understand interfaces in Go, the concept is like OOP but the implementation is weird, maybe according use make`s more sense
	Disable()
}

type Client struct {
	name   string
	age    int
	active bool
	Address
}

func (client Client) Disable() {
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
