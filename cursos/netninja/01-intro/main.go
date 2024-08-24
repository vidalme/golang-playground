package main

import "fmt"

func main() {
	fmt.Println("oi")

	var name1 string = "andre"
	var name2 = "vidal"
	var name3 string
	name4 := "Matilda"

	fmt.Println(name1, name2, name3, name4)

	// ints
	var age1 int = 5
	var age2 = 20
	age3 := 30
	age4 := 46

	fmt.Println(age1, age2, age3)

	// bits & memory
	// -126 a 126 (arredondei)
	var num1 int8 = 25
	var num2 int16 = 30000
	// 0 a 255
	var num3 uint8 = 255

	fmt.Println(num1, num2, num3)

	// decimals
	var float1 float32 = 22.33
	var float2 float64 = 33.2222223333344444455555
	fmt.Println(float1, float2)

	fmt.Print("Meu nome é ", name1, " ", name2, " e eu tenho ", age4, "\n")
	fmt.Printf("Minha filha é a %v e ela tem %v anos \n\n", name3, age1)

	fmt.Printf("Meu nome é %v e eu tenho %v anos \n", name1, age4)
	fmt.Printf("Meu nome é %q e eu tenho %v anos \n", name1, age4)
	fmt.Printf("O meu nome é do tipo %T e minha idade do tipo %T \n", name1, age4)
	fmt.Printf("Exemplo de um numero float -> %f \n", 3.22)
	fmt.Printf("Exemplo de um numero float porem com casas decimais limitadas-> %0.3f \n", 3.14)
}
