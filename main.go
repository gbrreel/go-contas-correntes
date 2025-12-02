package main

import "fmt"

type ContaCorrente struct {
	titurlar      string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaGabriel := ContaCorrente{"Gabriel", 123, 456789, 1000.0}

	fmt.Println(contaGabriel)

}
