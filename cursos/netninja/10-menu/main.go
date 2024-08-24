package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	n, err := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	return n, err
}

func createBill() bill {

	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Escreva o nome da conta: ", reader)

	b := newBill(name)
	fmt.Println("Conta foi criada com o nome: ", name)
	return b

}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	u, _ := getInput("Selecione uma opçao: ( [a] - adiciona item, [s] - salvar conta, [g] - adiciona gorjeta)", reader)

	switch u {
	case "a":
		item, _ := getInput("Digite o item consumido: ", reader)
		preco, _ := getInput("Digite o preço do item: ", reader)

		p, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Println("Preço inválido")
			promptOptions(b)
		}

		b.addItem(item, p)
		fmt.Println("Item adicionado ", item, preco)

		promptOptions(b)

	case "s":
		tip, _ := getInput("Digite a gorjeta: ", reader)
		fmt.Println(tip)
	case "g":
		fmt.Println("escolheu g")
	default:
		fmt.Println("opção inválida...")
		promptOptions(b)
	}

	fmt.Println(u)
}

func main() {

	mybill := createBill()
	promptOptions(mybill)

	// fmt.Println(mybill)

}
