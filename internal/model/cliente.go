package model

type Cliente struct {
	ID     int64 `json:"-"`
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}
