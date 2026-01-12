SET search_path TO bins;

update bins.bins bins
    set bandeira = 'MASTER2',
        pais = mbin.country_code,
        produto_base = mbin.maestro_product_code,
        banco_base = mbin.issuer_processor_id
    from bins.master_bin_single_message_system_fit mbin
    where cast(bins.bin || '000' as bigint) between mbin.identifier_low and mbin.identifier_high
    and bins.bandeira is null;

