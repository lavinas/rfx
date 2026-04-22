-- Active: 1774368236280@@192.168.100.78@5436@dev_regulat
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


select gross_revenue + rental_revenue + others_revenue - interchange_cost - marketing_cost - brand_access_cost - risk_cost - processing_cost - others_cost as cost
  from cadoc_6334_v2.luccred
 where year = 2026
   and quarter = 1;