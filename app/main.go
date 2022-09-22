package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}
// o c na frente equivale ao this/self
func (c *ContaCorrente) Sacar(valor float64) string {
	if c.saldo >= valor &&  valor > 0 {
		fmt.Println("\nSacando",valor, "de", c.titular)
		c.saldo -= valor
		fmt.Println("Saldo: ", c.saldo)
		return "Saque realizado com sucesso\n"
	} else if valor < 0 {
		return "Operação inválida!\n"
	} else {
		return "\nSaldo insuficiente\n"
	}
}
func (c *ContaCorrente) Depositar(valor float64) string {
    if valor > 0 {
			fmt.Println("\nDepositando",valor, "de", c.titular)
			c.saldo += valor
			fmt.Println("Saldo: ", c.saldo)
			return "Depósito realizado com sucesso\n"
		}
		return "Operação inválida!\n"
}

func (c *ContaCorrente) Show() {
	fmt.Printf(
		"Titular: %s\n"+
			"Agência: %d\n"+
			"Conta: %d\n"+
			"Saldo: %.2f\n\n", c.titular, c.numeroAgencia, c.numeroConta, c.saldo)
}

func main() {
	// Exemplo de preenchimento de um só parâmetro
	t := ContaCorrente{titular:"Thiago"}
	// Exemplo de preenchimento completo da struct
	e:= ContaCorrente{"Erica", 123, 456,1500.50}
	t.Show()
	e.Show()
	// Exemplo de preenchimento dinamico
	/*
	var cris *ContaCorrente
	cris = new(ContaCorrente)
	cris.titular = "Cris"
	fmt.Println(*cris)
	*/
	//comparar 2 ponteiros com conteúdo igual mas sem o *, retorna false
	cris := ContaCorrente{}
	cris.titular = "Cris"
	fmt.Println(cris, "\n")

	s:= ContaCorrente{}
	s.titular = "Silvia"
	s.saldo = 1500
	s.Show()
	fmt.Println(s.Sacar(300))
	fmt.Println(s.Sacar(3000))
	fmt.Println(s.Sacar(-300))
}
