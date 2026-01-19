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
having sum(b.valor_transacoes) > a.valor_transacoes;