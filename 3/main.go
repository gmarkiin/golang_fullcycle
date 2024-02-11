package main

const a = "Hello World"

type ID int // Declare of a type with receives a primal type (this point of the class FOR ME don't make sense)

var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	println(f)
}
