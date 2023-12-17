package main

import "fmt"

func main() {

	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}

	//ou

	n := 2
	for n <= 20 {
		fmt.Println(n)
		n += 2
	}

	/*fmt.Println(`1
	2 Com crase ele imprime como está estilizado
	3`)*/
}
