package main

import "banco/contas"

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoJean := contas.ContaCorrente{}

	PagarBoleto(&contaDoJean, 60)
}
