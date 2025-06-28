package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(51, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

}

func sum(a int, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("O soma maior do que 50")
	}
	return a + b, nil
}
