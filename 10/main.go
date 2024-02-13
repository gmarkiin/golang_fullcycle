package main

import (
	"fmt"
)

func main() {
	sumOfValues := func() int { // Just showing the use of closures
		return sum(1, 2, 3)
	}()
	fmt.Printf("The sum of numbers is %v", sumOfValues)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
