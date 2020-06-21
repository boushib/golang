package main

import "fmt"

func main() {
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
}
