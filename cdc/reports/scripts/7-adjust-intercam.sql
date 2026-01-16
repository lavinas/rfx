
-- verifica proporcoes a corrigir de acordo com o balanço
select round(489_434_030.63 / sum(valor_transacoes), 15) as prop_valor_transacoes,
       round(cast(10_603_159 as NUMERIC) / cast(sum(quantidade_transacoes) as numeric), 15) as prop_quantidade_transacoes
  from reports.intercam;


-- adjusta proporcionalmente os valores de intercambio para chegar no volume do balanço
create table reports.intercam_adj as
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
       round(a.valor_transacoes * 1.003381850595934, 2) as valor_transacoes,
       round(a.quantidade_transacoes * 1.002647240269749, 0) as quantidade_transacoes
  from reports.intercam a;

-- verifica a diferença restante de quantidades como arrendondamento
select 10_603_159 - sum(a.quantidade_transacoes) as dif_quantidade_transacoes_transacoes,
       489_434_030.63 - sum(a.valor_transacoes) as dif_valor_transacoes
  from reports.intercam_adj a;


-- seleciona o maior registro para alteração sem causar impacto relevante
select id,
       quantidade_transacoes,
       valor_transacoes
  from reports.intercam_adj
order by 2 desc
limit 1;

-- atualiza o registro selecionado com os valores a arredondar
update reports.intercam_adj
   set quantidade_transacoes = quantidade_transacoes + 214,
       valor_transacoes = valor_transacoes + 0.02
 where id = 77216;


-- verifica novamente a diferença restante de quantidades como arrendondamento
select 10_603_159 - sum(a.quantidade_transacoes) as dif_quantidade_transacoes_transacoes,
       489_434_030.63 - sum(a.valor_transacoes) as dif_valor_transacoes
  from reports.intercam_adj a;