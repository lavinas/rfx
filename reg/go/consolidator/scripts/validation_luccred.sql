
create table cadoc_6334_v2.tmp_luccred as
select extract(year from a.transaction_date) as year,
       extract(quarter from a.transaction_date) as quarter,
       sum(a.revenue_mdr_value) as gross_revenue,
       0 as rental_revenue,
       0 as others_revenue,
       sum(a.cost_interchange_value) as interchange_cost,
       0 as marketing_cost,
       0 as brand_access_cost,
       0 as risk_cost,
       0 as processing_cost,
       0 as others_cost
  from transaction_v4.transaction a
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-01-06'
  and a.status_id = 2
  group by 1, 2;



select count(1) from cadoc_6334_v2.tmp_luccred;

select count(1) from cadoc_6334_v2.luccred;

select count(1) from cadoc_6334_v2.luccred a
left join cadoc_6334_v2.tmp_luccred b
  on a.year = b.year
  and a.quarter = b.quarter
where a.year is null;


select count(1) from cadoc_6334_v2.tmp_luccred a
left join cadoc_6334_v2.luccred b
  on a.year = b.year
  and a.quarter = b.quarter
where a.year is null;


select count(1) 
from cadoc_6334_v2.luccred a
inner join cadoc_6334_v2.tmp_luccred b
  on a.year = b.year
  and a.quarter = b.quarter
  and 
  (a.gross_revenue != b.gross_revenue or 
   a.rental_revenue != b.rental_revenue or
   a.others_revenue != b.others_revenue or
   a.interchange_cost != b.interchange_cost or
   a.marketing_cost != b.marketing_cost or
   a.brand_access_cost != b.brand_access_cost or
   a.risk_cost != b.risk_cost or
   a.processing_cost != b.processing_cost or
   a.others_cost != b.others_cost
  )  ;