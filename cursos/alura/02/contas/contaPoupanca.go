package contas

import (
	"alura/banco/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

// Saque implements main.verificarConta.
// func (t ContaPoupanca) Saque(valor float64) string {
// 	panic("unimplemented")
// }

func (t *ContaPoupanca) Sacar(q float64) string {
	fmt.Print("helo")
	if t.saldo > q && q > 0 {
		t.saldo -= q
		fmt.Printf("Saque de %v realizado com sucesso.\n", q)
		fmt.Println("saldo atual :", t.saldo)
	} else {
		fmt.Println("Não ha fundos para o saque")
	}
	return "oi"
}

func (t *ContaPoupanca) Depositar(q float64) {
	if q > 0 {
		t.saldo += q
	}
	fmt.Printf("Deposito de %v realizado com sucesso.\n", q)
	fmt.Println("saldo atual :", t.saldo)
}

func (t *ContaPoupanca) Transferir(a *ContaPoupanca, q float64) {
	t.Sacar(q)
	a.Depositar(q)
	fmt.Println("transferencia completa")
}

func (t *ContaPoupanca) ObterSaldo() float64 {
	return t.saldo
}
