package main

import (
	"io/ioutil"
	"log"
)

func main() {
	//Define uma variável com o conteúdo do arquivo
	message := []byte("Hello, Gophers!")
	//ioutil.WriteFile vai gerar e escrever no arquivo.
	err := ioutil.WriteFile("arquivoGerado", message, 0664)
	if err != nil {
		log.Fatal(err)
	}
}
