--------------------------------------------------------------------------
--- TPV e qtts por bandeira - descontos e intercambio
--------------------------------------------------------------------------

-- Elo	     60785933.31	1184197
-- Master	236254047.12	5806973
-- Visa	    191458388.99	3608765
select case bandeira 
         when 1 then 'Visa'
         when 2 then 'Master'
         when 8 then 'Elo'
         else 'Erro'
       end as bandeira, 
       sum(valor_transacoes) tpv,
       sum(quantidade_transacoes) qtd_transacoes
  from reports.descontos
group by 1
order by 1;

-- Elo	     60785933.31	1184197
-- Master	236254047.12	5806973
-- Visa	    191458388.99	3608765
select case bandeira 
         when 1 then 'Visa'
         when 2 then 'Master'
         when 8 then 'Elo'
         else 'Erro'
       end as bandeira, 
         sum(a.sum_valor_transacoes) tpv,
         sum(a.quantidade_transacoes) qtd_transacoes
 from apoio.gestao a
group by 1
order by 1;

-- Elo	     60797428.93	1184490
-- Master	235469953.74	5780847
-- Visa	    191517033.94	3609827
select case bandeira 
         when 1 then 'Visa'
         when 2 then 'Master'
         when 8 then 'Elo'
         else 'Erro'
       end as bandeira, 
       sum(valor_transacoes) tpv,
       sum(quantidade_transacoes) qtd_transacoes
  from reports.intercam
group by 1
order by 1;

-- Elo	         60797428.93	1184490
-- Master	    235469953.74	5780847
-- Visa	        191517033.94	3609827
select bandeira,
       sum(sum_valor_transacoes) tpv,
       sum(quantidade_transacoes) qtd_transacoes
  from apoio.intercambio
group by 1
order by 1;

--------------------------------------------------------------------------
--- Receita total por bandeira - descontos
--------------------------------------------------------------------------

select case bandeira 
         when 1 then 'Visa'
         when 2 then 'Master'
         when 8 then 'Elo'
         else 'Erro'
       end as bandeira, 
       round(sum(valor_transacoes * taxa_desconto_media/100)/sum(valor_transacoes), 4) receita
  from reports.descontos
group by 1
order by 1;

select  round(sum(valor_transacoes * taxa_desconto_media/100)/sum(valor_transacoes), 4) receita
  from reports.descontos;



-- Elo	    456485.35
-- Master	2313796.21
-- Visa	    1790649.89
select case bandeira 
         when 1 then 'Visa'
         when 2 then 'Master'
         when 8 then 'Elo'
         else 'Erro'
       end as bandeira, 
       round(sum(a.valor_transacoes * a.tarifa_intercambio/100)/sum(a.valor_transacoes), 4) custo_medio
 from reports.intercam a
group by 1
order by 1;


select round(sum(a.valor_transacoes * a.tarifa_intercambio/100)/sum(a.valor_transacoes), 4) custo_medio
 from reports.intercam a;