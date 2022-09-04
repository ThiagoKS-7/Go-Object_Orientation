package models

import "fmt"

type Conta struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func ContaCorrente(titular string, numeroAgencia int, numeroConta int, saldo float64) Conta {
	conta := Conta{titular, numeroAgencia, numeroConta, saldo}
	fmt.Printf(
		"Titular: %s\n"+
			"Agência: %d\n"+
			"Conta: %d\n"+
			"Saldo: %.2f\n", conta.titular, conta.numeroAgencia, conta.numeroConta, conta.saldo)
	return conta
}

func teste(Conta) {

}
