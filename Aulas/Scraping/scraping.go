//Scraping Web
//usa biblioteca colly para extrair dados de um site
//ele acessa o www pelo protocolo http
//serve: análise de dados, gerar leads
//para cibersegurança: Levantamento de informações
//Para estudos.

package main

import (
	"encoding/csv" //vai gerar um arquivo CSV
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Não foi possível criar o arquivo, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			writer.Write([]string{
				el.ChildText("td:nth-child(1)"),
				el.ChildText("td:nth-child(2)"),
				el.ChildText("td:nth-child(3)"),
			})
		})
		fmt.Println("Varredura Completa")
	})
	c.Visit("https://www.w3schools.com/html/html_tables.asp")
}
