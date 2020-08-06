package main

import (
	"fmt"
	"sort"
)

// When iterating thru maps with range the order is not determined (random).
// To iterate in the same order as inserted you can use a third party module
// You can also sort the result by name (string)
// See example 5

//Coordinates struct
type Coordinates struct {
	Lat, Long float64
}

var countryCapitalCoordinatesMap map[string]Coordinates

func main() {
	countryCapital()
	countryCapitalCoordinates()
	countryCapitalCoordinates2()
	countryCapitalCoordinates3()
	countryCapitalCoordinates4()
	countryCapitalCoordinates5()
}

// Example 1: Map Countries with capitals
func countryCapital() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["Norway"] = "Oslo"
	countryCapitalMap["Sweden"] = "Stockholm"
	countryCapitalMap["Denmark"] = "Copenhagen"
	countryCapitalMap["Finland"] = "Helsinki"

	for countryKey := range countryCapitalMap {
		fmt.Println("Capital of", countryKey, "is", countryCapitalMap[countryKey])
	}
}

//Example 2: Map Countries with capital coordinates inside a struct
func countryCapitalCoordinates() {
	countryCapitalCoordinatesMap = make(map[string]Coordinates)
	countryCapitalCoordinatesMap["Norway"] = Coordinates{
		59.9216697, 10.7513373,
	}
	countryCapitalCoordinatesMap["Sweden"] = Coordinates{
		59.3382617, 18.0588233,
	}
	countryCapitalCoordinatesMap["Denmark"] = Coordinates{
		55.683967, 12.5859583,
	}
	countryCapitalCoordinatesMap["Finland"] = Coordinates{
		60.1687671, 24.9373578,
	}

	fmt.Println("Coordinate of capital in Norway is:", countryCapitalCoordinatesMap["Norway"])

	for capitalCoordinatesKey := range countryCapitalCoordinatesMap {
		fmt.Println("")
		fmt.Println("Coordinates of capital in", capitalCoordinatesKey, "is", countryCapitalCoordinatesMap[capitalCoordinatesKey])
		fmt.Println("Lat coordinate of capital in", capitalCoordinatesKey, "is", countryCapitalCoordinatesMap[capitalCoordinatesKey].Lat)
		fmt.Println("Long coordinate of capital in", capitalCoordinatesKey, "is", countryCapitalCoordinatesMap[capitalCoordinatesKey].Long)
	}
}

//Example 3: Map literals

func countryCapitalCoordinates2() {
	fmt.Println("")
	fmt.Println("Example 3")
	fmt.Println("")

	var countryCapitalCoordinatesMap2 = map[string]Coordinates{
		"Norway": Coordinates{
			59.9216697, 10.7513373,
		},
		"Sweden": Coordinates{
			59.3382617, 18.0588233,
		},
		"Denmark": Coordinates{
			55.683967, 12.5859583,
		},
		"Finland": Coordinates{
			60.1687671, 24.9373578,
		},
	}

	for capitalCoordinatesKey2 := range countryCapitalCoordinatesMap2 {
		fmt.Println("")
		fmt.Println("Coordinates of capital in", capitalCoordinatesKey2, "is", countryCapitalCoordinatesMap2[capitalCoordinatesKey2])
		fmt.Println("Lat coordinate of capital in", capitalCoordinatesKey2, "is", countryCapitalCoordinatesMap2[capitalCoordinatesKey2].Lat)
		fmt.Println("Long coordinate of capital in", capitalCoordinatesKey2, "is", countryCapitalCoordinatesMap2[capitalCoordinatesKey2].Long)
	}
}

//Example 4: Map literals (shortened)
func countryCapitalCoordinates3() {
	fmt.Println("")
	fmt.Println("Example 4")
	fmt.Println("")

	var countryCapitalCoordinatesMap3 = map[string]Coordinates{
		"Norway":  {59.9216697, 10.7513373},
		"Sweden":  {59.3382617, 18.0588233},
		"Denmark": {55.683967, 12.5859583},
		"Finland": {60.1687671, 24.9373578},
	}

	for capitalCoordinatesKey3 := range countryCapitalCoordinatesMap3 {
		fmt.Println("")
		fmt.Println("Coordinates of capital in", capitalCoordinatesKey3, "is", countryCapitalCoordinatesMap3[capitalCoordinatesKey3])
		fmt.Println("Lat coordinate of capital in", capitalCoordinatesKey3, "is", countryCapitalCoordinatesMap3[capitalCoordinatesKey3].Lat)
		fmt.Println("Long coordinate of capital in", capitalCoordinatesKey3, "is", countryCapitalCoordinatesMap3[capitalCoordinatesKey3].Long)
	}
}

//Example 5: Map literals (shortened) with correct order and a slice
func countryCapitalCoordinates4() {
	fmt.Println("")
	fmt.Println("Example 5")
	fmt.Println("")

	var countryCapitalCoordinatesMap4 = map[string]Coordinates{
		"Norway":  {59.9216697, 10.7513373},
		"Sweden":  {59.3382617, 18.0588233},
		"Denmark": {55.683967, 12.5859583},
		"Finland": {60.1687671, 24.9373578},
	}

	// Create slice and append Coordinates string keys
	var keysOrder []string
	for key := range countryCapitalCoordinatesMap4 {
		keysOrder = append(keysOrder, key)
	}
	fmt.Println(keysOrder)
	// Sort slice by county name (string key)
	sort.Strings(keysOrder)

	for _, key := range keysOrder {
		fmt.Println("")
		fmt.Println("Coordinates of capital in", key, "is", countryCapitalCoordinatesMap4[key])
		fmt.Println("Lat coordinate of capital in", key, "is", countryCapitalCoordinatesMap4[key].Lat)
		fmt.Println("Long coordinate of capital in", key, "is", countryCapitalCoordinatesMap4[key].Long)
	}
}

//Example 6: Modify maps
func countryCapitalCoordinates5() {
	fmt.Println("")
	fmt.Println("Example 6")
	fmt.Println("")

	var countryCapitalCoordinatesMap5 = map[string]Coordinates{
		"Norway":  {59.9216697, 10.7513373},
		"Sweden":  {59.3382617, 18.0588233},
		"Denmark": {55.683967, 12.5859583},
		"Finland": {60.1687671, 24.9373578},
	}

	fmt.Println(countryCapitalCoordinatesMap5["Norway"])

	//Modify
	countryCapitalCoordinatesMap5["Norway"] = Coordinates{100, 200}
	fmt.Println(countryCapitalCoordinatesMap5["Norway"])

	//Check if exists
	value, exist := countryCapitalCoordinatesMap5["Norway"]
	fmt.Println("The value:", value, "exists?:", exist)

	//Delete
	delete(countryCapitalCoordinatesMap5, "Norway")
	fmt.Println(countryCapitalCoordinatesMap5["Norway"])

	//Check if exists
	value, exist = countryCapitalCoordinatesMap5["Norway"]
	fmt.Println("The value:", value, "exists?:", exist)
}
