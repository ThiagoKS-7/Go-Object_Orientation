package models
import "fmt"

//Maiúsculo = public
//minusculo = private
type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}
// o c na frente equivale ao this/self
func (c *ContaCorrente) Sacar(valor float64) string {
	if c.Saldo >= valor &&  valor > 0 {
		fmt.Println("\nSacando",valor, "de", c.Titular)
		c.Saldo -= valor
		fmt.Println("Saldo: ", c.Saldo)
		return "Saque realizado com sucesso\n"
	} else if valor < 0 {
		return "Operação inválida!\n"
	} else {
		return "\nSaldo insuficiente\n"
	}
}
func (c *ContaCorrente) Depositar(valor float64) string {
    if valor > 0 {
			fmt.Println("\nDepositando",valor, "de", c.Titular)
			c.Saldo += valor
			fmt.Println("Saldo: ", c.Saldo)
			return "Depósito realizado com sucesso\n"
		}
		return "Operação inválida!\n"
}

func (c *ContaCorrente) Transferir(valor float64, destino ContaCorrente) {
	c.Sacar(valor)
	destino.Depositar(valor)
}

func (c *ContaCorrente) Show() {
	fmt.Printf(
		"Titular: %s\n"+
			"Agência: %d\n"+
			"Conta: %d\n"+
			"Saldo: %.2f\n\n", c.Titular, c.NumeroAgencia, c.NumeroConta, c.Saldo)
}
