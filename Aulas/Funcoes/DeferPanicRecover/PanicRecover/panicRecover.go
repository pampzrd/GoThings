package main

import "fmt"

func main() {
	defer func() { //O recover vai permitir que ao invés do erro + Pânico ocorra só a string doPanic
		x := recover()
		fmt.Println(x)
	}()
	panic("Pânico")
}
