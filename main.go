package main

import (
	"fmt"
	"app/models"

)

func main() {
	// Exemplo de preenchimento de um só parâmetro
	t := models.ContaCorrente{Titular:"Thiago"}
	// Exemplo de preenchimento completo da struct
	e:= models.ContaCorrente{"Erica", 123, 456,1500.50}
	//t.Show()
	e.Show()
	// Exemplo de preenchimento dinamico
	/*
	var cris *ContaCorrente
	cris = new(ContaCorrente)
	cris.titular = "Cris"
	fmt.Println(*cris)
	*/
	//comparar 2 ponteiros com conteúdo igual mas sem o *, retorna false
	cris := models.ContaCorrente{}
	cris.Titular = "Cris"
	fmt.Println(cris, "\n")

	s:= models.ContaCorrente{}
	s.Titular = "Silvia"
	s.Saldo = 1500
	s.Show()
	fmt.Println(s.Sacar(300))
	fmt.Println(s.Sacar(3000))
	fmt.Println(s.Sacar(-300))
	s.Transferir(100, t)
}
