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
	u, _ := getInput("Selecione uma opçao: ( [a] - adiciona item, [g] - adiciona gorjeta, [s] - salvar conta)", reader)

	switch u {
	case "a":
		item, _ := getInput("Digite o item consumido: ", reader)
		preco, _ := getInput("Digite o preço do item: ", reader)

		p, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Println("Valor inválido")
			promptOptions(b)
		}

		b.addItem(item, p)
		fmt.Println("Item adicionado ", item, preco)

		promptOptions(b)

	case "g":
		gorjeta, _ := getInput("Digite o valor da gorjeta: ", reader)
		g, err := strconv.ParseFloat(gorjeta, 64)
		if err != nil {
			fmt.Println("Valor inválido")
		}
		b.updateTip(g)
		fmt.Println("Gorjeta adicionada: ", g)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("Voce salvou a conta", b)

	default:
		fmt.Println("opção inválida...")
		promptOptions(b)
	}
}

func main() {

	mybill := createBill()
	promptOptions(mybill)

}
