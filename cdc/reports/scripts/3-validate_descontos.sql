-- totais
select sum(valor_transacoes) as total_valor_transacoes,
      sum(quantidade_transacoes) as total_quantidade_transacoes,
      round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket,
      round(avg(taxa_desconto_media), 2) as avg_taxa_desconto_media,
      round(avg(taxa_desconto_minima), 2) as avg_taxa_desconto_minima,
      round(avg(taxa_desconto_maxima), 2) as avg_taxa_desconto_maxima,
      round(avg(desvio_padrao_taxa_desconto), 2) as avg_desvio_padrao_taxa,
      round(sum(valor_transacoes * taxa_desconto_media / 100), 2) as valor_descontado
  from descontos;

-- duplicidades - ok
select ano, trimestre, funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento, count(1)
  from descontos
group by 1, 2, 3, 4, 5, 6, 7
having count(1) > 1
limit 100000;

-- null - ok
select *
  from descontos
where not (descontos is not null)
limit 100000;

-- ano - 2025 - ok
select *
  from descontos
where ano != 2025
limit 100000;

-- trimstre - 4 - ok
select *
  from descontos
where trimestre != 4
limit 100000;

-- valores negativos ou zero - ok
select *
  from descontos
where valor_transacoes <= 0
limit 100000;

-- taxa desconto_media negativa ou zero - ok
select *
  from descontos
where taxa_desconto_media <= 0
limit 100000;

-- quantidade transacoes negativa ou zero - ok
select *
  from descontos
where quantidade_transacoes  <= 0
limit 100000;

-- taxa minima negativa - ok
select *
  from descontos
where taxa_desconto_minima < 0
limit 100000;

-- taxa desconto minima maior que 8
select count(1)
  from descontos
where taxa_desconto_minima > 8
limit 100000;

-- taxa desconto maxima maior que 10
select count(1)
  from descontos
where taxa_desconto_maxima > 10
limit 100000;

-- desvio padrao negativo - ok
select count(1)
  from descontos
where desvio_padrao_taxa_desconto < 0
limit 100000;

-- round
select count(1)
  from descontos
where desvio_padrao_taxa_desconto = 0
  and not (round(taxa_desconto_minima,1) = round(taxa_desconto_media,1) and round(taxa_desconto_media,1) = round(taxa_desconto_maxima,1))
limit 100000;

-- outliers 8 desvios
select taxa_desconto_minima, taxa_desconto_maxima, taxa_desconto_media, desvio_padrao_taxa_desconto,
       taxa_desconto_media - 8 * desvio_padrao_taxa_desconto as limite_inferior,
       taxa_desconto_media + 8 * desvio_padrao_taxa_desconto as limite_superior
  from descontos
where taxa_desconto_maxima > taxa_desconto_media + 8 * desvio_padrao_taxa_desconto
  and taxa_desconto_minima < taxa_desconto_media - 8 * desvio_padrao_taxa_desconto;

-- min <= med <= max
select count(1)
  from descontos
where not (taxa_desconto_minima <= taxa_desconto_media and taxa_desconto_media <= taxa_desconto_maxima)
limit 100000;

-- funcao inválida
select count(1)
  from descontos
where coalesce(funcao, 'X') not in ('C', 'D', 'E') 
limit 100000;


-- C	199670698.89	2729760	73.15	3.10	2.44	4.10	0.37
-- D	288827670.53	7870175	36.70	1.60	0.39	3.11	0.39
select funcao,
sum(valor_transacoes) as total_valor_transacoes,
      sum(quantidade_transacoes) as total_quantidade_transacoes,
      round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket,
      round(avg(taxa_desconto_media), 2) as avg_taxa_desconto_media,
      round(avg(taxa_desconto_minima), 2) as avg_taxa_desconto_minima,
      round(avg(taxa_desconto_maxima), 2) as avg_taxa_desconto_maxima,
      round(avg(desvio_padrao_taxa_desconto), 2) as avg_desvio_padrao_taxa 
  from descontos
group by 1
order by 2 desc
limit 100000;


-- check bandeira - 0 linhas
select *
  from descontos
where bandeira is null or bandeira not in (1,2,3,4,5,6,7,8,99)
limit 100000;

-- 1	191458388.99	3608765	53.05	2.71	2.03	3.86	0.39
-- 2	236254047.12	5806973	40.68	2.73	2.03	3.93	0.40
-- 8	 60785933.31	1184197	51.33	3.32	2.47	4.13	0.31
select bandeira,
sum(valor_transacoes) as total_valor_transacoes,
      sum(quantidade_transacoes) as total_quantidade_transacoes,
      round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket,
      round(avg(taxa_desconto_media), 2) as avg_taxa_desconto_media,
      round(avg(taxa_desconto_minima), 2) as avg_taxa_desconto_minima,
      round(avg(taxa_desconto_maxima), 2) as avg_taxa_desconto_maxima,
      round(avg(desvio_padrao_taxa_desconto), 2) as avg_desvio_padrao_taxa 
  from descontos
group by 1
order by 1
limit 100000;

-- check forma_captura - 0 linhas
select count(1)
  from descontos
where forma_captura is null or forma_captura not in (1,2,3,4,5,6)
limit 100000;

-- 1	   499208.32	  15601	32.00	2.21	1.54	2.83	0.31
-- 2	209531012.48	2847093	73.59	2.96	2.24	4.11	0.40
-- 5	278468148.62	7737241	35.99	2.88	2.13	3.92	0.34
select forma_captura,
sum(valor_transacoes) as total_valor_transacoes,
      sum(quantidade_transacoes) as total_quantidade_transacoes,  
      round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket,
      round(avg(taxa_desconto_media), 2) as avg_taxa_desconto_media,
      round(avg(taxa_desconto_minima), 2) as avg_taxa_desconto_minima,
      round(avg(taxa_desconto_maxima), 2) as avg_taxa_desconto_maxima,
      round(avg(desvio_padrao_taxa_desconto), 2) as avg_desvio_padrao_taxa 
  from descontos
group by 1
order by 1
limit 100000;


select *
  from descontos
where numero_parcelas is null or numero_parcelas not in (1,2,3,4,5,6,7,8,9,10,11,12)
limit 100000;


-- 1	432575310.22	10480614	  41.27	2.10	0.67	3.63
-- 2	 15716637.69	   64629	 243.18	3.14	2.47	4.62
-- 3	 14717796.52	   31512	 467.05	3.12	2.50	4.46
-- 4	  7419807.84	    9925	 747.59	3.10	2.56	4.14
-- 5	  4495385.52	    4253	1056.99	3.01	2.57	3.83
-- 6	  5591935.40	    4748	1177.75	3.05	2.63	3.81
-- 7	   499108.01	     334	1494.34	3.41	3.14	3.88
-- 8	  1139532.60	     691	1649.11	3.32	2.96	3.95
-- 9	   152200.11	      66	2306.06	3.37	3.20	3.50
-- 10	  4582181.19	    2462	1861.16	3.42	2.93	4.28
-- 11	    26438.63	      16	1652.41	3.72	3.69	3.75
-- 12	  1582035.69	     685	2309.54	3.24	3.01	3.65
select numero_parcelas,
sum(valor_transacoes) as total_valor_transacoes,
      sum(quantidade_transacoes) as total_quantidade_transacoes,
      round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket,
      round(avg(taxa_desconto_media), 2) as avg_taxa_desconto_media,
      round(avg(taxa_desconto_minima), 2) as avg_taxa_desconto_minima,
      round(avg(taxa_desconto_maxima), 2) as avg_taxa_desconto_maxima,
      round(avg(desvio_padrao_taxa_desconto), 2) as avg_desvio_padrao_taxa 
  from descontos
group by 1
order by 1
limit 100000;


select *
  from descontos
where codigo_segmento is null or not (codigo_segmento >= 401 and codigo_segmento <= 428);

-- 
select codigo_segmento,
sum(valor_transacoes) as total_valor_transacoes,
      sum(quantidade_transacoes) as total_quantidade_transacoes,
      round(sum(valor_transacoes) / sum(quantidade_transacoes), 2) as ticket,
      round(avg(taxa_desconto_media), 2) as avg_taxa_desconto_media,
      round(avg(taxa_desconto_minima), 2) as avg_taxa_desconto_minima,
      round(avg(taxa_desconto_maxima), 2) as avg_taxa_desconto_maxima,
      round(avg(desvio_padrao_taxa_desconto), 2) as avg_desvio_padrao_taxa 
  from descontos
group by 1
order by 2 desc
limit 100000;