// Consultaremos a POKEAPI https://pokeapi.co/api/v2/pokedex/kanto
// Usaremos o verbo GET
// GET -> Obtêm a lista dos dados/registros
// POST -> Adiciona um novo registro
// DELETE -> Remove um registro
// PUT e PATCH -> Edita (Pesquisar as diferenças)
// Criar uma estrutura para cada retorno da API (Response, Pokemón, Espécie do Pokemón)
// https://tutorialedge.net/golang/consuming-restful-api-with-go/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// As estruturas vão indicar o que a função deve procurar no arquivo e atribuir para um campo.
type Response struct {
	Nome     string    `json:"name"`
	Pokemons []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	Numero  int            `json:"entry_number"`
	Especie PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Nome string `json:"name"`
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto") //Mapeia os dados da api.
	//  Pesquisar o NIL.
	if err != nil { //tratamento se houver erro.
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body) // Escaneia e obtêm os dados da api.(?)
	if err != nil {                                    //Tratamento de erro.
		log.Fatal(err)
	}

	//fmt.Println(string(responseData)) // Converte para string porque os dados brutos vem em bytes(Porém não foi necessário). (Ela converteu aqui mas no exemplo não tem a conversão string aqui.)

	var responseObject Response //Vai trabalhar com o responseData

	json.Unmarshal(responseData, &responseObject) //Vai jogar os dados obtidos de responseData dentro de responseObject
	//Imprime ?????
	fmt.Println(responseObject.Nome)
	fmt.Println(len(responseObject.Pokemons))

	for i := 0; i < len(responseObject.Pokemons); i++ {
		fmt.Println(responseObject.Pokemons[i].Especie.Nome)
	}

}

// NOTA : Deu o erro de reconhecimento com letras minúsculas nome de struct, campos e tipos.
// NOTA : O resultado do exercício da professora retornava dados brutos porque ela converteu para string antes de chamar o retorno? Usando a conversão, não alterou o retorno.
