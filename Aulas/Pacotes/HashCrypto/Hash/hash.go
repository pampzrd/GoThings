package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	//criar a hash
	h := crc32.NewIEEE()

	//Escrever nossos dados no Hash
	h.Write([]byte("código com pacote hash"))

	//Calcular o hash
	v := h.Sum32()
	fmt.Println(v)
}
