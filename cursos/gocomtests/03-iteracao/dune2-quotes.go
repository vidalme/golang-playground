package dune2quotes

import (
	"fmt"
)

var QUOTES = []string{
	"Lead them to paradise.",
	"I am Paul Muad'Dib Atreides, Duke of Arrakis. The Hand of God be my witness, I am the Voice from the Outer World! I will lead you to PARADISE!",
	"SILENCE!",
	"You fought well, Atreides...",
	"Power over spice is power over all.",
}

func say_quote() string {
	return "oi"
}

func main() {
	q := []string{"a", "b", "c"}
	fmt.Println(q)
	fmt.Println(say_quote())
	fmt.Println("--------------")
}
