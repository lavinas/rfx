-- Active: 1767998905059@@127.0.0.1@5435@cdc
-- cria tabela de validacao
create table ranking_tmp_validate as
select funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento, 
       sum(valor_transacoes) valor_transacoes, 
       sum(quantidade_transacoes) quantidade_transacoes,
       round(sum(valor_transacoes * taxa_desconto_media / 100) / sum(valor_transacoes) * 100, 2)
  from reports.ranking
group by 1,2,3,4,5;


-- verifica excedentes
select count(1)
  from reports.ranking_tmp_validate a
left join reports.descontos_ch b 
   on b.funcao = a.funcao
   and b.bandeira = a.bandeira
   and b.forma_captura = a.forma_captura
   and b.numero_parcelas = a.numero_parcelas
   and b.codigo_segmento = a.codigo_segmento
where b.funcao is null;

-- verifica totais


select count(1)
  from (
select a.funcao, a.bandeira, a.forma_captura, a.numero_parcelas, a.codigo_segmento, a.valor_transacoes val_desc, a.quantidade_transacoes qtd_desc,
       sum(b.valor_transacoes) val_rank, sum(b.quantidade_transacoes) qtd_rank
   from reports.descontos a
inner join reports.ranking_tmp_validate b 
  on b.funcao = a.funcao
   and b.bandeira = a.bandeira
   and b.forma_captura = a.forma_captura
   and b.numero_parcelas = a.numero_parcelas
   and b.codigo_segmento = a.codigo_segmento
group by 1,2,3,4,5,6,7
having sum(b.valor_transacoes) > a.valor_transacoes
  ) a;


select sum(a.sum_valor_transacoes), sum(a.quantidade_transacoes)
  from apoio.gestao a
inner join apoio.segmentos b on cast(a.mcc as INTEGER) between b.mcc_init and b.mcc_end
where a.funcao = 'C' and a.bandeira = 1 and a.numero_parcelas = 1 and b.segment = 401 and a.forma_captura in (-1, 6)

-- 4179605.28 71125
select sum(a.sum_valor_transacoes), sum(a.quantidade_transacoes)
  from apoio.gestao a
inner join apoio.segmentos b on cast(a.mcc as INTEGER) between b.mcc_init and b.mcc_end
where b.segment = 401;

-- 4179605.28 71125
select sum(valor_transacoes), sum(quantidade_transacoes)
  from reports.descontos
where codigo_segmento = 401;


select codigo_estabelecimento,
       sum(a.sum_valor_transacoes) as tpv
  from apoio.gestao a
inner join apoio.segmentos b on cast(a.mcc as INTEGER) between b.mcc_init and b.mcc_end
where b.segment = 401 and a.
group by 1
order by 2 desc
limit 15


drop table reports.ranking_tmp_validate;