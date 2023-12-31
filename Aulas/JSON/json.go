package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Criando as estruturas (A dos sociais, a dos dados dos usuarios, a array de todos os dados)
// Do último para o primeiro, de baixo para cima baseado no json

type Usuarios struct {
	Usuarios []Usuario `json: "usuarios"`
}

type Usuario struct {
	Nome   string `json: "Nome"`
	Tipo   string `json: "Tipo"`
	Idade  int    `json: "Idade"`
	Social Social `json: "Social"`
}

type Social struct {
	Facebook string `json: "facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	//Para abrir o arquivo
	jsonFile, err := os.Open("usuarios.json")
	//Se ocorrer algum erro ao abrir:
	if err != nil {
		fmt.Println(err)
	}
	//se não tiver erro:
	fmt.Println("Arquivo aberto com sucesso.")
	defer jsonFile.Close()

	//Vai ler tudo que está no json
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var usuarios Usuarios
	//Descompactar  em usuarios
	json.Unmarshal(byteValue, &usuarios)

	//Imprimir os dados solicitados
	for i := 0; i < len(usuarios.Usuarios); i++ {
		fmt.Println("Usuário Tipo: " + usuarios.Usuarios[i].Tipo)
		fmt.Println("Usuário Idade: " + strconv.Itoa(usuarios.Usuarios[i].Idade))
		fmt.Println("Usuário Nome: " + usuarios.Usuarios[i].Nome)
	}
}
