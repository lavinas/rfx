select case bandeira 
         when 1 then 'Visa'
         when 2 then 'Master'
         when 8 then 'Elo'
         else 'Erro'
       end as bandeira, 
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes * taxa_desconto_media / 100)/sum(valor_transacoes), 4) as valor_intercambio
  from reports.descontos_ch
group by 1
order by 1;


select 
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes * taxa_desconto_media / 100)/sum(valor_transacoes), 4) as valor_intercambio
  from reports.descontos_ch;

select ano, trimestre, funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento, count(1)
  from descontos_ch
group by 1, 2, 3, 4, 5, 6, 7
having count(1) > 1
limit 100000;

select *
  from descontos_ch
where not (descontos_ch is not null)
limit 100000;

select *
  from descontos_ch
where ano != 2025 or ano is null
limit 100000;

select *
  from descontos
where trimestre != 4 or trimestre is null
limit 100000;


select *
  from descontos_ch
where valor_transacoes <= 0
limit 100000;


select *
  from descontos_ch
where taxa_desconto_media <= 0 or taxa_desconto_media is null
limit 100000;

-- quantidade transacoes negativa ou zero - ok
select *
  from descontos_ch
where quantidade_transacoes  <= 0 or quantidade_transacoes is null
limit 100000;

-- taxa minima negativa - ok
select *
  from descontos_ch
where taxa_desconto_minima < 0 or taxa_desconto_minima is null
limit 100000;

-- taxa desconto minima maior que 8
select count(1)
  from descontos_ch
where taxa_desconto_minima > 8 or taxa_desconto_minima is null
limit 100000;

-- taxa desconto maxima maior que 10
select count(1)
  from descontos_ch
where taxa_desconto_maxima > 10 or taxa_desconto_maxima is null
limit 100000;

-- desvio padrao negativo - ok
select count(1)
  from descontos_ch
where desvio_padrao_taxa_desconto < 0 or desvio_padrao_taxa_desconto is null
limit 100000;

-- round
select count(1)
  from descontos_ch
where desvio_padrao_taxa_desconto = 0
  and not (round(taxa_desconto_minima,1) = round(taxa_desconto_media,1) and round(taxa_desconto_media,1) = round(taxa_desconto_maxima,1))
limit 100000;

-- outliers 8 desvios
select taxa_desconto_minima, taxa_desconto_maxima, taxa_desconto_media, desvio_padrao_taxa_desconto,
       taxa_desconto_media - 8 * desvio_padrao_taxa_desconto as limite_inferior,
       taxa_desconto_media + 8 * desvio_padrao_taxa_desconto as limite_superior
  from descontos_ch
where taxa_desconto_maxima > taxa_desconto_media + 8 * desvio_padrao_taxa_desconto
  and taxa_desconto_minima < taxa_desconto_media - 8 * desvio_padrao_taxa_desconto;

-- min <= med <= max
select count(1)
  from descontos_ch
where not (taxa_desconto_minima <= taxa_desconto_media and taxa_desconto_media <= taxa_desconto_maxima)
limit 100000;

-- funcao inválida
select count(1)
  from descontos_ch
where coalesce(funcao, 'X') not in ('C', 'D', 'E') 
limit 100000;

select *
  from descontos_ch
where bandeira is null or bandeira not in (1,2,3,4,5,6,7,8,99)
limit 100000;

select count(1)
  from descontos_ch
where forma_captura is null or forma_captura not in (1,2,3,4,5,6)
limit 100000;

select *
  from descontos_ch
where numero_parcelas is null or numero_parcelas not in (1,2,3,4,5,6,7,8,9,10,11,12)
limit 100000;

select *
  from descontos_ch
where codigo_segmento is null or not (codigo_segmento >= 401 and codigo_segmento <= 428);
