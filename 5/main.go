package main

import "fmt"

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
	var myArray [3]int // Declaration of array, is needed declare the "size" of the array

	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for key, value := range myArray {
		fmt.Printf("The value of index is %d and the value is %d\n", key, value)
	}
}
