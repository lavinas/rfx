select *
  from reports.intercam;


drop table if exists reports.intercam_ch;
CREATE TABLE IF NOT EXISTS reports.intercam_ch (
    id BIGINT PRIMARY KEY NULL,
    sync_status SMALLINT NULL,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    ano SMALLINT NOT NULL,
    trimestre SMALLINT NOT NULL,
    produto SMALLINT NOT NULL,
    modalidade_cartao CHAR(1) NOT NULL,
    funcao CHAR(1) NOT NULL,
    bandeira SMALLINT NOT NULL,
    forma_captura SMALLINT NOT NULL,
    numero_parcelas SMALLINT NOT NULL,
    codigo_segmento INTEGER NOT NULL,
    tarifa_intercambio NUMERIC(7,4) NOT NULL,
    valor_transacoes NUMERIC(18,2) NOT NULL,
    quantidade_transacoes INTEGER NOT NULL
);


drop table if exists reports.intercam_tmp;
create table reports.intercam_tmp as
select funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento, 
       sum(valor_transacoes) as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
  from reports.intercam
 group by 1, 2, 3, 4, 5;

insert into reports.intercam_ch
select a.id,
       a.sync_status,
       a.created_at,
       a.updated_at,
       a.ano,
       a.trimestre,
       a.produto,
       a.modalidade_cartao,
       a.funcao,
       a.bandeira,
       a.forma_captura,
       a.numero_parcelas,
       a.codigo_segmento,
       a.tarifa_intercambio,
       round(b.valor_transacoes / b.valor_transacoes * c.valor_transacoes, 2) as valor_transacoes,
       round(b.quantidade_transacoes / b.quantidade_transacoes * c.quantidade_transacoes, 0) as quantidade_transacoes
  from reports.intercam a
inner join reports.intercam_tmp b
    on a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento
inner join reports.descontos_ch c
    on a.funcao = c.funcao
   and a.bandeira = c.bandeira
   and a.forma_captura = c.forma_captura
   and a.numero_parcelas = c.numero_parcelas
   and a.codigo_segmento = c.codigo_segmento;
