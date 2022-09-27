package models
import "fmt"

//Maiúsculo = public
//minusculo = private
type ContaPoupanca struct {
	Titular                     Titular
	NumeroAgencia, NumeroConta  int
	Saldo                       float64
}
// o c na frente equivale ao this/self
func (c *ContaPoupanca) Sacar(valor float64) string {
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
func (c *ContaPoupanca) Depositar(valor float64) string {
    if valor > 0 {
			fmt.Println("\nDepositando",valor, "de", c.Titular)
			c.Saldo += valor
			fmt.Println("Saldo: ", c.Saldo)
			return "Depósito realizado com sucesso\n"
		}
		return "Operação inválida!\n"
}

func (c *ContaPoupanca) Transferir(valor float64, destino ContaPoupanca) {
	c.Sacar(valor)
	destino.Depositar(valor)
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.Saldo
}

func (c *ContaPoupanca) Show() {
	c.Titular.Show()
	fmt.Printf(
			"Agência: %d\n"+
			"Conta: %d\n"+
			"Saldo: %.2f\n\n", c.NumeroAgencia, c.NumeroConta, c.Saldo)
}

func (c *ContaPoupanca) PagarBoleto(valor float64) {
  c.Sacar(valor)
  fmt.Println("Boleto de %.2f pago com sucesso!\n", valor)
}
