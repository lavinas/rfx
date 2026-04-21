-- Active: 1766518799113@@127.0.0.1@5434@reg



create table cadoc_6334_v2.tmp_infresta as
select federation_unit, 
       count(1) establishment_total_quantity,
       0 establishment_manual_capture_quantity,
       count(1) establishment_eletronic_capture_quantity,
       0 establishment_remote_capture_quantity
  from raw_data_v2.establishments
where accreditation_date < '2026-04-01'::date
group by 1;


select count(1) from cadoc_6334_v2.tmp_infresta;

select count(1) from cadoc_6334_v2.infresta;


select count(1)
  from cadoc_6334_v2.infresta a
left join cadoc_6334_v2.tmp_infresta b
  on a.federation_unit = b.federation_unit
 where a.establishment_total_quantity != b.establishment_total_quantity
    or a.establishment_manual_capture_quantity != b.establishment_manual_capture_quantity
    or a.establishment_eletronic_capture_quantity != b.establishment_eletronic_capture_quantity
    or a.establishment_remote_capture_quantity != b.establishment_remote_capture_quantity;


select count(1)
  from cadoc_6334_v2.tmp_infresta a
left join cadoc_6334_v2.infresta b
  on a.federation_unit = b.federation_unit
 where a.establishment_total_quantity != b.establishment_total_quantity
    or a.establishment_manual_capture_quantity != b.establishment_manual_capture_quantity
    or a.establishment_eletronic_capture_quantity != b.establishment_eletronic_capture_quantity
    or a.establishment_remote_capture_quantity != b.establishment_remote_capture_quantity;

drop table cadoc_6334_v2.tmp_infresta;