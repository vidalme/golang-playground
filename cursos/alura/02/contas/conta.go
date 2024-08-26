package contas

import "fmt"

type ContaBancaria struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (t *ContaBancaria) saque(q float64) {
	fmt.Print("helo")
	if t.saldo > q && q > 0 {
		t.saldo -= q
		fmt.Printf("Saque de %v realizado com sucesso.\n", q)
		fmt.Println("Saldo atual :", t.saldo)
	} else {
		fmt.Println("Não ha fundos para o saque")
	}
}

func (t *ContaBancaria) depositar(q float64) {
	if q > 0 {
		t.saldo += q
	}
	fmt.Printf("Deposito de %v realizado com sucesso.\n", q)
	fmt.Println("Saldo atual :", t.saldo)
}

func (t *ContaBancaria) transferir(a *ContaBancaria, q float64) {
	t.saque(q)
	a.depositar(q)
	fmt.Println("transferencia completa")
}
