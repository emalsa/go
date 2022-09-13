package main

import "fmt"

func main() {
	fmt.Println()
	// statesPopulation := make(map[string]int) // Other option -> then remove := assignment below
	statesPopulation := map[string]int{
		"California": 3921919,
		"Texas":      9299232,
		"Florida":    2121332,
		"Ohio":       7212123,
	}
	fmt.Println(statesPopulation)
	fmt.Println(statesPopulation["Florida"])
	statesPopulation["Georgia"] = 4294929
	fmt.Println(statesPopulation["Georgia"])
	fmt.Println()

	fmt.Println(statesPopulation["notAvailable"]) // return 0
	// ok will be false -> useful for check if key exist
	//pop, ok := statesPopulation["notAvailable"]
	_, ok := statesPopulation["notAvailable"]
	fmt.Println(ok)

	// By reference
	sp := statesPopulation
	delete(sp,"Ohio")
	fmt.Println(statesPopulation)
	fmt.Println(sp)

}
