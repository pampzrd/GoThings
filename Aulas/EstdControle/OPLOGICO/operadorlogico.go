package main

import "fmt"

func main() {
	x := 9
	if x == 1 || x == 2 {
		fmt.Println("UP!")
	} else {
		fmt.Println("DOWN!")
	}

	y := 0

	if x == 3 && y == 0 {
		fmt.Println("v")
	} else {
		fmt.Println("f")
	}
}
