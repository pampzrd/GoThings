package main

import "fmt"

func main() {
	x := 0
	for {
		if x < 20 {
			fmt.Println("X menor que 20")
			x++
		} else {
			break
		}

	}

	// Continue
	for y := 0; y < 20; y++ {
		if y == 3 || y == 6 || y == 9 { //Continue se 3, Ele pula o 3. Ignora o 3.
			continue
		} else {
			fmt.Println(y)
		}
	}
}
