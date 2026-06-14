package main

import "fmt"

func Map() {
	// membuat map dengan cara 1
	var myMap1 map[string]int
	myMap1 = map[string]int{}

	myMap1["Satu-1"] = 1

	// membuat map dengan cara 2
	myMap2 := map[string]int{}

	myMap2["Satu-2"] = 1
	myMap2["Dua-2"] = 2
	myMap2["Tiga-3"] = 3

	fmt.Println(myMap2)

	// cara menghapus value di map
	delete(myMap2, "Tiga-3")

	for key, value := range myMap2 {
		fmt.Println("Key :", key, "Value :", value)
	}
}
