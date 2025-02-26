package main

import "fmt"

// Employee struct with ID, Name, and Salary fields
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// UpdateSalary updates the salary of an employee
func (e *Employee) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
}

// Display prints employee details
func (e *Employee) Display() {
	fmt.Println("ID:", e.ID, "| Name:", e.Name, "| Salary:", e.Salary)
}

func AddEmployee(emp *Employee, db map[int]*Employee) {
	if _, exists := db[emp.ID]; exists {
		fmt.Println("Employee with ID", emp.ID, "already exists.")
		return
	}
	db[emp.ID] = emp
	fmt.Println("Employee added successfully!")
}

func RemoveEmployee(id int, db map[int]*Employee) {
	if _, exists := db[id]; exists {
		delete(db, id)
		fmt.Println("Employee with ID", id, "removed successfully!")
	} else {
		fmt.Println("Employee with ID", id, "not found.")
	}
}

func main() {

	db := make(map[int]*Employee)

	AddEmployee(&Employee{ID: 101, Name: "Rashmith", Salary: 60000}, db)
	AddEmployee(&Employee{ID: 102, Name: "John", Salary: 50000}, db)
	AddEmployee(&Employee{ID: 103, Name: "Ram", Salary: 70000}, db)

	fmt.Println("\nEmployee List:")
	for _, emp := range db {
		emp.Display()
	}

	fmt.Println("\nUpdating Salary:")
	if emp, exists := db[102]; exists {
		emp.UpdateSalary(55000)
		emp.Display()
	} else {
		fmt.Println("Employee not found!")
	}

	fmt.Println("\nRemoving Employee:")
	RemoveEmployee(101, db)

	fmt.Println("\nFinal Employee List:")
	for _, emp := range db {
		emp.Display()
	}
}
