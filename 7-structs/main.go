package main

import "fmt"

type Employee struct {
	Name string
	Age int
	JobTitle string
}

type Company struct {
	Name string
	Employees []Employee
	Location string
}

func main() {

	tyk := Company{
		Name: "Tyk Technologies",
		Employees: []Employee{
			{Name: "Yaara", Age: 25, JobTitle: "Consulting Engineer"},
			{Name: "Josh", Age: 45, JobTitle: "Guitarrist"},
		},
		Location: "London",
	}

	fmt.Printf("%+v\n", tyk)

	fmt.Println("Company Name", tyk.Name)
}
