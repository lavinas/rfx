select * from reports.infrterm_ch;


select uf,
       count(1)
  from reports.infrterm_ch
  group by uf
  having count(1) != 1;

select count(1)
  from reports.infrterm_ch
where quantidade_total != quantidade_pdv + quantidade_pos_leitora_chip + quantidade_pos_compartilhados;

-- totals 
select sum(quantidade_total) terms
  from reports.infrterm_ch
union all
select count(1) terms
  from apoio.terminais;


select uf,
   'infrterm' tipo,
   sum(quantidade_total) quantidade_total
 from reports.infrterm_ch
group by 1, 2
union all
select uf,
       'apoio.terminais' tipo,
       count(1) quantidade_total
  from apoio.terminais a
 left join apoio.estabelecimentos b
    on a.codigo_estabelecimento = b.codigo_estabelecimento
    where b.codigo_estabelecimento is not null
group by 1, 2
order by uf, tipo;
