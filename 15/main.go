package main

func main() {
	a := 10
	var pointer *int = &a // This * is a notation in memory of the variable and the & is used for get the address in memory
	*pointer = 20         // The var pointer using reference changed the original value of 'a'
	b := &a
	println(*b)
}
