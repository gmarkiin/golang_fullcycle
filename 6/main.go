package main

import "fmt"

func main() {
	/*
		A slice is basically an array without a size limit
		If your slice need more size, than declared, they duplicate the initial size
	*/
	mySlice := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}

	fmt.Printf("len=%d cap=%d %v \n", len(mySlice), cap(mySlice), mySlice)
	// Cap is a function how returns the capacity of a data
	fmt.Printf("len=%d cap=%d %v \n", len(mySlice[:0]), cap(mySlice[:0]), mySlice[:0])
	// When do you use a slice[:*], is removed all the values in right side of a slice from the value (*)

	fmt.Printf("len=%d cap=%d %v \n", len(mySlice[:4]), cap(mySlice[:4]), mySlice[:4])
	fmt.Printf("len=%d cap=%d %v \n", len(mySlice[2:]), cap(mySlice[2:]), mySlice[2:])

	mySlice = append(mySlice, 110)
	fmt.Printf("len=%d cap=%d %v \n", len(mySlice), cap(mySlice), mySlice)
}
