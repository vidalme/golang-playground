package main

import "fmt"

func mudaPointers(c string, t string, ca map[string]int16) {

}

func updateFunc(c string) {
	c = "amarelo"
}

func updatePointerFunc(c *string) {
	*c = "violeta"
}

func main() {

	cor := "verde"

	corp := &cor

	fmt.Println(cor)
	updateFunc(cor)
	fmt.Println(cor)
	updatePointerFunc(corp)
	fmt.Println(cor)
	fmt.Println("-------------------")

}
