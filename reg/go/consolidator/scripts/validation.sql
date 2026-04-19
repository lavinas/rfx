-- Active: 1766518799113@@127.0.0.1@5434@reg@cadoc_6334_v2

create table cadoc_6334_v2.tmp_desconto as
select extract(year from a.transaction_date) as year,
       extract(quarter from a.transaction_date) as quarter,
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
       round(avg(a.revenue_mdr_value / a.transaction_amount * 100), 2) as avg_mdr_fee,
       round(min(a.revenue_mdr_value / a.transaction_amount * 100), 2) as min_mdr_fee,
       round(max(a.revenue_mdr_value / a.transaction_amount * 100), 2) as max_mdr_fee,
       round(stddev(a.revenue_mdr_value / a.transaction_amount * 100), 2) as stdev_mdr_fee,
       round(stddev(a.revenue_mdr_value / a.transaction_amount * 100), 2) as sqrdiff_mdr_fee,
       round(sum(transaction_amount), 2) as transaction_amount,
       count(1) as transaction_quantity
  from transaction_v4.transaction a
left join apoio.segmentos b
  on a.establishment_mcc >= b.mcc_init
  and a.establishment_mcc <= b.mcc_end
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-01-06'
  and a.status_id = 2
  group by 1, 2, 3, 4, 5, 6, 7
  order by 1, 2, 3, 4, 5, 6, 7;

select count(1)
  from cadoc_6334_v2.tmp_desconto a
left join cadoc_6334_v2.desconto b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.function = b.function
  and a.brand = b.brand
  and a.capture_mode = b.capture_mode
  and a.installments = b.installments
  and a.segment_code = b.segment_code
where b.id is null;


select count(1)
  from cadoc_6334_v2.desconto a
left join cadoc_6334_v2.tmp_desconto b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.function = b.function
  and a.brand = b.brand
  and a.capture_mode = b.capture_mode
  and a.installments = b.installments
  and a.segment_code = b.segment_code
where b.year is null;

select max(a.avg_mdr_fee - b.avg_mdr_fee) as max_diff_avg_mdr_fee,
       min(a.avg_mdr_fee - b.avg_mdr_fee) as min_diff_avg_mdr_fee,
       max(a.min_mdr_fee - b.min_mdr_fee) as max_diff_min_mdr_fee,
       min(a.min_mdr_fee - b.min_mdr_fee) as min_diff_min_mdr_fee,
       max(a.max_mdr_fee - b.max_mdr_fee) as max_diff_max_mdr_fee,
       min(a.max_mdr_fee - b.max_mdr_fee) as min_diff_max_mdr_fee,
       max(a.stdev_mdr_fee - b.stdev_mdr_fee) as max_diff_stdev_mdr_fee,
       min(a.stdev_mdr_fee - b.stdev_mdr_fee) as min_diff_stdev_mdr_fee,
       max(a.transaction_amount - b.transaction_amount) as max_diff_transaction_amount,
       min(a.transaction_amount - b.transaction_amount) as min_diff_transaction_amount,
       max(a.transaction_quantity - b.transaction_quantity) as max_diff_transaction_quantity,
       min(a.transaction_quantity - b.transaction_quantity) as min_diff_transaction_quantity
  from cadoc_6334_v2.desconto a
inner join cadoc_6334_v2.tmp_desconto b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.function = b.function
  and a.brand = b.brand
  and a.capture_mode = b.capture_mode
  and a.installments = b.installments
  and a.segment_code = b.segment_code;


select count(1) from (
select a.stdev_mdr_fee as desconto,
       b.stdev_mdr_fee as tmp_desconto,
       a.stdev_mdr_fee - b.stdev_mdr_fee as diff,
       a.transaction_quantity
  from cadoc_6334_v2.desconto a
inner join cadoc_6334_v2.tmp_desconto b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.function = b.function
  and a.brand = b.brand
  and a.capture_mode = b.capture_mode
  and a.installments = b.installments
  and a.segment_code = b.segment_code
where abs(a.stdev_mdr_fee - b.stdev_mdr_fee) > 0.01
order by 3
) as x;


select a.stdev_mdr_fee std1,
       b.stdev_mdr_fee std2,
       a.stdev_mdr_fee - b.stdev_mdr_fee diff, 
       a.transaction_quantity qtt
  from cadoc_6334_v2.desconto a
inner join cadoc_6334_v2.tmp_desconto b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.function = b.function
  and a.brand = b.brand
  and a.capture_mode = b.capture_mode
  and a.installments = b.installments
  and a.segment_code = b.segment_code
order by 4 desc;