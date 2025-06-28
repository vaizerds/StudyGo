package main

import (
	"fmt"

	"studygo/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Println(carro.Andar())

	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)
}
