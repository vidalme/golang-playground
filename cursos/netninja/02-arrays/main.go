package main

import "fmt"

func main() {

	// Arrays - cant be altered (only its values)
	// var ages [3]int = [3]int{5, 31, 46}
	var ages = [3]int{5, 30, 46}
	fmt.Println(ages, len(ages))

	ages[0] = 6
	fmt.Println(ages)

	// slices - behave a bit more like a regular array
	techs := []string{"k8s", "docker", "terraform", "bash"}
	fmt.Println(techs, len(techs))

	techs[2] = "containers"
	fmt.Println(techs, len(techs))

	techs = append(techs, "aws", "golang")
	fmt.Println(techs, len(techs))

	// slice ranges
	range1 := techs[1:3]
	range2 := techs[3:]

	fmt.Println(range1, range2)

}
