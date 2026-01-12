SET search_path TO bins;

update bins a
set banco_intermediario = c.issuer
from visa_issuers_bin b
inner join visa_issuers_main c on c.id = b.id
where a.bandeira = 'VISA'
  and a.banco_base = b.bin;

update bins a
set banco_intermediario = c.issuer
from visa_issuers_bin b
inner join visa_issuers_main c on c.id = b.id
where a.bandeira = 'VISA'
  and a.banco_base || '00' = b.bin
 and a.banco_intermediario is null;

update bins a
set banco_intermediario = c.issuer
from visa_issuers_bin b
inner join visa_issuers_main c on c.id = b.id
where a.bandeira = 'VISA'
  and a.bin = b.bin
  and a.banco_intermediario is null;

update bins a
set banco_intermediario = c.issuer
from visa_issuers_bin b
inner join visa_issuers_main c on c.id = b.id
where a.bandeira = 'VISA'
  and substr(a.bin, 1, 6) = b.bin
  and a.banco_intermediario is null;

update bins a
set banco_intermediario = c.issuer
from visa_issuers_bin b
inner join visa_issuers_main c on c.id = b.id
where a.bandeira = 'VISA'
  and substring(b.bin, 1, length(a.banco_base)) = a.banco_base
  and a.banco_intermediario is null;

update bins a
set banco_intermediario = c.issuer
from visa_issuers_bin b
inner join visa_issuers_main c on c.id = b.id
where a.bandeira = 'VISA'
  and substring(b.bin, 1, length(a.banco_base)) = a.banco_base
  and a.banco_intermediario is null;

update bins
    set banco_intermediario = 'JPMorgan Chase Bank N.A.'
where bandeira = 'VISA'
  and banco_intermediario is null
  and banco_base = '400087';

update bins
    set banco_intermediario = 'Neon Pagamentos S.A.'
where bandeira = 'VISA'
  and banco_intermediario is null
  and banco_base = '707742';

update bins
    set banco_intermediario = 'JPMorgan Chase Bank N.A.'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'US';


update bins
    set banco_intermediario = 'Banco Bradesco S.A.'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'BR';

update bins
    set banco_intermediario = 'DBS Bank Ltd'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'SG';

update bins
    set banco_intermediario = 'Banco de Credito del Peru'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'PE';

update bins
    set banco_intermediario = 'BAWAG P.S.K. Bank fuer Arbeit und Wirtschaft und Osterreichische Postsparkasse A'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'DE';

update bins
    set banco_intermediario = 'Cuscal Limited'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'AU';

update bins
    set banco_intermediario = 'The Hongkong and Shanghai Banking Corporation Limited'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'HK';

update bins
    set banco_intermediario = 'Banco Popular de Puerto Rico'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'PR';

update bins
    set banco_intermediario = 'The Toronto-Dominion Bank'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'CA';

update bins
    set banco_intermediario = 'Foreningen af Danske Kortudstedere F.M.B.A'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'DK';

update bins
    set banco_intermediario = 'Banco Pichincha C.A.'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'EC';

update bins
    set banco_intermediario = 'Asociacion Cibao de Ahorros y Prestamos'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'DO';

update bins
    set banco_intermediario = 'Wise Europe SA/NV'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'BE';

update bins
    set banco_intermediario = 'Intesa Sanpaolo S.P.A.'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'IT';

update bins
    set banco_intermediario = 'BARCLAYS BANK UK PLC'
where bandeira = 'VISA'
  and banco_intermediario is null
  and pais = 'GB';

update bins
    set banco_intermediario = 'JPMorgan Chase Bank N.A.'
where bandeira = 'VISA'
  and banco_intermediario is null;
