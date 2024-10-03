package main

import "fmt"

type details struct {
	genre string
}

type game struct {
	details
	name  string
	price float64
}

func (d details) showDetails() {
	fmt.Println("Genre:", d.genre)
}
func (g game) show() {
	fmt.Println("Name:", g.name)
	fmt.Println("Price:", g.price)
	g.showDetails()
}

func main() {
	newGame := game{
		name:  "COunter Strike",
		price: 149.99,
		details: details{
			genre: "+18",
		},
	}
	newGame.show()
	newGame.showDetails()
}
