package main

import (
	"fmt"
)

func main() {
	var isReady bool = true
	isReady = !isReady
	if isReady {
		fmt.Println("I'm ready!")
	} else {
		fmt.Println("I'm not ready!")
	}
	fmt.Println(1 == 1)
}
