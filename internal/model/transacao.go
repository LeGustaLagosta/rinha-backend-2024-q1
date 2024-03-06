package model

type Transacao struct {
	ID         int64
	ID_cliente int64
	Valor      int64 `json:"valor" binding:"required,gte=0"`
	Tipo       string  `json:"tipo" binding:"required,len=1,oneof=c d"`
	Descricao  string  `json:"descricao" binding:"required,max=10"`
}

func (Transacao) TableName() string {
    return "transacoes"
}