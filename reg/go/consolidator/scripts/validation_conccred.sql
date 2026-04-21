-- Active: 1766518799113@@127.0.0.1@5434@reg
-- amounts

drop table if exists cadoc_6334_v2.tmp_conccred;

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
       0 number_accredited_establishments,
       0 number_active_establishments, 
       round(sum(transaction_amount), 2) as transaction_amount,
       count(1) as transaction_quantity
  from transaction_v4.transaction a
left join apoio.segmentos b
  on a.establishment_mcc >= b.mcc_init
  and a.establishment_mcc <= b.mcc_end
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2
  group by 1, 2, 3, 4;




update cadoc_6334_v2.tmp_conccred a
  set number_accredited_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_visa = true
                                            and b.has_credit = true),
      number_active_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_visa = true
                                            and b.has_credit = true) * 0.76
where a.year = 2026
  and a.quarter = 1
  and a.brand = 1
  and a.function = 'C';

  update cadoc_6334_v2.tmp_conccred a
  set number_accredited_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_mastercard = true
                                            and b.has_credit = true),
      number_active_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_mastercard = true
                                            and b.has_credit = true) * 0.76
where a.year = 2026
  and a.quarter = 1
  and a.brand = 2
  and a.function = 'C';

  update cadoc_6334_v2.tmp_conccred a
  set number_accredited_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_elo = true
                                            and b.has_credit = true),
      number_active_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_elo = true
                                            and b.has_credit = true) * 0.76
where a.year = 2026
  and a.quarter = 1
  and a.brand = 8
  and a.function = 'C';


update cadoc_6334_v2.tmp_conccred a
  set number_accredited_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_visa = true
                                            and b.has_debit = true),
      number_active_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_visa = true
                                            and b.has_debit = true) * 0.76
where a.year = 2026
  and a.quarter = 1
  and a.brand = 1
  and a.function = 'D';

  update cadoc_6334_v2.tmp_conccred a
  set number_accredited_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_mastercard = true
                                            and b.has_debit = true),
      number_active_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_mastercard = true
                                            and b.has_debit = true) * 0.76
where a.year = 2026
  and a.quarter = 1
  and a.brand = 2
  and a.function = 'D';

  update cadoc_6334_v2.tmp_conccred a
  set number_accredited_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_elo = true
                                            and b.has_debit = true),
      number_active_establishments = (select count(1)
                                           from raw_data_v2.establishments b
                                          where b.accreditation_date < '2026-04-01'
                                            and b.has_elo = true
                                            and b.has_debit = true) * 0.76
where a.year = 2026
  and a.quarter = 1
  and a.brand = 8
  and a.function = 'D';




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
       max(a.transaction_quantity - b.transaction_quantity) as max_diff_qtd,
       min(a.number_accredited_establishments - b.number_accredited_establishments) as min_diff_accredited,
        max(a.number_accredited_establishments - b.number_accredited_establishments) as max_diff_accredited,
        min(a.number_active_establishments - b.number_active_establishments) as min_diff_active,
        max(a.number_active_establishments - b.number_active_establishments) as max_diff_active
  from cadoc_6334_v2.tmp_conccred a
inner join cadoc_6334_v2.conccred b
on a.year = b.year
and a.quarter = b.quarter
and a.function = b.function
and a.brand = b.brand;

drop table cadoc_6334_v2.tmp_conccred;