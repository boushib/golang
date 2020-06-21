package main

import (
	"fmt"
)

func main() {
	x := 3 + 4i
	y := 1 + 2i
	z := x / y
	fmt.Println(real(z))
	fmt.Println(imag(z))
}
