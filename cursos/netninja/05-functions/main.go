package main

import (
	"fmt"
	"math"
)

func oi(p string) {
	fmt.Printf("Oi %v \n", p)
}
func tchau(p string) {
	fmt.Printf("Tchau %v \n", p)
}

func listP(lp []string, f func(string)) {
	for _, v := range lp {
		f(v)
	}
}

func areaCirc(r float64) float64 {
	return math.Pi * r * r
}

func main() {

	pessoas := []string{"Andre", "Manu", "Matilda"}

	listP(pessoas, oi)
	listP(pessoas, tchau)

	a1 := 10.5
	a2 := 4.2
	fmt.Printf("circle 1 %0.3f e ciruclo 2 %0.3f \n", a1, a2)

}
