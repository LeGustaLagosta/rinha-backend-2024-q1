package repository

import (
	"errors"
	"database/sql"
	_ "github.com/lib/pq"

	"rinha/internal/model"
)

var DB *sql.DB
func InitDB(db *sql.DB) {
	DB = db
}

func CloseDB() {
	DB.Close()
}

func ObterCliente(id_cliente int64) (*model.Cliente, error) {
	var cliente model.Cliente

	// stmt, err := DB.Prepare("select * from clientes where id = $1")
	// if err != nil {
	// 	return nil, errors.New("erro de statement")
	// }
	// defer stmt.Close()

	// err = stmt.QueryRow(id_cliente).Scan(&cliente.ID, &cliente.Limite, &cliente.Saldo)
	err := DB.QueryRow("select clientes.id, limite, sum(case when tipo = 'c' then coalesce(valor, 0) else coalesce(-valor, 0) end) saldo from clientes left join transacoes on transacoes.id_cliente = clientes.id where clientes.id = $1 group by clientes.id, limite", id_cliente).Scan(&cliente.ID, &cliente.Limite, &cliente.Saldo)
	if err != nil {
		return nil, errors.New("erro de queryrow")
	}

	return &cliente, nil
}

func InserirTransacao(transacao *model.Transacao, cliente *model.Cliente) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("insert into transacoes(id_cliente, valor, tipo, descricao, data_transacao) values ($1, $2, $3, $4, $5)")
	if err != nil {
		return errors.New("erro de statement insert")
	}
	if _, err := stmt.Exec(transacao.ID_cliente, transacao.Valor, transacao.Tipo, transacao.Descricao, transacao.Data); err != nil{
		_ = tx.Rollback()
		return errors.New("erro de exec")
	}
	stmt.Close()

	stmt, err = tx.Prepare("update clientes set saldo = $1 where id = $2")
	if err != nil {
		return errors.New("erro de statement update")
	}
	if _, err := stmt.Exec(cliente.Saldo, transacao.ID_cliente); err != nil{
		_ = tx.Rollback()
		return errors.New("erro de exec")
	}
	stmt.Close()
	
	tx.Commit()

	return nil
}

func ObterTransacoes(id_cliente int64) ([]*model.Transacao, error) {
	var transacoes []*model.Transacao

	// stmt, err := DB.Prepare("select * from transacoes where id_cliente = $1 order by data_transacao desc limit 10")
	// if err != nil {
	// 	return nil, errors.New("erro de statement")
	// }
	// defer stmt.Close()

	// rows, err := stmt.Query(id_cliente)
	rows, err := DB.Query("select * from transacoes where id_cliente = $1 order by data_transacao desc limit 10", id_cliente)
	if err != nil {
		return nil, errors.New("erro de query")
	}

	for rows.Next() {
		transacao := &model.Transacao{}
		err = rows.Scan(&transacao.ID, &transacao.Valor, &transacao.Tipo, &transacao.Descricao, &transacao.Data, &transacao.ID_cliente)
		if err != nil {
			return nil, errors.New("erro de scan")
		}
		transacoes = append(transacoes, transacao)
	}

	if len(transacoes) == 0 {
		return []*model.Transacao{}, nil
	}

	return transacoes, nil
}