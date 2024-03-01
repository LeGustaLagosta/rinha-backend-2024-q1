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

func PersistirCadastroInicial(clientes []model.Cliente) {
    DB.AutoMigrate(&model.Cliente{}, &model.Transacao{})
	DB.Create(&clientes)
}

func ObterCliente(id_cliente int64) (*model.Cliente, error) {
	var cliente model.Cliente

	err := DB.First(&cliente, "id = ?", id_cliente).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("cliente não encontrado")
	}

	return &cliente, nil
}

func inserirTransacao(transacao model.Transacao) {

}
