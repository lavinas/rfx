-- validatar segementos possiveis - 0 linjas
set search_path to reports; 


-- validatar segementos possiveis - 0 linjas
select *
  from ranking
where codigo_segmento is null or not (codigo_segmento >= 401 and codigo_segmento <= 428)
limit 100000;


-- validar estabelecimentos por segmento
-- deveria ter 15 por segmento maiores
select codigo_segmento,
       count(distinct codigo_estabelecimento) as quantidade_estabelecimentos
  from reports.ranking
where codigo_estabelecimento != '999999'
group by 1
order by 1;


select ano, trimestre, codigo_estabelecimento, funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento,
       count(1) as quantidade_registros
  from reports.ranking
group by 1,2,3,4,5,6,7,8
having count(1) > 1


select *
  from ranking
where codigo_estabelecimento is null
  or funcao is null or funcao not in ('D', 'C', 'E')
  or bandeira is null or bandeira not in (1,2,3,4,5,6,7,8,99)
  or forma_captura is null or forma_captura not in (1,2,3,4,5,6)
  or numero_parcelas is null or numero_parcelas < 1 or numero_parcelas > 12
  or codigo_segmento is null or not (codigo_segmento >= 401 and codigo_segmento <= 428)
  or trimestre is null or trimestre != 4
  or ano is null or ano != 2025;


select *
  from ranking 
where codigo_estabelecimento < '0';



select *
  from ranking
where valor_transacoes <= 0 and quantidade_transacoes <= 0 and taxa_desconto_media <= 0;


select case codigo_estabelecimento
         when 'group200' then 'MENORES' 
         else 'MAIORES'
       end as segmento,
       sum(a.valor_transacoes),
       sum(a.quantidade_transacoes),
       round(sum(a.valor_transacoes) / sum(a.quantidade_transacoes), 2) as ticket_medio,
       round(sum(a.taxa_desconto_media * a.valor_transacoes / 100), 2) as valor_descontado
  from reports.ranking a
  group by 1
  order by 1;


select sum(tpv) from (
select codigo_estabelecimento,
       sum(a.sum_valor_transacoes) as tpv
  from apoio.gestao a
inner join apoio.segmentos b on cast(a.mcc as INTEGER) between b.mcc_init and b.mcc_end
where b.segment = 422
group by 1
order by 2 desc
limit 15
);

select sum(valor_transacoes) 
from reports.ranking
where codigo_estabelecimento != 'group200'
  and codigo_segmento = 422;


select sum(tpv) from (
select codigo_estabelecimento,
       sum(a.sum_valor_transacoes) as tpv
  from apoio.gestao a
inner join apoio.segmentos b on cast(a.mcc as INTEGER) between b.mcc_init and b.mcc_end
where b.segment = 422
group by 1
order by 2
limit 200
);

-- 40091.21
select sum(valor_transacoes) 
from reports.ranking
where codigo_estabelecimento = 'group200'
  and codigo_segmento = 422;



select * from 