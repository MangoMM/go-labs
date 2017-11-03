package main

import "fmt"

const A = 10
const B = 9

func main() {

	myAdder := func(a int, b int) int {
		return a + b
	}

	mySubtracter := func(a int, b int) int {
		return a - b
	}

	fmt.Println("Adder", compute(myAdder))
	fmt.Println("Subtracter", compute(mySubtracter))
}

func compute(fn func(int, int) int) int {
	return fn(A, B)
}
