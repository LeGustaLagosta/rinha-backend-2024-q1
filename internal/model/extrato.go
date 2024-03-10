package model

type Extrato struct {
	Saldo_Cliente *Saldo `json:"saldo"`
	Ultimas_Transacoes []*Transacao `json:"ultimas_transacoes"`
}