package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("X is not too terribly big at all")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

}
