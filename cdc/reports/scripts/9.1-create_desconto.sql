DO $$
DECLARE
   -- visa 
   target_bandeira_visa INT := 1;
   target_tpv_visa NUMERIC := 191_892_055.39;
   target_qtd_visa INT := 3_609_877;
   -- mastercard
   target_bandeira_mc INT := 2;
   target_tpv_mc NUMERIC := 236_661_956.14;
   target_qtd_mc INT := 5_808_775;
   -- elo
   target_bandeira_elo INT := 8;
   target_tpv_elo NUMERIC := 60_880_019.10;
   target_qtd_elo INT := 1_184_507;
BEGIN
   CALL reports.create_desconto(target_bandeira_visa, target_tpv_visa, target_qtd_visa);
   CALL reports.create_desconto(target_bandeira_mc, target_tpv_mc, target_qtd_mc);
   CALL reports.create_desconto(target_bandeira_elo, target_tpv_elo, target_qtd_elo);
END; $$; LANGUAGE plpgsql;