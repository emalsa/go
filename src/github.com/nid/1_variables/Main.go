package main

import (
	"fmt"
	"strconv"
)

var moduleScope int = 111

var (
	actorName    string = "Eva Mendes"
	doctorNumber int    = 120
)

func main() {
	var i int
	i = 10
	fmt.Println(i)

	var j int = 89
	fmt.Println(j)

	d := 100
	fmt.Println(d)

	fmt.Println(moduleScope)

	fmt.Println(actorName)
	fmt.Println(doctorNumber)

	actorName := 1

	fmt.Println(actorName)
	fmt.Println()

	// Converting
	var ii int = 423
	fmt.Printf("%v, %T\n", ii, ii)

	var jj float32
	jj = float32(ii)
	fmt.Printf("%v, %T\n", jj, jj)

	var tt string
	tt = strconv.Itoa(ii)
	fmt.Printf("%v, %T\n",tt,tt)

}
