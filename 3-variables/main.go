package main

import "fmt"

var a string = "Hello"
var b bool
var c int

type myType struct {
	MyString string
	MyInt int
}

func main() {

	var d string
	var e int = 3

	var f = 3.0

	g := myType{"Hello", 33}

	fmt.Printf("%T: %v\n", a, a)
	fmt.Printf("%T: %v\n", b, b)
	fmt.Printf("%T: %v\n", c, c)
	fmt.Printf("%T: %v\n", d, d)
	fmt.Printf("%T: %v\n", e, e)
	fmt.Printf("%T: %v\n", f, f)

	fmt.Printf("%T: %v\n", g, g)
}
