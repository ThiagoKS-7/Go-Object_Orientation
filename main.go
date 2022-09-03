package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	conta1 := ContaCorrente{titular: "Thiago", numeroAgencia: 123, numeroConta: 4343, saldo: 1500}
	fmt.Printf(
		"Titular: %s\n"+
			"Agência: %d\n"+
			"Conta: %d\n"+
			"Saldo: %.2f\n", conta1.titular, conta1.numeroAgencia, conta1.numeroConta, conta1.saldo)

	/*Exemplo de ponteiro*/
	contaPonteiro := new(ContaCorrente)
	contaPonteiro.titular = "Teste"
	fmt.Println(contaPonteiro)
}
