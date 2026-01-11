SET search_path TO bins;

-- apply
update bins
   set bandeira = 'ELO',
         pais = 'BR',
        produto_base = elo.f07_c007_card_type,
        banco_base = iss.f04_c005_bank_name,
        produto_intermediario = elo.f07_c007_card_type,
        banco_intermediario = iss.f04_c005_bank_name
from elo_bin elo
left join elo_issuer iss
  on elo.f03_c003_bank_id = iss.f02_c002_bank_id
where cast(bin as integer) between elo.bin_range_low and elo.bin_range_high;

-- checks
select count(1) from bins where bandeira is not null;
select count(1) from bins where bandeira is not null and banco_intermediario is null;
