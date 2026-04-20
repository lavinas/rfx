
create table cadoc_6334_v2.tmp_ranking as
select extract(year from a.transaction_date) as year,
       extract(quarter from a.transaction_date) as quarter,
       a.establishment_code,
       case when a.transaction_product = 'CR' then 'C'
            when a.transaction_product = 'DB' then 'D'
            else null end as function,
       case when a.transaction_brand = 'V' then 1
            when a.transaction_brand = 'M' then 2
            when a.transaction_brand = 'E' then 8
            else null end as brand,
       case when a.transaction_capture = 'CHP' then 2
            when a.transaction_capture = 'CTC' then 5
            when a.transaction_capture = 'TAR' then 1 
            else null end as capture_mode,
        a.transaction_installments as installments,
        b.segment as segment_code,
       round(sum(transaction_amount), 2) as transaction_amount,
       count(1) as transaction_quantity,
       round(avg(a.revenue_mdr_value / a.transaction_amount * 100), 2) as avg_mcc_fee
  from transaction_v4.transaction a
left join apoio.segmentos b
  on a.establishment_mcc >= b.mcc_init
  and a.establishment_mcc <= b.mcc_end
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-01-06'
  and a.status_id = 2
  group by 1, 2, 3, 4, 5, 6, 7, 8
  order by 1, 2, 3, 4, 5, 6, 7, 8;


select count(1) from cadoc_6334_v2.tmp_ranking;


select count(1) from cadoc_6334_v2.ranking;


select count(1)
  from cadoc_6334_v2.tmp_ranking a
  left join cadoc_6334_v2.ranking b
    on a.year = b.year
   and a.quarter = b.quarter
   and a.establishment_code = b.establishment_code
   and a.function = b.function
   and a.brand = b.brand
   and a.capture_mode = b.capture_mode
   and a.installments = b.installments
   and a.segment_code = b.segment_code
 where b.id is null;


 select count(1)
  from cadoc_6334_v2.ranking a
  left join cadoc_6334_v2.tmp_ranking b
    on a.year = b.year
   and a.quarter = b.quarter
   and a.establishment_code = b.establishment_code
   and a.function = b.function
   and a.brand = b.brand
   and a.capture_mode = b.capture_mode
   and a.installments = b.installments
   and a.segment_code = b.segment_code
 where b.year is null;


 select min(a.avg_mcc_fee - b.avg_mcc_fee) as min_diff_avg_mcc_fee,
        max(a.avg_mcc_fee - b.avg_mcc_fee) as max_diff_avg_mcc_fee,
        min(a.transaction_amount - b.transaction_amount) as min_diff_transaction_amount,
        max(a.transaction_amount - b.transaction_amount) as max_diff_transaction_amount,
        min(a.transaction_quantity - b.transaction_quantity) as min_diff_transaction_quantity,
        max(a.transaction_quantity - b.transaction_quantity) as max_diff_transaction_quantity
  from cadoc_6334_v2.tmp_ranking a
  inner join cadoc_6334_v2.ranking b
    on a.year = b.year
   and a.quarter = b.quarter
   and a.establishment_code = b.establishment_code
   and a.function = b.function
   and a.brand = b.brand
   and a.capture_mode = b.capture_mode
   and a.installments = b.installments
   and a.segment_code = b.segment_code;

