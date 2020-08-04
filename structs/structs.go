package main

import (
	"fmt"
	"strconv"
)

//Person struct with nested struct
type Person struct {
	Name string
	Age  int
	Work
}

//Work struct
type Work struct {
	Company string
	Title   string
	Salary  int
}

func assignPerson() {
	person := Person{}
	person.Name = "Thomas"
	person.Age = 35
	person.Work.Company = "Kvael INC"
	person.Work.Title = "Developer"
	person.Work.Salary = 100000

	// Change Salary using a pointer
	personWithPointer := &person
	personWithPointer.Salary = 200000

	writeStuff(person)
}

func writeStuff(person Person) {
	fmt.Println("Test")
	fmt.Println("Name: " + person.Name)
	fmt.Println("Company: " + person.Work.Company)
	// Convert integer to Ascii (Itoa)
	fmt.Println("Salary: " + strconv.Itoa(person.Work.Salary))
}

func main() {
	assignPerson()
}
