-- Active: 1766518799113@@127.0.0.1@5434@reg@cadoc_6334_v2


drop table if exists cadoc_6334_v2.tmp_ranking_group;

CREATE TABLE cadoc_6334_v2.tmp_ranking_group AS
SELECT
    a.establishment_code,
    a.segment_code,   
    SUM(a.transaction_amount) as transaction_amount
FROM cadoc_6334_v2.ranking a
GROUP BY 1, 2;


CREATE TABLE cadoc_6334_v2.tmp_ranking_15 AS
SELECT
    establishment_code,
    segment_code,
    transaction_amount,
    rn
FROM (
    SELECT
        establishment_code,
        segment_code,
        transaction_amount,
        ROW_NUMBER() OVER (PARTITION BY segment_code ORDER BY transaction_amount DESC) AS rn
    FROM cadoc_6334_v2.tmp_ranking_group
) ranked
WHERE rn <= 15;


CREATE TABLE cadoc_6334_v2.tmp_ranking_200 AS
SELECT
    establishment_code,
    segment_code,
    transaction_amount,
    '200ultimos'
FROM (
    SELECT
        establishment_code,
        segment_code,
        transaction_amount,
        ROW_NUMBER() OVER (PARTITION BY segment_code ORDER BY transaction_amount ASC) AS rn
    FROM cadoc_6334_v2.tmp_ranking_group
) ranked
WHERE rn <= 200;


drop table if exists cadoc_6334_v2.tmp_ranking_filtered;

create table cadoc_6334_v2.tmp_ranking_filtered as
select b.year, b.quarter, b.establishment_code, b.function, b.brand, b.capture_mode, b.installments, b.segment_code, 
       b.transaction_amount, b.transaction_quantity, b.avg_mcc_fee
  from cadoc_6334_v2.tmp_ranking_15 a
inner join cadoc_6334_v2.ranking b
  on a.establishment_code = b.establishment_code
  and a.segment_code = b.segment_code;



insert into cadoc_6334_v2.tmp_ranking_filtered
select b.year, b.quarter, -1 as establishment_code, b.function, b.brand, b.capture_mode, b.installments, b.segment_code, 
       sum(b.transaction_amount) as transaction_amount, 
       sum(b.transaction_quantity) as transaction_quantity, 
       round(sum(b.avg_mcc_fee / 100 * b.transaction_amount) / sum(b.transaction_amount) * 100, 2) as avg_mcc_fee
  from cadoc_6334_v2.tmp_ranking_200 a
inner join cadoc_6334_v2.ranking b
  on a.establishment_code = b.establishment_code
  and a.segment_code = b.segment_code
group by 1, 2, 3, 4, 5, 6, 7, 8;


-- validate top 15

select count(1)
  from cadoc_6334_v2.ranking_filtered
where establishment_code != -1;

select count(1)
  from cadoc_6334_v2.tmp_ranking_filtered
where establishment_code != -1;


select count(1)
  from cadoc_6334_v2.ranking_filtered a
left join cadoc_6334_v2.tmp_ranking_filtered b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.establishment_code = b.establishment_code
  and a.function = b.function
  and a.brand = b.brand
  and a.capture_mode = b.capture_mode
  and a.installments = b.installments
  and a.segment_code = b.segment_code
where b.year is null
  and a.establishment_code != -1;


select count(1)
  from cadoc_6334_v2.tmp_ranking_filtered a
left join cadoc_6334_v2.ranking_filtered b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.establishment_code = b.establishment_code
  and a.function = b.function
  and a.brand = b.brand
  and a.capture_mode = b.capture_mode
  and a.installments = b.installments
  and a.segment_code = b.segment_code
where b.year is null
  and a.establishment_code != -1;

select min(a.transaction_amount - b.transaction_amount) as min_diff_transaction_amount,
       max(a.transaction_amount - b.transaction_amount) as max_diff_transaction_amount,
       min(a.transaction_quantity - b.transaction_quantity) as min_diff_transaction_quantity,
       max(a.transaction_quantity - b.transaction_quantity) as max_diff_transaction_quantity,
       min(a.avg_mcc_fee - b.avg_mcc_fee) as min_diff_avg_mcc_fee,
       max(a.avg_mcc_fee - b.avg_mcc_fee) as max_diff_avg_mcc_fee
  from cadoc_6334_v2.tmp_ranking_filtered a
inner join cadoc_6334_v2.ranking_filtered b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.establishment_code = b.establishment_code
  and a.function = b.function
  and a.brand = b.brand
  and a.capture_mode = b.capture_mode
  and a.installments = b.installments
  and a.segment_code = b.segment_code
where a.establishment_code != -1;


-- validate bottom
select count(1)
  from cadoc_6334_v2.ranking_filtered
where establishment_code = -1;

select count(1)
  from cadoc_6334_v2.tmp_ranking_filtered
where establishment_code = -1;


select count(1)
  from cadoc_6334_v2.ranking_filtered a
left join cadoc_6334_v2.tmp_ranking_filtered b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.establishment_code = b.establishment_code
  and a.function = b.function
  and a.brand = b.brand
  and a.capture_mode = b.capture_mode
  and a.installments = b.installments
  and a.segment_code = b.segment_code
where b.year is null
  and a.establishment_code = -1;


select count(1)
  from cadoc_6334_v2.tmp_ranking_filtered a
left join cadoc_6334_v2.ranking_filtered b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.establishment_code = b.establishment_code
  and a.function = b.function
  and a.brand = b.brand
  and a.capture_mode = b.capture_mode
  and a.installments = b.installments
  and a.segment_code = b.segment_code
where b.year is null
  and a.establishment_code = -1;

select min(a.transaction_amount - b.transaction_amount) as min_diff_transaction_amount,
       max(a.transaction_amount - b.transaction_amount) as max_diff_transaction_amount,
       min(a.transaction_quantity - b.transaction_quantity) as min_diff_transaction_quantity,
       max(a.transaction_quantity - b.transaction_quantity) as max_diff_transaction_quantity,
       min(a.avg_mcc_fee - b.avg_mcc_fee) as min_diff_avg_mcc_fee,
       max(a.avg_mcc_fee - b.avg_mcc_fee) as max_diff_avg_mcc_fee
  from cadoc_6334_v2.tmp_ranking_filtered a
inner join cadoc_6334_v2.ranking_filtered b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.establishment_code = b.establishment_code
  and a.function = b.function
  and a.brand = b.brand
  and a.capture_mode = b.capture_mode
  and a.installments = b.installments
  and a.segment_code = b.segment_code
where a.establishment_code = -1;

drop table cadoc_6334_v2.tmp_ranking_group;
drop table cadoc_6334_v2.tmp_ranking_15;
drop table cadoc_6334_v2.tmp_ranking_200;
drop table cadoc_6334_v2.tmp_ranking_filtered;

