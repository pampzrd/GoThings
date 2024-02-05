package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Response struct {
	Created    string `json:"created_at"`
	Categories string `json:"categories"`
	Value      string `json:"value"`
}

func main() {
	for {
		var decision string
		responseapi, err := http.Get("https://api.chucknorris.io/jokes/random")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(responseapi.Body)
		if err != nil {
			log.Fatal(err)
		}
		var responseObject Response
		json.Unmarshal(responseData, &responseObject)

		fmt.Println("[Z] TO START - [X] TO END")
		fmt.Scan(&decision)
		if decision == "z" {
			fmt.Print(formatterOut(responseObject) + "\n")
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("End of Jokes")
			os.Exit(1)
		}
	}
}

func formatterOut(objeto Response) string {
	created := string(objeto.Created)
	value := string(strings.ToUpper(objeto.Value))
	result := created + "\n" + value
	return result
}
