package model

import (
	"time"
)

type Saldo struct {
	Total int64 `json:"total"`
	Data_extrato time.Time `json:"data_extrato"`
	Limite int64 `json:"limite"`
}
