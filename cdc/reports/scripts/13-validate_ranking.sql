-- funcao
select funcao,
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes) / sum(quantidade_transacoes),2) as ticket_medio,
       round(avg(taxa_desconto_media),2)
 from reports.ranking
 group by 1
order by 1;
select funcao,
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes) / sum(quantidade_transacoes),2) as ticket_medio,
       round(avg(taxa_desconto_media),2)
 from reports.descontos_up
 group by 1
order by 1;

-- bandeira
select bandeira,
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes) / sum(quantidade_transacoes),2) as ticket_medio,
       round(avg(taxa_desconto_media),2)
 from reports.ranking
 group by 1
order by 1;
select bandeira,
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes) / sum(quantidade_transacoes),2) as ticket_medio,
       round(avg(taxa_desconto_media),2)
 from reports.descontos_up
 group by 1
order by 1;

select * from reports.ranking;

-- forma_captura
select forma_captura,
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes) / sum(quantidade_transacoes),2) as ticket_medio,
       round(avg(taxa_desconto_media),2)
 from reports.ranking
 group by 1
order by 1;
select forma_captura,
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes) / sum(quantidade_transacoes),2) as ticket_medio,
       round(avg(taxa_desconto_media),2)
 from reports.descontos_up
 group by 1
order by 1;

-- numero_parcelas
select numero_parcelas,
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes) / sum(quantidade_transacoes),2) as ticket_medio,
       round(avg(taxa_desconto_media),2)
 from reports.ranking
 group by 1
order by 1;
select numero_parcelas,
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes) / sum(quantidade_transacoes),2) as ticket_medio,
       round(avg(taxa_desconto_media),2)
 from reports.descontos_up
 group by 1
order by 1;

-- codigo_segmento
select codigo_segmento,
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes) / sum(quantidade_transacoes),2) as ticket_medio,
       round(avg(taxa_desconto_media),2)
 from reports.ranking
 group by 1
order by 1;
select codigo_segmento,
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes) / sum(quantidade_transacoes),2) as ticket_medio,
       round(avg(taxa_desconto_media),2)
 from reports.descontos_up
 group by 1
order by 1;