package main

import "fmt"

func main() {

	fmt.Println(multiply(10, 3))

	name, age := GetNameAndAge()
	fmt.Println("Name", name, "Age", age)
}

func multiply(x int, y int) int {
	return x * y
}

func GetNameAndAge() (string, int) {
	return "Ahmet", 35
}