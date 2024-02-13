package main

import (
	"fmt"
)

func main() {
	sumOfValues := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("The sum of numbers is %v", sumOfValues)
}

func sum(numbers ...int) int { // When we use the "..." in a function param, the allows we pass how many params we want (with the same type of course)
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
