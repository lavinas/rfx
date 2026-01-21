select uf,
       count(1)
  from reports.infresta_ch
 group by uf
 having count(1) != 1;


 select * 
   from reports.infresta_ch
  where quantidade_estabelecimentos_totais != 
        quantidade_estabelecimentos_captura_eletronica + quantidade_estabelecimentos_captura_manual + quantidade_estabelecimentos_captura_remota;