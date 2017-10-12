package main

import "fmt"

type Employee struct {
	ID      int
	Name    string
	Address string
	Salary  float64
}

func main() {
	var p1 = Employee{}
	//var p1 Employee
	p1.Name = "P1"
	p1.Address = "1-2-6"
	p1.ID = 1
	p1.Salary = 10000
	fmt.Println(p1)
}
