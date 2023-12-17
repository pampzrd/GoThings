package main

import (
	"io"
	"log"
	"os"
)

func main() {
	//io.WriteString escreve uma string
	if _, err := io.WriteString(os.Stdout, "Hello, Mars!"); err != nil {
		log.Fatal(err)
	}
}
