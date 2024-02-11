package main

import "fmt" // Declaration of import

const a = "Hello World"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("The type of E is %v", e) // Using the function of lib, to print the string interpolating the value
}
