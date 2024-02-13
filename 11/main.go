package main

import "fmt"

type Client struct { // Struct is a way to organize field but probably more complex than that
	name   string
	age    int
	active bool
}

func main() {
	marcos := Client{
		name:   "Marcos",
		age:    24,
		active: true,
	}

	fmt.Printf(
		"Client name: %v \nClient age: %v \nClient status: %v \n", marcos.name, marcos.age, marcos.active)
}
