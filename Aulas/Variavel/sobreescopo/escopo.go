package main

import "fmt"

var y = "HELLO" // É acessível porque está com escopo global

func main() {

	var resultado = mat(3, 5)
	fmt.Println(y)
	fmt.Println(resultado)
	//fmt.Println(hu) fora do escopo
}

func mat(a int, b int) int {
	return (a + b)
}

/*func numperdido() {
	var hu = 2
}
*/
