package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	SSH_SERVICE      = "sshd"
	PASSWORD_FAILURE = "Failed password"
	THRESHOLD        = 10
)

// Estrutura do que será necessário extrair do arquivo para análise
type authoLoginInfo struct {
	sourceIP   string
	sourcePort string
	count      int
	username   string
}

// Checa se houve muitas tentativas de login e se é ou não possível brute force.
func bruteDetect(filename string) {
	ipToFailure := map[string]authoLoginInfo{} // Cria uma forma que irá receber os dados da struct
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR could not open file", err)
		os.Exit(-1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for { // O for vai pegar linha por linha do arquivo e checar a condição
		line, err := reader.ReadString('\n')
		info := parseLog(line)
		if value, ok := ipToFailure[info.sourceIP]; ok { // Se houver IP repetido ele vai armazenar no mapa junto a quantidade
			value.count += 1
			ipToFailure[info.sourceIP] = value
		} else {
			ipToFailure[info.sourceIP] = info // Info retorna a linha
		}
		if err == io.EOF { // Se ocorrer erro de leitura no arquivo por falta de mais inputs, ele interromperá
			break
		}
	}
	for ip, failure := range ipToFailure { // Confere se há mais tentativas do que o limite estabelecido e determina se houve ou não suspeita de brute force
		if failure.count >= THRESHOLD {
			fmt.Println("Possible brute force detection from IP %s with count %d", ip, failure.count)
		}
	}
}

// Deserializa e faz análise do arquivo
func parseLog(logLine string) authoLoginInfo {
	val := authoLoginInfo{}
	strLog := string(logLine) // Converte para string  para ser analisado
	// Verifica as palavras-chave necessárias para análise
	if strings.Contains(strLog, SSH_SERVICE) && strings.Contains(strLog, PASSWORD_FAILURE) {
		afterFor := strings.Split(strLog, "for")
		if len(afterFor) < 2 {
			fmt.Println("ERROR could not understand log format, unable to split log using character")
			os.Exit(-1)
		}
		infoS := strings.Split(afterFor[1], " ")
		if len(infoS) < 6 {
			fmt.Println("Length not enough")
		}
		// Atribui os valores obtidos nos campos da struct
		val.sourceIP = infoS[3]
		val.username = infoS[1]
		val.sourcePort = infoS[5]
	}
	return val // Retorna todos os dados inseridos.
}

func main() {
	if len(os.Args) < 2 { // Checa os comandos main.go sample.log  (Tem dois argumentos: o arquivo script e o arquivo log)
		fmt.Println("ERROR please provide log file path as argument")
		return
	}
	bruteDetect(os.Args[1]) // insere o argumento 1 que deve ser o arquivo ou caminho do arquivo a ser analisado.
}
