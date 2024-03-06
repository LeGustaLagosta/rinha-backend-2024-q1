package repository

import (
	"errors"

	"gorm.io/gorm"

	"rinha/internal/model"
)

var DB *gorm.DB

func InitDB(db *gorm.DB) {
	DB = db
}

func ObterCliente(id_cliente uint64) (*model.Cliente, error) {
	var cliente model.Cliente

	err := DB.First(&cliente, "id = ?", id_cliente).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("cliente não encontrado")
	}

	return &cliente, nil
}

func InserirTransacao(transacao *model.Transacao) error {
	result := DB.Create(transacao)

	if result.Error != nil{
		return errors.New("erro ao cadastrar transação")
	}

	return nil
}
