
-- create new table
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

-- create temporary table to hold total values from 
drop table if exists reports.intercam_tmp;
create table reports.intercam_tmp as
select funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento, 
       sum(valor_transacoes) as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
  from reports.intercam
 group by 1, 2, 3, 4, 5;

create table reports.intercam_tmp_2 as
select funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento, 
       sum(valor_transacoes) as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
  from reports.intercam
 group by 1, 2, 3, 4, 5;


-- generate main intersection values
truncate table reports.intercam_ch;
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
       round(a.valor_transacoes / b.valor_transacoes * c.valor_transacoes, 2) as valor_transacoes,
       round(cast(a.quantidade_transacoes as numeric) / cast(b.quantidade_transacoes as numeric) * cast (c.quantidade_transacoes as numeric), 0) as quantidade_transacoes
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


-- ver quantas faltam para completar
-- 6	2219.97	12
select count(1), sum(a.valor_transacoes), sum(a.quantidade_transacoes)
  from reports.descontos_ch a
left join reports.intercam_tmp b
    on a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento
 where b.funcao is null;

-- gerar tabela temporaria com os registros que faltam 
-- linkando com o id mais proximo da tabela intercam
-- deve-se inserir a mesma quantidade de registros que do count da query acima
drop table if exists reports.intercam_tmp_missing;
create table reports.intercam_tmp_missing as
select a.funcao, a.bandeira, a.forma_captura, a.numero_parcelas, a.codigo_segmento, 
       a.valor_transacoes,
       a.quantidade_transacoes,
       min(c.id) as intercam_id
  from reports.descontos_ch a
left join reports.intercam_tmp b
    on a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento
inner join reports.intercam c
    on a.funcao = c.funcao
   and a.bandeira = c.bandeira
   and a.forma_captura = c.forma_captura
   and a.numero_parcelas = c.numero_parcelas
where b.funcao is null
group by 1, 2, 3, 4, 5, 6, 7;

-- inserir os registros que faltam na tabela principal
insert into reports.intercam_ch
select -1 * b.id,
       b.sync_status,
       b.created_at,
       b.updated_at,
       b.ano,
       b.trimestre,
       b.produto,
       b.modalidade_cartao,
       a.funcao,
       a.bandeira,
       a.forma_captura,
       a.numero_parcelas,
       a.codigo_segmento,
       b.tarifa_intercambio,
       a.valor_transacoes,
       a.quantidade_transacoes
  from reports.intercam_tmp_missing a
inner join reports.intercam b
    on a.intercam_id = b.id;
-- apagar tabela temporaria
drop table if exists reports.intercam_tmp_missing;


-- pegar as diferenças de arredondamento
create table intercam_tmp_adjust as
select a.funcao, a.bandeira, a.forma_captura, a.numero_parcelas, a.codigo_segmento,
       a.valor_transacoes valor_desc,
       sum(b.valor_transacoes) valor_inter,
       a.quantidade_transacoes qtde_desc,
       sum(b.quantidade_transacoes) qtde_inter
  from reports.descontos_ch a
inner join reports.intercam_ch b
    on a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento
group by 1, 2, 3, 4, 5, 6, 8;

-- pegar daqueles que tem diferença, aqueles com a quantidade de transacoes maior
drop table if exists reports.intercam_tmp_adjust_2;
create table reports.intercam_tmp_adjust_2 as
select a.funcao, a.bandeira, a.forma_captura, a.numero_parcelas, a.codigo_segmento,
       a.valor_desc, a.valor_inter,
       a.qtde_desc, a.qtde_inter, max(b.quantidade_transacoes) as max_qtde
  from intercam_tmp_adjust a
inner join reports.intercam_ch b
    on a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento
where a.valor_desc <> a.valor_inter
   or a.qtde_desc <> a.qtde_inter
group by 1, 2, 3, 4, 5, 6, 7, 8, 9;


-- pegar o id correto para ajuste
drop table if exists reports.intercam_tmp_adjust_3;
create table reports.intercam_tmp_adjust_3 as
select a.funcao, a.bandeira, a.forma_captura, a.numero_parcelas, a.codigo_segmento,
       a.valor_desc, a.valor_inter,
       a.qtde_desc, a.qtde_inter, 
       a.max_qtde,
       min(b.id) as intercam_id
  from reports.intercam_tmp_adjust_2 a
inner join reports.intercam_ch b
    on a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento
 where b.quantidade_transacoes = a.max_qtde
group by 1, 2, 3, 4, 5, 6, 7, 8, 9, 10;

-- realizar o ajuste
update reports.intercam_ch a
   set valor_transacoes = a.valor_transacoes + b.valor_desc - b.valor_inter,
       quantidade_transacoes = a.quantidade_transacoes + b.qtde_desc - b.qtde_inter
  from reports.intercam_tmp_adjust_3 b
 where a.id = b.intercam_id;

-- pegar as que ficaram com quantidade menor ou igual a zero
create table intercam_tmp_adjust_4 as
select *  
  from reports.intercam_ch a
where a.quantidade_transacoes <= 0;

-- calcular o ajuste necessário na tabela de descontos_ch
create table reports,intercam_tmp_adjust_5 as
select funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento,
       sum(-1*quantidade_transacoes+1) as quantidade_transacoes
  from reports.intercam_tmp_adjust_4
 group by 1, 2, 3, 4, 5
 order by 1, 2, 3, 4, 5;

-- realizar o ajuste na tabela principal
update reports.intercam_ch a
   set quantidade_transacoes = 1
where quantidade_transacoes <= 0;

-- faz backup da tabela descontos_ch antes do ajuste final
create table reports.descontos_ch_back as
select *
  from reports.descontos_ch;

-- realizar o ajuste final na tabela descontos_ch
update descontos_ch a
   set quantidade_transacoes = a.quantidade_transacoes + b.quantidade_transacoes
  from reports.intercam_tmp_adjust_5 b
 where a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento;

-- cria tabela de validacao para intercambio
create table intercam_tmp_valida as
select funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento,
       sum(valor_transacoes) as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
  from reports.intercam_ch
 group by 1, 2, 3, 4, 5;

-- cria a tabela de validacao para descontos
create table descontos_tmp_valida as
select funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento,
       sum(valor_transacoes) as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
  from reports.descontos_ch
 group by 1, 2, 3, 4, 5;

-- faz a validacao
 select count(1)
   from reports.intercam_tmp_valida a
 inner join reports.descontos_tmp_valida b
     on a.funcao = b.funcao
    and a.bandeira = b.bandeira
    and a.forma_captura = b.forma_captura
    and a.numero_parcelas = b.numero_parcelas
    and a.codigo_segmento = b.codigo_segmento
  where a.valor_transacoes <> b.valor_transacoes
     or a.quantidade_transacoes <> b.quantidade_transacoes;

-- apaga tabelas auxiliares
drop table reports.descontos_tmp_valida;
drop table reports.intercam_tmp_adjust;
drop table reports.intercam_tmp_adjust_2;
drop table reports.intercam_tmp_adjust_3;
drop table reports.intercam_tmp_adjust_4;
drop table reports.intercam_tmp_adjust_5;
drop table reports.intercam_tmp_valida;
drop table reports.descontos_ch_back;