package main

import (
	"fmt"
)
type Doctor struct {
	number     int
	actorName  string
	companions []string
	otherName  string
}

func main() {
	fmt.Println()
	aDoctor := Doctor{
		number:    21,
		actorName: "Daniele Nicastro",
		companions: []string{
			"Kein Name",
			"Der Name",
			"Weiss ned",
		},
		otherName: "Mein anderer Name",
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[2])
	fmt.Println()

	// Positional syntax, not recommended
	bDoctor := Doctor{
		21,
		"Daniele Nicastro",
		[]string{
			"Kein Name",
			"Der Name",
			"Weiss ned",
		},
		"Mein anderer Name",
	}
	fmt.Println(bDoctor)
	fmt.Println()

	// Anonymous struct
	theDoctor := struct {
		name string
		age  int
	}{name: "Daniele", age: 21}
	// Copy by value
	otherDoctor := theDoctor
	otherDoctor.name = "My Name"
	otherDoctor.age = 44
	fmt.Println(theDoctor)
	fmt.Println(theDoctor.name)
	fmt.Println(otherDoctor)
	fmt.Println(otherDoctor.name)
	fmt.Println()

	// Composition
	type Animal struct {
		Name string
	}
	type Bird struct {
		Animal   // <- no type required
		SpeedKMH int
		CanFly   bool
	}

	bird := Bird{}
	bird.Name = "Emu"
	bird.SpeedKMH = 122
	bird.CanFly = true
	fmt.Println(bird)
	fmt.Println(bird.Name)
	fmt.Println()
	// Literal
	otherBird := Bird{
		Animal:   Animal{Name: "Other Emu"},
		SpeedKMH: 45,
		CanFly:   false,
	}
	fmt.Println(otherBird)
	fmt.Println()
}
