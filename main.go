package main

import (
	//Colocando um apelido, para usar o pacote
	"fmt"
	c "test/contas"
)

type Verificador interface {
	SaqueConta(valorDoBoleto float64) string
}

func PagarBoleto(conta Verificador, valorDoBoleto float64) {
	conta.SaqueConta(valorDoBoleto)
}

func main() {
	var contaGustavo *c.ContaCorrente
	var contaAna *c.ContaCorrente
	//Para usar composição em estruturas aninhadas(Primeira maneira de usar multiplas estruturas - tipos aninhados)
	contaGustavo = new(c.ContaCorrente)
	contaGustavo.Titular.Nome = "Gustavo"
	contaGustavo.Titular.Cpf = "123.456.789-11"
	contaGustavo.Titular.Profissao = "Analista"
	contaGustavo.NroConta = 123
	contaGustavo.NroAgencia = 1234567
	contaGustavo.Saldo = 150.0

	//Usando import para usar composiçao(Segunda maneira de usar multiplas estruturas - tipos aninhados)
	contaAna = new(c.ContaCorrente)

	contaAna.Titular.Nome = "Ana"
	contaAna.Titular.Cpf = "065.547.891.46"
	contaAna.Titular.Profissao = "Desenvolvedor"
	contaAna.NroConta = 123
	contaAna.NroAgencia = 1234567
	contaAna.Saldo = 150.0

	fmt.Println(contaGustavo)
	fmt.Println(contaAna)
	//NÃO PRECISEI USAR &(ENDEREÇO NA MEMORIA DA CONTA NO PARAMETRO DA FUNÇÃO)
	status := contaAna.TransferenciaConta(150.00, contaGustavo)
	fmt.Println(contaGustavo.ExibeSaldo())
	fmt.Println(status)
	fmt.Println(contaGustavo)
	fmt.Println(contaAna)

	//Conta Poupança
	contaCaio := c.ContaPoupanca{}
	contaMaria := c.ContaCorrente{}
	contaCaio.DepositoConta(100)
	contaMaria.DepositoConta(100)
	fmt.Println(contaCaio, contaMaria)
	PagarBoleto(&contaCaio, 50)
	fmt.Println(contaCaio)
}
