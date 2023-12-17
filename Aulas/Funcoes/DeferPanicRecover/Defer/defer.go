package main

import "fmt"

func num1() {
	fmt.Println("1")
}

func num2() {
	fmt.Println("2")
}

func num3() {
	fmt.Println("3")
}

func num4() {
	fmt.Println("4")
}

func main() {
	num4()
	defer num2()
	num3()
	num1()

}
