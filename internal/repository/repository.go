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

func ObterCliente(id_cliente int64) (*model.Cliente, error) {
	var cliente model.Cliente

	err := DB.First(&cliente, "id = ?", id_cliente).Error
	if err != nil {
		return nil, errors.New("cliente não encontrado")
	}

	return &cliente, nil
}

func InserirTransacao(transacao *model.Transacao, cliente *model.Cliente) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	
	
	if err := DB.Create(transacao).Error; err != nil{
		return errors.New("erro ao registrar transação")
	}

	if err := DB.Save(cliente).Error; err != nil{
		return errors.New("erro ao atualizar saldo")
	}

	return tx.Commit().Error
}

func ObterTransacoes(id_cliente int64) (*[]model.Transacao, error) {
	var transacoes []model.Transacao

	err := DB.Limit(10).Order("data_transacao desc").Where("id_cliente = ?", id_cliente).Find(&transacoes).Error
	if err != nil {
		return nil, errors.New("transações não encontradas")
	}

	return &transacoes, nil
}