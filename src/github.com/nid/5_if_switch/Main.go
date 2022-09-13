package main

import "fmt"

func main() {
	fmt.Println()
	if true {
		fmt.Println("That's True")
	}

	statePopulations := map[string]int{
		"Ohio":    1234,
		"Florida": 8998,
	}

	if popul, ok := statePopulations["Florida"]; ok {
		fmt.Println(popul)
	}
	// fmt.Println(popul) <-- popul not available, on in the if block

	number := 50
	guess := 30

	if guess < number {
		fmt.Println("too low")
	}
	if guess > number {
		fmt.Println("too high")
	}
	if guess == number {
		fmt.Println("got it")
	}
	if guess <= number {
		fmt.Println("<=")
	}
	if guess != number {
		fmt.Println("guess is not same as number, got it")
	}
	if guess < 1 || guess > 100 {
		fmt.Println("Guess have to be between 1 and 100")
	}
	fmt.Println()

	if guess == 100 {
		fmt.Println("If ()")
	} else if guess >= 200 {
		fmt.Println("else if ()")
	} else {
		fmt.Println("else ()")
	}
}
