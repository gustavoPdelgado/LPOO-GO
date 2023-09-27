package contas

import (
	t "test/clientes"
)

type ContaPoupanca struct {
	Titular                        t.Titular
	NroAgencia, NroConta, operacao int
	saldo                          float64
}

func (c *ContaPoupanca) SaqueConta(valorSaque float64) string {
	saque := valorSaque > 0 && valorSaque < c.saldo

	if saque {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente!"
	}
}
func (c *ContaPoupanca) DepositoConta(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor de deposito incorreto", c.saldo
	}
}
func (c *ContaPoupanca) ExibeSaldo() float64 {
	return c.saldo
}
