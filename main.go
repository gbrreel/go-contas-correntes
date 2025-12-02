package main

import "fmt"

type ContaCorrente struct {
	titular       string
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

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito < 0 {
		c.saldo += valorDeposito
		return "Valor de depósito inválido", c.saldo
	} else {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso", c.saldo
	}
}

func (c *ContaCorrente) ransferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaSilva := ContaCorrente{titular: "Silva", saldo: 300}
	contaGallo := ContaCorrente{titular: "Gallo", saldo: 500}

	status := contaSilva.ransferir(200, &contaGallo)

	fmt.Println(status)

	fmt.Println(contaSilva)
	fmt.Println(contaGallo)

}
