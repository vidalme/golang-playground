package main

import "fmt"

type trabalhando interface {
	receber()
	trabalhar(h int)
}

type colaborador struct {
	nome    string
	id      int
	idade   int
	horas   int
	conta   float64
	salario float64
}

type seguranca struct {
	colaborador
	turno string
	nivel int
}

type vendedor struct {
	colaborador
	regiao   string
	comissao float64
}

type vendedor_externo struct {
	vendedor
	carro    string
	telefone int
}

type vendedor_interno struct {
	vendedor
	sala int
}

func (c vendedor_externo) receber() {
	c.conta += c.salario
}

func (c vendedor_externo) trabalhar(h int) {
	c.horas += h
}

func (c vendedor_interno) receber() {
	c.conta += c.salario
}

func (c vendedor_interno) trabalhar(h int) {
	c.horas += h
}

func main() {

	// v1 := vendedor_interno{
	// 	vendedor: vendedor{
	// 		regiao: "nordeste",
	// 		colaborador: colaborador{
	// 			nome:  "andre",
	// 			id:    223344,
	// 			idade: 46,
	// 		},
	// 	},
	// 	sala: 11,
	// }
	// v2 := vendedor_interno{
	// 	vendedor: vendedor{
	// 		regiao: "sudeste",
	// 		colaborador: colaborador{
	// 			nome:  "manu",
	// 			id:    666444,
	// 			idade: 32,
	// 		},
	// 	},
	// 	sala: 11,
	// }
	// v3 := vendedor_externo{
	// 	vendedor: vendedor{
	// 		regiao: "norte",
	// 		colaborador: colaborador{
	// 			nome:  "fulana",
	// 			id:    222222,
	// 			idade: 19,
	// 		},
	// 	},
	// 	carro: "chevete",
	// }

	vendedores := []trabalhando{
		vendedor_interno{
			vendedor: vendedor{
				regiao: "nordeste",
				colaborador: colaborador{
					nome:  "andre",
					id:    223344,
					idade: 46,
				},
			},
			sala: 11,
		},
		vendedor_interno{
			vendedor: vendedor{
				regiao: "sudeste",
				colaborador: colaborador{
					nome:  "manu",
					id:    666444,
					idade: 32,
				},
			},
			sala: 21,
		},
		vendedor_externo{
			vendedor: vendedor{
				regiao: "norte",
				colaborador: colaborador{
					nome:  "fulana",
					id:    222222,
					idade: 19,
				},
			},
			carro: "chevete",
		},
	}

	// fmt.Println(vendedores)

	for _, i := range vendedores {
		fmt.Println(i)
	}

	// fmt.Println(v1.vendedor.colaborador.nome)
	// fmt.Println(v1.vendedor.colaborador.idade)
	// fmt.Println(v1.sala)
	// fmt.Println(v1.vendedor.regiao)
	// fmt.Println(v2.vendedor.colaborador.nome)
	// fmt.Println(v2.vendedor.colaborador.idade)
	// fmt.Println(v2.sala)
	// fmt.Println(v2.vendedor.regiao)
	// fmt.Println(v3.vendedor.colaborador.nome)
	// fmt.Println(v3.vendedor.colaborador.idade)
	// fmt.Println(v3.vendedor.regiao)
	// fmt.Println(v3.carro)
}
