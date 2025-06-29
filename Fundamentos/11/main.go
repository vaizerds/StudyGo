package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	Wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	Wesley.Ativo = false

	//fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", Wesley.Nome, Wesley.Idade, Wesley.Ativo)

	fmt.Println(Wesley.Ativo)
}
