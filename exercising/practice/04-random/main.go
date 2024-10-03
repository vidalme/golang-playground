package main

import (
	"bufio"
	"os"
)

func otherThings() {

	os_movies, _ := os.Open("movies.txt")
	defer os_movies.Close()

	scanner := bufio.NewScanner(os_movies)

	scanner.Scan()

}

func main() {
	// movies := []string{"Dunkirk", "Interstellar", "Moon", "King", "Mad Max", "Inception"}
	// // n := []int{1, 2, 3, 4, 5, 6}

	// slices.Sort(movies)

	// ins_movies1 := slices.Insert(movies, 0, "Oppenheimer", "Shutter Island")
	// ins_movies2 := slices.Insert(ins_movies1, len(ins_movies1), "Dune 1")

	// fmt.Println(movies)
	// fmt.Println(ins_movies1)
	// fmt.Println(ins_movies2)

	// slices.SortFunc(ins_movies2, func(a, b string) int {
	// 	return strings.Compare(b, a)
	// })

	// fmt.Println(ins_movies2)
	otherThings()
}
