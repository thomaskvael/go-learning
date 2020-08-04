package main

import "fmt"

// Slices is similar to arrays but has a dynamic size.
// Due to this, slices are more common than arrays

func main() {
	var MySlice []string
	MySlice = append(MySlice, "Hello")
	MySlice = append(MySlice, "World")

	fmt.Println(MySlice[0], MySlice[1])

	//When a slice is defined the capacity will start at 1
	AnotherSlice := []string{}
	AnotherSlice = append(AnotherSlice, "Hello")
	AnotherSlice = append(AnotherSlice, "World")
	AnotherSlice = append(AnotherSlice, "Moon")

	fmt.Println(AnotherSlice)
	fmt.Println(AnotherSlice[0], AnotherSlice[2])
	a := AnotherSlice[1:3]
	fmt.Println(a)
	printSlice(AnotherSlice)

	// Another way to define and fill a slice
	LastSlice := []string{
		"Hello",
		"World",
		"Moon",
	}

	//Slice length and capacity
	printSlice(LastSlice)
}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
