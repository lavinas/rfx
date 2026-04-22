

create table cadoc_6334_v2.tmp_infrterm as
select b.federation_unit,
       count(1) as pos_total_quantity,
       0 pos_shared_quantity,
       sum(case when a.terminal_type = 'POS' then 1 else 0 end) pos_chip_quantity,
       sum(case when a.terminal_type = 'TEF' then 1 else 0 end) pdv_quantity
  from raw_data_v2.terminals_transaction a
inner join raw_data_v2.establishments b
    on a.establishment_code = b.establishment_code
   and b.accreditation_date < '2026-04-01'::date
group by 1;


select count(1)
  from cadoc_6334_v2.tmp_infrterm;

select count(1)
  from cadoc_6334_v2.infrterm;


select a.*
  from cadoc_6334_v2.infresta a
left join cadoc_6334_v2.infrterm b
   on a.federation_unit = b.federation_unit
where b.federation_unit is null
  and a.establishment_total_quantity > 2;


select count(1)
  from cadoc_6334_v2.infrterm a
left join cadoc_6334_v2.tmp_infrterm b
   on a.federation_unit = b.federation_unit
 and a.pos_total_quantity = b.pos_total_quantity
 and a.pos_shared_quantity = b.pos_shared_quantity
 and a.pos_chip_quantity = b.pos_chip_quantity
 and a.pdv_quantity = b.pdv_quantity
where b.federation_unit is null;


select count(1)
  from cadoc_6334_v2.tmp_infrterm a
left join cadoc_6334_v2.infrterm b
   on a.federation_unit = b.federation_unit
 and a.pos_total_quantity = b.pos_total_quantity
 and a.pos_shared_quantity = b.pos_shared_quantity
 and a.pos_chip_quantity = b.pos_chip_quantity
 and a.pdv_quantity = b.pdv_quantity
where b.federation_unit is null;

drop table cadoc_6334_v2.tmp_infrterm;