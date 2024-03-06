package model

type Cliente struct {
	ID     int64
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}
