package main

import "fmt"

const myConst float32 = 43.4
const (
	a = iota
	b = iota
	c = iota
)

func main() {
	const myConst int = 43
	fmt.Println(myConst)
	//var ab int16 = 27
	//fmt.Printf("%v %T\n", myConst+ab, myConst+ab)
	fmt.Println()

	const myConstOther = 32
	var be int16 = 27
	fmt.Printf("%v %T\n", myConstOther+be, myConstOther+be)
	fmt.Println()

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}
