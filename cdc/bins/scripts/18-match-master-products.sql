
-- find produto_intermediario for MASTER cards
update bins
    set produto_intermediario = codes.description
from master_product_codes codes
where bins.produto_base = codes.licensed_product_id
  and bins.bandeira like 'MASTER%';
-- work around for missing values
update bins
    set produto_intermediario = 'MAESTRO'
where bandeira = 'MASTER2'
  and produto_intermediario is null;

update bins
    set produto_intermediario = 'MASTERCARD WORLD ELITE'
where bandeira = 'MASTERCARD'
  and produto_intermediario is null
  and pais IN ('USA', 'SGP', 'ARE', 'SAU');

update bins
    set produto_intermediario = 'MASTERCARD STANDARD CARD'
where bandeira = 'MASTERCARD'
  and produto_intermediario is null
  and pais IN ('BRA');