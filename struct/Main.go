package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	p := Person{Name: "Rashmith", Age: 25}

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
