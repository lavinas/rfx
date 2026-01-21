-- validatar segementos possiveis - 0 linjas
set search_path to reports; 


-- validatar segementos possiveis - 0 linjas
select *
  from ranking_ch
where codigo_segmento is null or not (codigo_segmento >= 401 and codigo_segmento <= 428)
limit 100000;


-- validar estabelecimentos por segmento
-- deveria ter 15 por segmento maiores
select codigo_segmento,
       count(distinct codigo_estabelecimento) as quantidade_estabelecimentos
  from reports.ranking_ch
where codigo_estabelecimento != 'group200'
group by 1
order by 1;

select codigo_segmento,
       count(distinct codigo_estabelecimento) as quantidade_estabelecimentos
  from reports.ranking_ch
where codigo_estabelecimento = 'group200'
group by 1
order by 1;



select ano, trimestre, codigo_estabelecimento, funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento,
       count(1) as quantidade_registros
  from reports.ranking_ch
group by 1,2,3,4,5,6,7,8
having count(1) > 1


select *
  from ranking_ch
where codigo_estabelecimento is null
  or funcao is null or funcao not in ('D', 'C', 'E')
  or bandeira is null or bandeira not in (1,2,3,4,5,6,7,8,99)
  or forma_captura is null or forma_captura not in (1,2,3,4,5,6)
  or numero_parcelas is null or numero_parcelas < 1 or numero_parcelas > 12
  or codigo_segmento is null or not (codigo_segmento >= 401 and codigo_segmento <= 428)
  or trimestre is null or trimestre != 4
  or ano is null or ano != 2025;


select *
  from ranking_ch
where codigo_estabelecimento < '0';



select *
  from ranking_ch
where valor_transacoes <= 0 and quantidade_transacoes <= 0 and taxa_desconto_media <= 0;


select case codigo_estabelecimento
         when 'group200' then 'MENORES' 
         else 'MAIORES'
       end as segmento,
       sum(a.valor_transacoes),
       sum(a.quantidade_transacoes),
       round(sum(a.valor_transacoes) / sum(a.quantidade_transacoes), 2) as ticket_medio,
       round(sum(a.taxa_desconto_media * a.valor_transacoes / 100), 2) as valor_descontado
  from reports.ranking_ch a
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


select count(1)
  from reports.ranking_ch a
left join reports.descontos_ch b 
   on b.funcao = a.funcao
   and b.bandeira = a.bandeira
   and b.forma_captura = a.forma_captura
   and b.numero_parcelas = a.numero_parcelas
   and b.codigo_segmento = a.codigo_segmento
where b.funcao is null;


select codigo_segmento
  from (
select a.funcao, a.bandeira, a.forma_captura, a.numero_parcelas, a.codigo_segmento, a.valor_transacoes val_desc, a.quantidade_transacoes qtd_desc,
       sum(b.valor_transacoes) val_rank, sum(b.quantidade_transacoes) qtd_rank
   from reports.descontos a
inner join reports.ranking_ch b 
  on b.funcao = a.funcao
   and b.bandeira = a.bandeira
   and b.forma_captura = a.forma_captura
   and b.numero_parcelas = a.numero_parcelas
   and b.codigo_segmento = a.codigo_segmento
group by 1,2,3,4,5,6,7
having sum(b.valor_transacoes) > a.valor_transacoes
) a
group by 1
order by 1;


select segmento
  from ranking.ranking_group
group by 1
having count(1) < 215
order by 1;


select a.funcao, a.bandeira, a.forma_captura, a.numero_parcelas, a.codigo_segmento, a.valor_transacoes val_desc, a.quantidade_transacoes qtd_desc,
       sum(b.valor_transacoes) val_rank, sum(b.quantidade_transacoes) qtd_rank
   from reports.descontos a
inner join reports.ranking_ch b 
  on b.funcao = a.funcao
   and b.bandeira = a.bandeira
   and b.forma_captura = a.forma_captura
   and b.numero_parcelas = a.numero_parcelas
   and b.codigo_segmento = a.codigo_segmento
   and b.codigo_estabelecimento != 'group200'
group by 1,2,3,4,5,6,7
having sum(b.valor_transacoes) > a.valor_transacoes


select a.funcao, a.bandeira, a.forma_captura, a.numero_parcelas, a.codigo_segmento, a.valor_transacoes val_desc, a.quantidade_transacoes qtd_desc,
       sum(b.valor_transacoes) val_rank, sum(b.quantidade_transacoes) qtd_rank
   from reports.descontos a
inner join reports.ranking_ch b 
  on b.funcao = a.funcao
   and b.bandeira = a.bandeira
   and b.forma_captura = a.forma_captura
   and b.numero_parcelas = a.numero_parcelas
   and b.codigo_segmento = a.codigo_segmento
   and b.codigo_estabelecimento = 'group200'
group by 1,2,3,4,5,6,7
having sum(b.valor_transacoes) > a.valor_transacoes