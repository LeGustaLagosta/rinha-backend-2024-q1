package model

type Transacao struct {
	ID         uint64
	ID_cliente uint64
	Valor      float32 `json:"valor"`
	Tipo       string  `json:"tipo"`
	Descricao  string  `json:"descricao"`
}
