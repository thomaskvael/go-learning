package main

import (
	"fmt"
)

//Arrays has a fixed size (MyStringArray has a size of 2)
func main() {
	var MyStringArray [2]string
	MyStringArray[0] = "Hello"
	MyStringArray[1] = "World"

	//Replace World with Moon using a pointer
	MyStringArrayWithPointer := &MyStringArray
	MyStringArrayWithPointer[1] = "Moon"

	fmt.Println(MyStringArray[0], MyStringArray[1])

	//Array inside array
	ContainerSlice := [2][2]string{}
	ContainerSlice[0][0] = "Hello"
	ContainerSlice[0][1] = "World"
	ContainerSlice[1][0] = "Go"
	ContainerSlice[1][1] = "Moon"

	println(ContainerSlice[0][0])
	println(ContainerSlice[0][1])
}
