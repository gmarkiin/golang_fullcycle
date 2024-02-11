package main

import "fmt"

func main() {
	/*
		A map is data struct how associate unique keys and values (like a array)
	*/
	myMap := map[string]int{
		"Will": 1000,
		"Joe":  1000,
		"Bill": 1000,
	}
	delete(myMap, "Will") // Remove a item in map
	myMap["Juan"] = 2985  // Add a item in map
	fmt.Println(myMap)

	myMap2 := map[string]int{
		"Logan":  2134,
		"Xavier": 2312,
	}

	fmt.Println(myMap2)
}
