-- Active: 1767998905059@@127.0.0.1@5435@cdc
create SCHEMA bins;
SET search_path TO bins;

drop table bins;
create table bins (
bin varchar(8) primary KEY,
bandeira varchar(10),
qtty integer,
pais varchar(50),
-- inicial
produto_base varchar(100),
banco_base varchar(100),
-- intermediario
produto_intermediario varchar(100),
banco_intermediario varchar(100),
-- final
produto_final integer, 
modalidade_final varchar(1)
);
