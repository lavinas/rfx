SET search_path TO bins;

update  bins bins4
    set bandeira = 'VISA',
        pais = f16_country,
        produto_base = b.f29_product_id,
        banco_base = b.f08_issuer_identifier
  from visa_bin_ardef b
  where cast(bins4.bin || '0' as integer) between b.key_range_low and b.key_range_high
    and bins4.bandeira is NULL;
update bins bins4
    set bandeira = 'VISA',
        pais = f16_country,
        produto_base = b.f29_product_id,
        banco_base = b.f08_issuer_identifier
  from visa_bin_ardef b
  where cast(bins4.bin as integer) between b.key_range_low and b.key_range_high
    and bins4.bandeira is NULL;
-- atualizar visa 3
update bins bins4
    set bandeira = 'VISA',
        pais = f16_country,
        produto_base = b.f29_product_id,
        banco_base = b.f08_issuer_identifier
  from visa_bin_ardef b
  where cast(substr(bins4.bin, 1, 7) as integer) between b.key_range_low and b.key_range_high
    and bins4.bandeira is NULL;