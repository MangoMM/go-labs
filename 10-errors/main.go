package main

import (
	"math"
	"fmt"
	"errors"
	"os"
)

func main() {

	answer, err := sqrt(-100)
	handleError(err)

	fmt.Println(answer)
}

func handleError(err error) {
	if err != nil {
		// Handle the error
		fmt.Println("Oops", err)
		os.Exit(1)
	}
}

func sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, errors.New("sqrt cannot be negative")
	}

	return math.Sqrt(x), nil
}
