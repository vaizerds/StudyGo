package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1, 3, 46, 6, 34, 654, 7645, 534, 543, 543, 543))

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}
