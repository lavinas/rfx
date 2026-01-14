
-- 487784416.61	10575164	4560930.44
select sum(valor_transacoes) as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes,
       round(sum(tarifa_intercambio * valor_transacoes / 100), 2) as custo_intercambio
  from intercam;

-- duplicate check - 0 linhas
select ano, trimestre, produto, modalidade_cartao, funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento, count(1)
  from intercam
group by 1, 2, 3, 4, 5, 6, 7, 8, 9
having count(1) > 1
limit 100000;

-- check nulls - 0 linhas
select *
  from intercam
where not (intercam is not null)
limit 100000;

-- check ano
select *
  from intercam
where ano != 2025
limit 100000;

select *
  from intercam
where trimestre != 4
limit 100000;


-- valor-transacoes - 0 linhas
select *
  from intercam
where valor_transacoes <= 0
limit 100000;

-- tarifa_intercambio - 0 linhas
select *
  from intercam
where tarifa_intercambio <= 0 or tarifa_intercambio > 9
limit 100000;

-- quantidade_transacoes - 0 linhas
select *
  from intercam
where quantidade_transacoes <= 0
limit 100000;


-- valida produto
select *
  from intercam
where produto < 31 or produto > 38 or produto is null;

-- 33	159199481.78	4337656	36.70	0.0300	2.3957	1.58	0.57
-- 31	134875399.09	2974842	45.34	0.0200	5.6300	1.31	0.62
-- 32	 57801547.75	1227263	47.10	0.0500	2.4300	1.44	0.52
-- 34	 64357995.94	1188583	54.15	0.0394	2.6000	1.79	0.66
-- 36	 32972906.33	 447022	73.76	0.0001	2.4900	1.54	0.66
-- 35	 34568400.31	 361734	95.56	0.0300	2.7500	1.80	0.69
-- 37	  2781917.45	  21876	127.17	0.5000	2.3255	1.70	0.60
-- 38	  1226767.96	  16188	75.78	0.4900	2.2600	1.45	0.62
select produto, 
       sum(valor_transacoes) as valor_transacoes, 
       sum(quantidade_transacoes) as quantidade_transacoes,
       round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket_medio,
       min(tarifa_intercambio) as min_tarifa_intercambio,
       max(tarifa_intercambio) as max_tarifa_intercambio,
       round(avg(tarifa_intercambio), 2) as avg_tarifa_intercambio,
       round(stddev(tarifa_intercambio), 2) as stddev_tarifa_intercambio
  from intercam
group by produto
order by 3 desc;

-- valida modalidade_cartao
select *
  from intercam
where modalidade_cartao not in ('P', 'C') or modalidade_cartao is null;

-- P	481894216.11	10416719	46.26	0.0001	5.6300	1.62	0.65
-- C	  5890200.50	  158445	37.18	0.0600	2.4300	1.38	0.57
select modalidade_cartao,   
       sum(valor_transacoes) as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes,
       round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket_medio,
       min(tarifa_intercambio) as min_tarifa_intercambio,
       max(tarifa_intercambio) as max_tarifa_intercambio,
       round(avg(tarifa_intercambio), 2) as avg_tarifa_intercambio,
       round(stddev(tarifa_intercambio), 2) as stddev_tarifa_intercambio
  from intercam
group by modalidade_cartao
order by 2 desc;

-- valida funcao
select *
  from intercam
where funcao not in ('C', 'D', 'E') or funcao is null;

-- D	288055173.18	7844708	36.72	0.55	0.0001	1.6000	0.17
-- C	199729243.43	2730456	73.15	1.78	0.0200	5.6300	0.48
select funcao,
      sum(valor_transacoes) as valor_transacoes,
      sum(quantidade_transacoes) as quantidade_transacoes,
      round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket_medio,
      round(avg(tarifa_intercambio), 2) as avg_tarifa_intercambio,
      min(tarifa_intercambio) as min_tarifa_intercambio,
      max(tarifa_intercambio) as max_tarifa_intercambio,
      round(stddev(tarifa_intercambio), 2) as stddev_tarifa_intercambio
  from intercam
group by funcao
order by 2 desc;  

select *
  from intercam
where bandeira is null or bandeira not in (1,2,3,4,5,6,7,8,99)
limit 100000;

-- 1	191517033.94	3609827	53.05	1.58	0.0001	5.6300	0.61
-- 2	235469953.74	5780847	40.73	1.58	0.0200	2.7500	0.72
-- 8	60797428.93	1184490	51.33	1.56	0.5000	2.3255	0.58
select bandeira,
      sum(valor_transacoes) as valor_transacoes,
      sum(quantidade_transacoes) as quantidade_transacoes,
      round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket_medio,
      round(avg(tarifa_intercambio), 2) as avg_tarifa_intercambio,
      min(tarifa_intercambio) as min_tarifa_intercambio,
      max(tarifa_intercambio) as max_tarifa_intercambio,
      round(stddev(tarifa_intercambio), 2) as stddev_tarifa_intercambio
  from intercam
group by 1
order by 1; 


select *
  from intercam
where forma_captura is null or forma_captura not in (1,2,3,4,5,6)

-- 1	501488.48	15702	31.94	0.69	0.35	1.80	0.23
-- 2	209156617.44	2831071	73.88	1.61	0.00	2.70	0.62
-- 5	278126310.69	7728391	35.99	1.56	0.02	5.63	0.65
select forma_captura,
      sum(valor_transacoes) as valor_transacoes,
      sum(quantidade_transacoes) as quantidade_transacoes,
      round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket_medio,
      round(avg(tarifa_intercambio), 2) as avg_tarifa_intercambio,
      round(min(tarifa_intercambio), 2) as min_tarifa_intercambio,
      round(max(tarifa_intercambio), 2) as max_tarifa_intercambio,
      round(stddev(tarifa_intercambio), 2) as stddev_tarifa_intercambio
  from intercam
group by 1
order by 1; 

select *
  from intercam
where numero_parcelas is null or numero_parcelas < 1 or numero_parcelas > 12
limit 100000;

-- 1	430321373.58	10453140	41.17	1.03	0.00	5.63	0.57
-- 2	16181618.66	66281	244.14	1.84	0.14	2.50	0.32
-- 3	15102271.31	32129	470.05	1.85	0.04	2.50	0.32
-- 4	7626863.42	10130	752.90	1.83	0.03	2.50	0.39
-- 5	4613506.29	4340	1063.02	1.84	0.09	2.50	0.37
-- 6	5692955.81	4795	1187.27	1.83	0.02	2.50	0.44
-- 7	522878.01	345	1515.59	2.13	0.06	2.70	0.40
-- 8	1188839.28	706	1683.91	2.17	0.06	2.75	0.43
-- 9	156170.11	68	2296.62	2.14	0.07	2.75	0.40
-- 10	4698992.88	2512	1870.62	2.14	0.04	2.75	0.46
-- 11	31468.63	18	1748.26	1.63	0.06	2.55	0.95
-- 12	1647478.63	700	2353.54	2.08	0.03	2.75	0.57
select numero_parcelas,
      sum(valor_transacoes) as valor_transacoes,
      sum(quantidade_transacoes) as quantidade_transacoes,
      round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket_medio,
      round(avg(tarifa_intercambio), 2) as avg_tarifa_intercambio,
      round(min(tarifa_intercambio), 2) as min_tarifa_intercambio,
      round(max(tarifa_intercambio), 2) as max_tarifa_intercambio,
      round(stddev(tarifa_intercambio), 2) as stddev_tarifa_intercambio
  from intercam
group by 1
order by 1;

select *
  from intercam
where codigo_segmento is null or not (codigo_segmento >= 401 and codigo_segmento <= 428);


select codigo_segmento,
      sum(valor_transacoes) as valor_transacoes,
      sum(quantidade_transacoes) as quantidade_transacoes,
      round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket_medio,
      round(avg(tarifa_intercambio), 2) as avg_tarifa_intercambio,
      round(min(tarifa_intercambio), 2) as min_tarifa_intercambio,
      round(max(tarifa_intercambio), 2) as max_tarifa_intercambio,
      round(stddev(tarifa_intercambio), 2) as stddev_tarifa_intercambio
  from intercam
group by 1
order by 1;
