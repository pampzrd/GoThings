package main

import (
	"fmt"
)

func main() {
	/*
		x := make(map[string]string)
		x["Chapolin"] = "Colorado"
		x["Bernardo"] = "E Bianca"
		fmt.Println(x)
		fmt.Println(x["Chapolin"])
	*/

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"
	fmt.Println(elemento["Li"])
}
