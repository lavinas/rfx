-- Active: 1766518799113@@127.0.0.1@5434@reg
create table cadoc_6334_v2.luccred_save as
select * from cadoc_6334_v2.luccred;


update cadoc_6334_v2.luccred
   set gross_revenue =      7884916.66,
       rental_revenue =     2183751.22,
       others_revenue =     1512410.24,
       interchange_cost =   3910021.60,
       marketing_cost =      191077.94,
       brand_access_cost =  1515123.96,
       risk_cost =                0.00,
       processing_cost =     406930.64,
       others_cost =        5686873.32
 where year = 2026
   and quarter = 1;

insert into cadoc_6334_v2.luccred (id, created_at, updated_at, year, quarter, gross_revenue, rental_revenue, others_revenue, interchange_cost, marketing_cost, brand_access_cost, risk_cost, processing_cost, others_cost)
values (1, now(), now(), 2026, 1, 7884916.66, 2183751.22, 1512410.24, 3910021.60, 191077.94, 1515123.96, 0.00, 406930.64, 5686873.32);

select * from cadoc_6334_v2.luccred


select gross_revenue + rental_revenue + others_revenue - interchange_cost - marketing_cost - brand_access_cost - risk_cost - processing_cost - others_cost as cost
  from cadoc_6334_v2.luccred
 where year = 2026
   and quarter = 1;