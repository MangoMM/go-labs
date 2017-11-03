package main

import "fmt"

func main() {

	//slice := []string{}
	//
	//slice = append(slice, "first")
	//slice = append(slice, "second")
	//slice = append(slice, "third")
	//slice = append(slice, "fourth")
	//
	////fmt.Printf("Array: %+v\n", slice)
	//
	//fmt.Println("slice[1]", slice[1])

	//nilSlice()

	slices := []int{10, 20, 30, 40, 50, 60}

	for key, value := range slices {
		fmt.Println(key, value)
	}
}



func nilSlice() {
	var s []int

	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil")
	}
}
