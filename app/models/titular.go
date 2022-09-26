package models

import "fmt"

type Titular struct {
	Nome      string
	Cpf       string
	Profissao string
	Cep       string
}

func (t *Titular) Show() {
	fmt.Printf(
		"Nome: %s\n"+
			"Cpf: %s\n"+
			"Profissao: %s\n"+
			"Cep: %s\n\n", t.Nome, t.Cpf, t.Profissao, t.Cep)
}
