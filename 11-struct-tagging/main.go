package main

import (
	"fmt"
	"encoding/json"
)

type Employee struct {
	name string `json:"first_name"`
	Age int `json:"age"`
	JobTitle string `json:"job_title"`
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
			{name: "Yaara", Age: 25, JobTitle: "Consulting Engineer"},
			{name: "Josh", Age: 45, JobTitle: "Guitarrist"},
		},
		Location: "London",
	}

	jsBytes, _ := json.MarshalIndent(tyk, "+", "  ")

	fmt.Println(string(jsBytes))
}
