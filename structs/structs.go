package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Work
}

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
	writeStuff(person)
}

func writeStuff(person Person) {
	fmt.Println("Test")
	fmt.Println("Name: " + person.Name)
	fmt.Println("Company: " + person.Work.Company)
}

func main() {
	assignPerson()
}
