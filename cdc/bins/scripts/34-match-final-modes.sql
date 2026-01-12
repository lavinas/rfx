update bins
   set banco_intermediario = trim(banco_intermediario);

update bins
   set modalidade_final = mod.codigo
from final_modes mod
where bins.banco_intermediario = mod.instituicao;

update bins
   set modalidade_final = 'P'
 where modalidade_final is null
   and upper(banco_intermediario) like '%BANK%';

update bins
   set modalidade_final = 'P'
 where modalidade_final is null
   and upper(banco_intermediario) like '%BANCO%';

update bins
   set modalidade_final = 'P'
 where modalidade_final is null
   and upper(banco_intermediario) like '%PAY%';

update bins
   set modalidade_final = 'P'
 where modalidade_final is null
   and upper(banco_intermediario) like '%PAID%';

update bins
   set modalidade_final = 'P'
 where modalidade_final is null
   and upper(banco_intermediario) like '%PAGA%';

update bins
   set modalidade_final = 'P'
 where modalidade_final is null
   and upper(banco_intermediario) like '%CARD%';


update bins
   set modalidade_final = 'P'
 where modalidade_final is null
   and upper(banco_intermediario) like '%CART%';

update bins
   set modalidade_final = 'P'
 where modalidade_final is null
   and upper(banco_intermediario) like '%CREDIT%';

update bins
   set modalidade_final = 'P'
 where modalidade_final is null
   and upper(banco_intermediario) like '%FINANC%';

-- finally, set all remaining nulls to 'P'
update bins
   set modalidade_final = 'P'
 where modalidade_final is null;



