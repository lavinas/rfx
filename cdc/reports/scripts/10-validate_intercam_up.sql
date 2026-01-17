select case bandeira 
         when 1 then 'Visa'
         when 2 then 'Master'
         when 8 then 'Elo'
         else 'Erro'
       end as bandeira, 
       sum(valor_transacoes),
         sum(quantidade_transacoes)
  from reports.intercam_up
group by 1
order by 1;


select count(1)
  from reports.intercam_up;

select count(1)
  from reports.intercam;