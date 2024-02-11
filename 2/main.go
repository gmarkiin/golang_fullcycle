package main

const a = "Hello World"

// Apparently a constant cannot share the name with a local variable

var (
	/*
	   The variables with global scope when declared, receive "0" based on type
	   Question: if Go allows the user create your type, what values receive a user type?
	   b bool
	   c int
	   d string
	   e float64
	*/
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
)

func main() {
	// var a string -- Complete way to declare a variable
	// a:= Short tag for declare a variable and his value
	a := "X"

	println(a)
}
