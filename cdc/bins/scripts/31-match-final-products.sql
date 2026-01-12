update bins bins
   set produto_intermediario = trim(produto_intermediario);

update bins bins
   set produto_final = prod.card_classification_code
from final_products prod
where bins.produto_intermediario = prod.card_name;

update bins bins
   set produto_final = 31 -- Basico Nacional
where bins.produto_intermediario = 'MAESTRO'
  and bins.pais = 'BRA';

update bins bins
   set produto_final = 32 -- Basico Internacional
where bins.produto_intermediario = 'MAESTRO'
  and bins.pais != 'BRA';


select * from bins


