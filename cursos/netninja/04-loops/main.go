package main

import "fmt"

func main() {
	// count := 10
	// for i := 0; i < count; i++ {
	// 	fmt.Println(i)
	// }
	salas := []string{"orion", "andromeda", "cassopeia", "hydra", "lynx"}

	// for i := 0; i < len(salas); i++ {
	// 	fmt.Println(salas[i])
	// }

	for index, value := range salas {
		if index == 1 {
			fmt.Printf("continuando na posição %v \n", index)
			continue
		}
		if index > 2 {
			fmt.Printf("aqui tem um break %v \n", index)
			break
		}
		fmt.Printf("interação de numero %v é a constelação %v \n", index, value)
	}

	fmt.Println("------------------------------------")

	for index, value := range salas {
		fmt.Printf("Value = %v e e direto do slice é %v \n", value, salas[index])
		value = "NOVO VALUE \n"
		fmt.Printf("Value = %v e direto do slice é %v \n", value, salas[index])
		// fmt.Printf("Tentei mudar, sera que foi? %v \n", value)
	}
}
