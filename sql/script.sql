create table clientes (
    id serial primary key,
    limite numeric,
    saldo numeric
);

create table transacoes (
    id serial primary key,
    valor numeric,
    tipo varchar(1),
    descricao varchar(20),
    data_transacao timestamp,
    id_cliente integer references clientes(id)
);

INSERT INTO clientes (limite, saldo) VALUES (100000, 0);
INSERT INTO clientes (limite, saldo) VALUES (80000, 0);
INSERT INTO clientes (limite, saldo) VALUES (1000000, 0);
INSERT INTO clientes (limite, saldo) VALUES (10000000, 0);
INSERT INTO clientes (limite, saldo) VALUES (500000, 0);
