package main

import (
	"fmt"
	"app/models"
)

func main() {
	// Exemplo de preenchimento de um só parâmetro
	thiago := models.Titular{Nome:"Thiago",Cpf: "9038943",Profissao: "Teste",	Cep: "3434343"}
	erica := models.Titular{Nome:"Erica",	Cpf: "9136945",	Profissao: "Teste2",Cep: "344554"}
	t := models.ContaCorrente{Titular:thiago}
	// Exemplo de preenchimento completo da struct
	e:= models.ContaCorrente{erica, 123, 456,1500.50}
	t.Show()
	e.Show()
	// Exemplo de preenchimento dinamico
	/*
	var cris *ContaCorrente
	cris = new(ContaCorrente)
	cris.titular = "Cris"
	fmt.Println(*cris)
	*/
	fmt.Println(e.Sacar(300))
	fmt.Println(e.Depositar(3000))
	fmt.Println(e.Sacar(-300))
	e.Transferir(100, t)
}
