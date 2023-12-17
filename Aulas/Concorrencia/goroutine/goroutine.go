package main

import (
	"fmt"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		//time.Sleep(time.Second * 1)
	}
}

func main() {
	go f(0)
	//Modo de deixar o programa aberto até 'apertar o enter'
	var escreva string
	fmt.Scanln(&escreva)
}
