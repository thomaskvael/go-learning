package main

import "fmt"

//Coordinates struct
type Coordinates struct {
	Lat, Long float64
}

var countryCapitalCoordinatesMap map[string]Coordinates

func main() {
	countryCapital()
	countryCapitalCoordinates()
	countryCapitalCoordinates2()
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
	//countryCapitalCoordinatesMap = make(map[string]Coordinates)
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
