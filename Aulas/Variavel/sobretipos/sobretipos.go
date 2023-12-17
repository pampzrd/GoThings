package main

import "fmt"

var x int
var y float64

func main() {
	x = 10
	y = 10.8
	fmt.Println("X = ", x)
	fmt.Println("Y = ", y)
	fmt.Printf("x = %d, tipo %T e y = %g, tipo %T", x, x, y, y)
}
