package model

type Cliente struct {
	ID     uint64
	Limite float32 `json:"limite"`
	Saldo  float32 `json:"saldo"`
}
