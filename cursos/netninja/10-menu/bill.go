package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
func (b bill) format() string {

	fs := fmt.Sprintln("-----------------------------------------")
	fs += "Itens Consumidos: \n"
	fs += fmt.Sprintln("-----------------------------------------")
	var soma float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...R$%v \n", k+":", v)
		soma += v
	}

	fs += fmt.Sprintln("-----------------------------------------")
	fs += fmt.Sprintf("%-25v ...R$%v \n", "Gorjeta:", b.tip)
	fs += fmt.Sprintln("-----------------------------------------")
	fs += fmt.Sprintf("%-25v ...R$%0.2f \n", "Total:", soma+b.tip)
	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(n string, p float64) {
	b.items[n] = p
}
