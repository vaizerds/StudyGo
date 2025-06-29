package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O clinte %s foi desativado", c.Nome)
}

func main() {

	Wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	Wesley.Ativo = false
	Wesley.Desativar()
}
