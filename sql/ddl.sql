create table clientes (
    id serial primary key,
    limite numeric(10,2),
    saldo numeric(10,2)
);

create table transacoes (
    id serial primary key,
    valor numeric(10,2),
    tipo varchar(1),
    descricao varchar(20),
    id_cliente integer references clientes(id)
);