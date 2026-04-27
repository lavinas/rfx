

-- ranking

select count(1)
  from cadoc_6334_v3.ranking;

select count(1)
from cadoc_6334_v2.ranking_filtered;


select year, quarter, establishment_code, function, brand, capture_mode, installments, segment_code, transaction_amount, transaction_quantity, avg_mcc_fee
  from cadoc_6334_v2.ranking_filtered
except
select year, quarter, establishment_code, function, brand, capture_mode, installments, segment_code, transaction_amount, transaction_quantity, avg_mcc_fee
  from cadoc_6334_v3.ranking;

-- segmento

select count(1)
  from cadoc_6334_v2.segmento;

select count(1)
  from cadoc_6334_v3.segmento;

select year, quarter, segment_name, segment_description, segment_code
  from cadoc_6334_v2.segmento
except
select year, quarter, segment_name, segment_description, segment_code
  from cadoc_6334_v3.segmento;

-- conccred

select count(1)
  from cadoc_6334_v2.conccred;

select count(1)
  from cadoc_6334_v3.conccred;

select year, quarter, brand, function, number_accredited_establishments, number_active_establishments, transaction_amount, transaction_quantity
  from cadoc_6334_v2.conccred
except
select year, quarter, brand, function, number_accredited_establishments, number_active_establishments, transaction_amount, transaction_quantity
  from cadoc_6334_v3.conccred;


-- desconto

select count(1)
  from cadoc_6334_v2.desconto;

select count(1)
  from cadoc_6334_v3.desconto;


select year, quarter, function, brand, capture_mode, installments, segment_code
  from cadoc_6334_v2.desconto
except
select year, quarter, function, brand, capture_mode, installments, segment_code
  from cadoc_6334_v3.desconto;


select min(a.avg_mdr_fee - b.avg_mdr_fee) as min_avg_mdr_fee_diff,
       max(a.avg_mdr_fee - b.avg_mdr_fee) as max_avg_mdr_fee_diff,
       min(a.min_mdr_fee - b.min_mdr_fee) as min_min_mdr_fee_diff,
       max(a.min_mdr_fee - b.min_mdr_fee) as max_min_mdr_fee_diff,
       min(a.max_mdr_fee - b.max_mdr_fee) as min_max_mdr_fee_diff,
       max(a.max_mdr_fee - b.max_mdr_fee) as max_max_mdr_fee_diff,
       min(a.stdev_mdr_fee - b.stdev_mdr_fee) as min_stdev_mdr_fee_diff,
       max(a.stdev_mdr_fee - b.stdev_mdr_fee) as max_stdev_mdr_fee_diff,
       min(a.transaction_amount - b.transaction_amount) as min_transaction_amount_diff,
       max(a.transaction_amount - b.transaction_amount) as max_transaction_amount_diff,
       min(a.transaction_quantity - b.transaction_quantity) as min_transaction_quantity_diff,
       max(a.transaction_quantity - b.transaction_quantity) as max_transaction_quantity_diff
  from cadoc_6334_v2.desconto a
inner join cadoc_6334_v3.desconto b
    on a.year = b.year
     and a.quarter = b.quarter
     and a.function = b.function
     and a.brand = b.brand
     and a.capture_mode = b.capture_mode
     and a.installments = b.installments
     and a.segment_code = b.segment_code;
 

 -- intercam

select count(1)
  from cadoc_6334_v2.intercam; 

select count(1)
  from cadoc_6334_v3.intercam;

select year, quarter, product_code, card_type, function, brand, capture_mode, installments, segment_code, interchange_fee, transaction_amount, transaction_quantity 
  from cadoc_6334_v2.intercam a
except
select year, quarter, product_code, card_type, function, brand, capture_mode, installments, segment_code, interchange_fee, transaction_amount, transaction_quantity 
  from cadoc_6334_v3.intercam a;


-- infresta
select count(1)
  from cadoc_6334_v2.infresta;

select count(1)
  from cadoc_6334_v3.infresta;

select year, quarter, federation_unit, establishment_total_quantity, establishment_manual_capture_quantity, establishment_eletronic_capture_quantity, establishment_remote_capture_quantity
    from cadoc_6334_v2.infresta a
except
select year, quarter, federation_unit, establishment_total_quantity, establishment_manual_capture_quantity, establishment_eletronic_capture_quantity, establishment_remote_capture_quantity
    from cadoc_6334_v3.infresta a;


-- infrterm

select count(1), sum(pos_total_quantity)
  from cadoc_6334_v2.infrterm; 

select count(1), sum(pos_total_quantity)
  from cadoc_6334_v3.infrterm;


select count(1)
  from raw_data_v2.terminals_transaction


select year, quarter, federation_unit, pos_total_quantity, pos_shared_quantity, pos_chip_quantity, pdv_quantity
    from cadoc_6334_v2.infrterm a
except
select year, quarter, federation_unit, pos_total_quantity, pos_shared_quantity, pos_chip_quantity, pdv_quantity 
  from cadoc_6334_v3.infrterm a;


select year, quarter, federation_unit, pos_total_quantity, pos_shared_quantity, pos_chip_quantity, pdv_quantity
    from cadoc_6334_v3.infrterm a
except
select year, quarter, federation_unit, pos_total_quantity, pos_shared_quantity, pos_chip_quantity, pdv_quantity 
  from cadoc_6334_v2.infrterm a;


select year, quarter, federation_unit, 'V3' which, pos_total_quantity, pos_shared_quantity, pos_chip_quantity, pdv_quantity
    from cadoc_6334_v3.infrterm a
union ALL
select year, quarter, federation_unit, 'V2' which, pos_total_quantity, pos_shared_quantity, pos_chip_quantity, pdv_quantity 
  from cadoc_6334_v2.infrterm a
order by 1, 2, 3, 4;