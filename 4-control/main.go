package main

import (
	"fmt"
	"time"
)



func main() {

	//forLoop()
	//whileLoop()
	//forever()

	doDefer()
}

func doDefer() {

	defer fmt.Println("Deferred")

	fmt.Println("Something")
}

func forever() {
	for {
		time.Sleep(time.Minute * 5)
		fmt.Println(time.Now())
	}
}

func whileLoop() {

	counter := 0
	finished := false

	for !finished {

		fmt.Println("Counter", counter)

		if counter == 5 {
			finished = true
		}
		counter ++
	}
}

func forLoop() {

	for i := 0; i < 10; i++ {
		fmt.Println("For Loop", i)
	}

}
