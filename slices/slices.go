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

	//Set length to 0
	LastSlice = LastSlice[:0]
	printSlice(LastSlice)

	//Set length to 2
	LastSlice = LastSlice[:2]
	printSlice(LastSlice)

	//Drop first value and set capacity to 2
	LastSlice = LastSlice[1:]
	printSlice(LastSlice)

	//Set length to capacity
	LastSlice = LastSlice[:cap(LastSlice)]
	printSlice(LastSlice)

	//Equilevant as above
	LastSlice = LastSlice[:2]
	printSlice(LastSlice)

	//Slices inside slice

	// Method 1
	ContainerSlice := [][]string{}
	row1 := []string{}
	row2 := []string{}

	row1 = append(row1, "Hello")
	row1 = append(row1, "World")
	row2 = append(row2, "Go")
	row2 = append(row2, "Moon")
	ContainerSlice = append(ContainerSlice, row1)
	ContainerSlice = append(ContainerSlice, row2)

	println(ContainerSlice[0][0])
	println(ContainerSlice[0][1])
	println(ContainerSlice[1][0])
	println(ContainerSlice[1][1])

	// Method 2
	ContainerSlice2 := [][]string{}
	ContainerSlice2Row1 := []string{"Hello", "World"}
	ContainerSlice2Row2 := []string{"Go", "Moon"}
	ContainerSlice2 = append(ContainerSlice2, ContainerSlice2Row1)
	ContainerSlice2 = append(ContainerSlice2, ContainerSlice2Row2)

	println(ContainerSlice2[0][0])
	println(ContainerSlice2[0][1])

	// Loop thru Slices

	// Single slice (with index)
	for i, value := range AnotherSlice {
		fmt.Printf("[%d] %s\n", i, value)
	}
	// Single slice (without index)
	for _, value := range AnotherSlice {
		fmt.Printf("%s\n", value)
	}
	// Single slice (without value)
	for i := range AnotherSlice {
		fmt.Printf("%s\n", AnotherSlice[i])
	}

	// 2D/Nested slice
	for i := range ContainerSlice {
		fmt.Println(ContainerSlice[i])
	}

	for ContainerSliceIndex := range ContainerSlice {
		for sliceIndex, value := range ContainerSlice[ContainerSliceIndex] {
			fmt.Printf("%d %d %s\n", ContainerSliceIndex, sliceIndex, value)
		}
	}

}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
