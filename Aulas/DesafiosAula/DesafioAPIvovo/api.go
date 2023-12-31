// API Rest para os clientes da vovó com o mux
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Clientes struct {
	Id       string    `json: "id"`
	Nome     string    `json: "nome"`
	Endereco *Endereco `json: "endereco"`
	Telefone string    `json: "telefone"`
}

type Endereco struct {
	Cep    string `json: cep`
	Cidade string `json: cidade`
	Rua    string `json: rua`
}

var clientes []Clientes

func ListarClientes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(clientes)
}

func VerCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for _, item := range clientes {
		if item.Id == vars["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Clientes{})
}

func main() {
	r := mux.NewRouter()

	clientes = append(clientes, Clientes{Id: "1", Nome: "Brigitte", Endereco: &Endereco{Cep: "00000-100", Cidade: "Paris", Rua: "Rua da Torre Eiffel"}, Telefone: "9090-9090"})
	clientes = append(clientes, Clientes{Id: "2", Nome: "Chalotte", Endereco: &Endereco{Cep: "00000-010", Cidade: "Londres", Rua: "Rua do Big Ben"}, Telefone: "8080-8080"})

	r.HandleFunc("/clientes", ListarClientes).Methods("GET")
	r.HandleFunc("/clientes/{id}", VerCliente).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
