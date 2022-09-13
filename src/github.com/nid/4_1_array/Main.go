package main

import "fmt"

func main() {

	grade := [...]int{2, 3, 9, 8}
	fmt.Println(grade)
	fmt.Println()

	var students [3]string
	students[0] = "Lisa"
	students[1] = "Daniele"
	//students[2] = "KeinName"
	fmt.Println(students)
	fmt.Println(students[1])
	// length of array
	fmt.Println(len(students))
	fmt.Println()

	// Array of array
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}

	// "Better"
	var identityMatrixOther [3][3]int
	identityMatrixOther[0] = [3]int{1, 0, 0}
	identityMatrixOther[0] = [3]int{0, 1, 0}
	identityMatrixOther[0] = [3]int{0, 0, 1}

	fmt.Println(identityMatrix)
	fmt.Println(identityMatrixOther)
	fmt.Println()

	// Copy an array is a copy. Not pointing to the same memory
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 99
	fmt.Println(a) // [1, 2, 3]
	fmt.Println(b) // [1, 99, 3]

	// We can use pointer
	c := &a
	c[1] = 44
	fmt.Println(a) // [1, 44, 3]
	fmt.Println(c) // [1, 44, 3]
	fmt.Println()

	// Slices
	numbersOf := []int{1, 2, 3, 4}
	fmt.Println(len(numbersOf))
	fmt.Println(cap(numbersOf))
	fmt.Println()
	// Slices are references -> copy shares the same point in memory
	sliceFirst := []int{1, 2, 3}
	sliceSecond := sliceFirst
	sliceSecond[1] = 33
	fmt.Println(sliceFirst)  // [1, 33, 3]
	fmt.Println(sliceSecond) // [1, 33, 3]
	fmt.Println()

	// Other ways
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// The first number is inclusive indexNumber
	// The second number is exclusive indexNumber
	bb := sliceA[:]   // slice of all elements
	cc := sliceA[3:]  // from 4th element to end ->  4, 5...
	dd := sliceA[:6]  // First 6 elements -> 1, 2, 3 ... 6
	ee := sliceA[3:6] // slice 4th to 6th element -> 4, 5, 6
	fmt.Println(bb)   // [1, 2, ... 10]
	fmt.Println(cc)   // [4, 5, 6 ... 10]
	fmt.Println(dd)   // [1, 2, 3 ... 6 ]
	fmt.Println(ee)   // [4, 5, 6]
	fmt.Println()

	// Make function
	makeSliceA := make([]int, 3)
	fmt.Println(makeSliceA)
	fmt.Println(len(makeSliceA))
	fmt.Println(cap(makeSliceA))
	fmt.Println()

	makeSliceB := make([]int, 3, 50)
	fmt.Println(makeSliceB)
	fmt.Println(len(makeSliceB))
	fmt.Println(cap(makeSliceB))

	// Append
	aa := []int{}
	fmt.Println(aa)
	fmt.Println(len(aa))
	fmt.Println(cap(aa))
	fmt.Println()
	aa = append(aa, 1)
	fmt.Println(aa)
	fmt.Println(len(aa))
	fmt.Println(cap(aa))
	fmt.Println()
	aa = append(aa, 2, 3, 4, 5)
	fmt.Println(aa)
	fmt.Println(len(aa))
	fmt.Println(cap(aa))
	fmt.Println()
	aa = append(aa, 6, 7, 8, 9)
	fmt.Println(aa)
	fmt.Println(len(aa))
	fmt.Println(cap(aa))
	fmt.Println()

	// Concatenate (spread operator)
	conc := []int{}
	conc = append(conc, 5)
	conc = append(conc, []int{6, 7, 8, 9}...)
	fmt.Println(conc)
	fmt.Println(len(conc))
	fmt.Println(cap(conc))
	fmt.Println()

	// Operations
	n := []int{1, 2, 3, 4, 5}
	ns := n[:len(n)-1]
	fmt.Println(n)
	fmt.Println(ns)
	fmt.Println()

	fm := []int{1, 2, 3, 4, 5}
	fmm := append(fm[:2], fm[3:]...)
	fmt.Println(fmm)
	fmt.Println()
	fmt.Println()

	na := [...]int{1, 2, 3}
	nb := na
	na[0] = 92
	fmt.Println(na) // [92, 2, 3]
	fmt.Println(nb) // [1, 2, 3]

	af := []int{1, 2, 3}
	bf := af
	af[0] = 92
	fmt.Println(af) // [92, 2, 3]
	fmt.Println(bf) // [92, 2, 3]

}
