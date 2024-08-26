package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 1
const LOG_FILE = "log.txt"

func main() {
	for {
		exibeIntro()
		o := selecaoOpcao()

		switch o {
		case 1:
			iniciandoMonitoramento()
		case 2:
			imprimeLogs()
			// fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Fechando o programa.")
			os.Exit(0)
		default:
			fmt.Println("Nao conheco esse comando")
			os.Exit(-1)
		}
	}
}

func selecaoOpcao() int {
	var o int
	fmt.Scan(&o)
	return o
}

func exibeIntro() {
	fmt.Println("Escolha uma opção:")
	fmt.Println("[1] Iniciar Monitoramento")
	fmt.Println("[2] Exibir logs")
	fmt.Println("[0] Sair do programa")
}

func iniciandoMonitoramento() {
	fmt.Println("Iniciando o monitoramento")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			resp, err := http.Get(site)
			fmt.Println("Testando site", site)
			time.Sleep(delay * time.Second)

			if err != nil {
				fmt.Println("Aconteceu um erro ao tentar acessar ", site)
				fmt.Println(err)
			}

			if resp.StatusCode == 200 {
				fmt.Println("Site", site, "está online")
				registraLog(site, true)
			} else {
				fmt.Println("Site", site, "está fora do ar.")
				fmt.Println("Erro ", resp.StatusCode)
				registraLog(site, false)
			}
			fmt.Println("------------------------------")
		}
	}
}

func leSitesDoArquivo() []string {
	sites := []string{}
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Aconteceu um erro ao ler o arquivo sites.txt")
		fmt.Println(err)
	}
	// arquivo, err := os.ReadFile("sites.txt")
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	fmt.Println(sites)
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile(LOG_FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao tentar criar ou escrever no arquivo de logs")
		fmt.Println(err)
	}
	l := time.Now().Format("02/01/2006 - 15:04:05") + "---" + site + " - online = " + strconv.FormatBool(status) + "\n"
	arquivo.WriteString(l)

	arquivo.Close()
}

func imprimeLogs() {
	log, err := os.ReadFile(LOG_FILE)
	if err != nil {
		fmt.Println("Houve erro na leitura do log")
		fmt.Println(err)
	}
	// leitor := bufio.NenwReader(log)

	fmt.Println(string(log))
}
