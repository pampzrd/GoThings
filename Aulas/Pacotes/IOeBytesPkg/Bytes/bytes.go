package main

import (
	"bytes"
	"fmt"
)

func main() {
	//Title deixa cada inicio de palavra em maiúscula.
	fmt.Printf("%s", bytes.Title([]byte("her royal highness")))
}
