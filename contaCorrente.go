package contas

import (
	t "test/clientes"
)

// PARA DEIXAR VISIVEL PARA OUTRAS PARTES DO CODIGO, PRIMEIRA LETRA
// DAS VARIAVEIS DEVEM SER MAIUSCULAS(Variaveis e funções)
type ContaCorrente struct {
	Titular    t.Titular
	NroAgencia int
	NroConta   int
	Saldo      float64
}

func (c *ContaCorrente) SaqueConta(valorSaque float64) string {
	saque := valorSaque > 0 && valorSaque < c.Saldo

	if saque {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente!"
	}
}
func (c *ContaCorrente) DepositoConta(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Valor de deposito incorreto", c.Saldo
	}
}
func (c *ContaCorrente) TransferenciaConta(valorTrans float64, contaDestino *ContaCorrente) bool {
	if valorTrans > c.Saldo || valorTrans <= 0 {
		return false
	}
	c.Saldo -= valorTrans
	contaDestino.DepositoConta(valorTrans)
	return true
}

func (c *ContaCorrente) ExibeSaldo() float64 {
	return c.Saldo
}
