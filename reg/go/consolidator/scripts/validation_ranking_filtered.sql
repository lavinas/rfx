-- Active: 1766518799113@@127.0.0.1@5434@reg@cadoc_6334_v2



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


select count(1)
   from cadoc_6334_v2.tmp_ranking_15 a
inner join cadoc



select sum(avg_mdr_fee * transaction_amount / 100) as total_mdr_fee, sum(transaction_amount) as total_transaction_amount
  from cadoc_6334_v2.desconto;



select segment_code, establishment_code, count(1) 
  from cadoc_6334_v2.ranking_filtered
group by 1, 2
order by 1, 2 desc
limit 500




select segment_code, count(distinct establishment_code)
  from cadoc_6334_v2.ranking_filtered
group by 1
order by 1
limit 500
