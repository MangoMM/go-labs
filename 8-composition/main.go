package main

import "fmt"

type Pooper struct{}

func (p Pooper) Poop() {
	fmt.Println("Pooping")
}

type Barker struct{}

func (b Barker) Bark() {
	fmt.Println("Woof")
}

type Sleeper struct{}

func (s Sleeper) Sleep() {
	fmt.Println("zzz")
}

type Meower struct{}

func (m Meower) Meow() {
	fmt.Println("meow")
}

type Dog struct {
	Barker
	Pooper
	Sleeper
}

type Cat struct {
	Meower
	Pooper
	Sleeper
}

type RobotDog struct {
	Barker
}



type Bird struct {
	Pooper
}

func main() {
	dog := Dog{}

	fmt.Println("---dog---")
	dog.Poop()
	dog.Sleep()
	dog.Bark()

	fmt.Println("---robot dog---")

	rd := RobotDog{}
	rd.Bark()
}
