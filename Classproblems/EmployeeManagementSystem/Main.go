package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func (e *Employee) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
}

func (e *Employee) Display() {
	fmt.Println("ID:", e.ID, "Name:", e.Name, "Salary:", e.Salary)
}

func main() {
	db := make(map[int]*Employee)
	db[1] = &Employee{45324, "Rashmith Reddy Podduturi", 60000}
	db[2] = &Employee{45325, "Raghavan", 50000}

	for _, emp := range db {
		emp.Display()
	}
	db[1].UpdateSalary(70000)
	fmt.Println(" \n Updated Salary:")
	db[1].Display()

	delete(db, 1)
	fmt.Println("\n After Deleted: ")
	for _, emp := range db {
		emp.Display()
	}
}
