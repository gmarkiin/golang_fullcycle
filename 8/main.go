package main

import (
	"fmt"
)

const MAXIMUM_SUM float64 = 100.0

func main() {

	var fristValue float64
	var secondValue float64

	fmt.Println("Type the frist number: ")
	fmt.Scan(&fristValue) // Function that's receive the input of user in CLI

	fmt.Println("Type the second number: ")
	fmt.Scan(&secondValue)

	sumOfValues, err := sum(fristValue, secondValue)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The sum of numbers is %v", sumOfValues)
}

func sum(a float64, b float64) (float64, error) { // Functions in Go can return more than one type
	if a+b >= 100 {
		return 0.0, fmt.Errorf("Your sum is more than %v", MAXIMUM_SUM)
	}
	return a + b, nil
}
