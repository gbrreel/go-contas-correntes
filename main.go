package main

import (
	"fmt"

	"go-contas-correntes/clientes"
	"go-contas-correntes/contas"
)

func main() {
	contaSilva := contas.ContaCorrente{
		Titular: clientes.Titular{Nome: "Silva", CPF: "000.000.000-00", Profissao: "Analista"},
		Saldo:   300,
	}
	contaGallo := contas.ContaCorrente{
		Titular: clientes.Titular{Nome: "Gallo", CPF: "111.111.111-11", Profissao: "Dev"},
		Saldo:   500,
	}

	status := contaSilva.Transferir(200, &contaGallo)

	fmt.Println(status)

	fmt.Println(contaSilva)
	fmt.Println(contaGallo)

}
