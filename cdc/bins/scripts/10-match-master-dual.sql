update bins.bins bins
    set bandeira = 'MASTERCARD',
        pais = mbin.country_code_alpha,
        produto_base = mbin.gcms_product_id,
        banco_base = mbin.member_id
    from master_bin_dual_message_system_mpe mbin
    where cast(bins.bin || '000' as bigint) between mbin.bin_range_low and mbin.bin_range_high
    and mbin.card_program_identifier NOT IN ('CIR', 'MSI')
    and bins.bandeira is null;
