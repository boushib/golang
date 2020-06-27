package main

import "fmt"

func main() {
	// 1:42:17
	// x := 3 + 4i
	// y := 1 + 2i
	// z := x / y
	// a := complex(5, 4)
	// fmt.Println(real(z))
	// fmt.Println(imag(z))
	// fmt.Println(a)

	// s := "Hello Go"
	// b := []byte(s)

	// fmt.Printf("%v - %T\n", b, b)
	// fmt.Printf("%v - %T", string(b), string(b))

    /*
	const (
		_  = iota // ignore first value
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)

	fileSize := 2000000000.0
	fmt.Printf("%.2f GB\n", fileSize/GB)
	numbers := []int {2, 6, 51, 56}
	fmt.Printf("- numbers: %v\n", numbers)
	/*numbers[0] = 100
	fmt.Printf("- updated numbers: %v\n", numbers)
	fmt.Printf("- numbers count: %v\n", len(numbers))
	arrayOfArrays := [][]int { {1, 2, 3}, {4, 5, 6}, {7, 8, 9} }
	fmt.Printf("Array of arrays: %v", arrayOfArrays)
     */

	// arrays
	/*
	a := [...]int { 3, 4, 5 }
	b := &a
	b[0] = 300
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	 */

	// slices
	a := []int { 1, 2, 3}
	b := []int {4, 5, 6}
	c := append(a, b...)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
}
