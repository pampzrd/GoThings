package main

import (
	"fmt"
	"time"
)

func pingar(c chan string, c1 chan string) {
	for i := 0; ; i++ {
		c <- "ping"
		c1 <- "pong"
	}
}

func imprimir(c chan string, c1 chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
		msg1 := <-c1
		fmt.Println(msg1)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	var c1 chan string = make(chan string)

	go pingar(c, c1)
	go imprimir(c, c1)

	var entrada string
	fmt.Scanln(&entrada)
}
