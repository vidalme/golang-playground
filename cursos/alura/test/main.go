package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 2

func main() {
	for {
		exibeIntro()
		o := selecaoOpcao()

		switch o {
		case 1:
			iniciandoMonitoramento()
		case 2:
			fmt.Println("opcao B")
		case 0:
			fmt.Println("opcao 0")
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
			} else {
				fmt.Println("Site", site, "está fora do ar.")
				fmt.Println("Erro ", resp.StatusCode)
			}
			fmt.Println("------------------------------")
		}
	}
}

func leSitesDoArquivo() []string {
	sites := []string{}
	// arquivo, err := os.Open("sites.txt")
	// arquivo, err := os.ReadFile("sites.txt")
	leitor := bufio.NewReader()

	if err != nil {
		fmt.Println("Aconteceu um erro ao ler o arquivo sites.txt")
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
	return sites
}
