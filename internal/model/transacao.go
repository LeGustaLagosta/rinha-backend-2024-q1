package model

import (
	"time"
)

type Transacao struct {
	ID         int64 `json:"-"`
	ID_cliente int64 `json:"-"`
	Valor      int64 `json:"valor" binding:"required,gte=0"`
	Tipo       string  `json:"tipo" binding:"required,len=1,oneof=c d"`
	Descricao  string  `json:"descricao" binding:"required,max=10"`
	Data	   time.Time `json:"realizada_em" gorm:"column:data_transacao"`
}

func (Transacao) TableName() string {
    return "transacoes"
}