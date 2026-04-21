-- amounts

create table cadoc_6334_v2.tmp_conccred as
select extract(year from a.transaction_date) as year,
       extract(quarter from a.transaction_date) as quarter,
       case when a.transaction_product = 'CR' then 'C'
            when a.transaction_product = 'DB' then 'D'
            else null end as function,
       case when a.transaction_brand = 'V' then 1
            when a.transaction_brand = 'M' then 2
            when a.transaction_brand = 'E' then 8
            else null end as brand,
       round(sum(transaction_amount), 2) as transaction_amount,
       count(1) as transaction_quantity
  from transaction_v4.transaction a
left join apoio.segmentos b
  on a.establishment_mcc >= b.mcc_init
  and a.establishment_mcc <= b.mcc_end
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-01-06'
  and a.status_id = 2
  group by 1, 2, 3, 4;


select count(1) as qtd_registros
  from cadoc_6334_v2.tmp_conccred;


select count(1) as qtd_registros
  from cadoc_6334_v2.conccred;


select count(1) 
  from cadoc_6334_v2.conccred a
left join cadoc_6334_v2.tmp_conccred b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.function = b.function
  and a.brand = b.brand
where b.year is null;


select count(1) 
  from cadoc_6334_v2.tmp_conccred a
left join cadoc_6334_v2.conccred b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.function = b.function
  and a.brand = b.brand
where b.year is null;

select min(a.transaction_amount - b.transaction_amount) as min_diff,
       max(a.transaction_amount - b.transaction_amount) as max_diff,
       min(a.transaction_quantity - b.transaction_quantity) as min_diff_qtd,
       max(a.transaction_quantity - b.transaction_quantity) as max_diff_qtd
  from cadoc_6334_v2.tmp_conccred a
inner join cadoc_6334_v2.conccred b
on a.year = b.year
and a.quarter = b.quarter
and a.function = b.function
and a.brand = b.brand;


select * from cadoc_6334_v2.conccred


select sum(establishment_total_quantity)
  from cadoc_6334_v2.infresta;