package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Zero", i)
		case 1:
			fmt.Println("Um", i)
		case 2:
			fmt.Println("Dois", i)
		case 3:
			fmt.Println("Três", i)

		}
		fmt.Println(i)

	}

}
