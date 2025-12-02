package main

import "fmt"

type ContaCorrente struct {
	titurlar      string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaSilva := ContaCorrente{}
	contaSilva.titurlar = "Silva"
	contaSilva.saldo = 1000.0

	fmt.Println(contaSilva.saldo)

	fmt.Println(contaSilva.Sacar(-300))

	fmt.Println(contaSilva.saldo)

}
