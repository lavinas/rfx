select case bandeira 
         when 1 then 'Visa'
         when 2 then 'Master'
         when 8 then 'Elo'
         else 'Erro'
       end as bandeira, 
       sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes * tarifa_intercambio / 100)/sum(valor_transacoes), 4) as valor_intercambio
  from reports.intercam_up
group by 1
order by 1;


select sum(valor_transacoes),
       sum(quantidade_transacoes),
       round(sum(valor_transacoes * tarifa_intercambio / 100)/sum(valor_transacoes), 4) as valor_intercambio
  from reports.intercam_up;

select * from reports.intercam_up


select count(1)
  from reports.intercam_up;

select count(1)
  from reports.intercam;