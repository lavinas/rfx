SET search_path TO bins;

-- find banco_intermediario for MASTER cards
update bins
    set banco_intermediario = trim(ip72.member_name)
from bins.master_ip0072t1 ip72
where bins.banco_base = ip72.member_id
  and bins.bandeira like 'MASTER%';
-- work around for missing values
update bins
    set banco_intermediario = 'CREDIT AGRICOLE S.A.'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'FRA';
update bins
    set banco_intermediario = 'BANCO C6 SA'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'CYM';
update bins
    set banco_intermediario = 'DEUTSCHER SPARKASSEN UND GIROV'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'DEU';
update bins
    set banco_intermediario = 'VISECA PAYMENT SERVICES SA'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'CHE';
update bins
    set banco_intermediario = 'BANCOLOMBIA S.A.'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'COL';
update bins
    set banco_intermediario = 'BANK OF AMERICA, NATIONAL ASSO'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'USA';
update bins
    set banco_intermediario = 'PSA PAYMENT SERVICES AUSTRIA G'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'AUT';
update bins
    set banco_intermediario = 'REVOLUT LTD'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'IRL';
update bins
    set banco_intermediario = 'MONZO BANK LIMITED'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'GBR';
update bins
    set banco_intermediario = 'THE HONGKONG AND SHANGHAI BANK'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'HKG';
update bins
    set banco_intermediario = 'CIMB BANK BERHAD'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'MYS';
update bins
    set banco_intermediario = 'ITAU UNIBANCO S.A.'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'BRA';
update bins
    set banco_intermediario = 'ING Bank'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'NLD';
update bins
    set banco_intermediario = 'POSTEPAY S.P.A'
where bandeira like 'MASTER%'
  and banco_intermediario is null
  and pais = 'ITA';