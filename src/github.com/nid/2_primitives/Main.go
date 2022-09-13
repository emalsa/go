package main

import (
	"fmt"
)

func main() {

	// Booleans
	var b bool = true
	fmt.Println(b)

	// Zero value are always false
	var isSet bool
	fmt.Println(isSet)

	i := 1 == 1
	ii := 1 == 2
	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T", ii, ii)
	fmt.Println()

	// Numeric
	var inte int
	fmt.Println(inte) // 0

	num := 42
	fmt.Printf("%v,%T\n", num, num)

	var unsignedInteger uint32 = 3
	fmt.Printf("%v,%T\n", unsignedInteger, unsignedInteger)

	// Floating
	f := 3.14
	fmt.Printf("%v,%T\n", f, f)
	f = 1.5e5
	fmt.Printf("%v,%T\n", f, f)
	f = 1.5E5
	fmt.Printf("%v,%T\n", f, f)

	// String
	s := "Hello World!"
	fmt.Printf("%v,%T\n", s[2], s)
}
