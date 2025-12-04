package main

import (
	"fmt"

	"go-contas-correntes/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) string {
	return conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaGabriel := contas.ContaPoupanca{}

	fmt.Println(contaGabriel)
}
