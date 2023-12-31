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
	nome    string    `json:"name"`
	pokemon []pokemon `json:"pokemon_entries"`
}

type pokemon struct {
	numero  int            `json:"entry_number"`
	especie pokemonSpecies `json:"pokemon_species"`
}

type pokemonSpecies struct {
	nome string `json:"name"`
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

	fmt.Println(string(responseData)) // Precisa converter para string porque os dados brutos vem em bytes.

	var responseObject Response //Is not a type?   //Vai trabalhar com o responseData

	json.Unmarshal(responseData, &responseObject) //Vai jogar os dados obtidos de responseData dentro de responseObject
	//Imprime ?????
	fmt.Println(responseObject.nome)
	fmt.Println(len(responseObject.pokemon))

	for _, pokemon := range responseObject.pokemon {
		fmt.Println(pokemon.especie.nome)
	}

}
