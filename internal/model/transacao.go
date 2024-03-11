package model

import (
	"time"
)

type Transacao struct {
	ID         int64 `json:"-"`
	ID_cliente int64 `json:"-"`
	Valor      int64 `json:"valor""`
	Tipo       string  `json:"tipo""`
	Descricao  string  `json:"descricao""`
	Data	   time.Time `json:"realizada_em"`
}

func (Transacao) TableName() string {
    return "transacoes"
}