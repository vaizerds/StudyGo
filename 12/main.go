package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome   string
	Idade  int
	Ativo  bool
	Adress Endereco
}

func main() {

	Wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	Wesley.Ativo = false
	Wesley.Adress.Logradouro = "São Paulo"

	//fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", Wesley.Nome, Wesley.Idade, Wesley.Ativo)

	fmt.Println(Wesley.Adress.Logradouro)
}
