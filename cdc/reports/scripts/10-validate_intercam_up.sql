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


select ano, trimestre, produto, modalidade_cartao, funcao, bandeira, forma_captura, numero_parcelas, codigo_segmento, count(1)
  from intercam_up
group by 1, 2, 3, 4, 5, 6, 7, 8, 9
having count(1) > 1
limit 100000;

select *
  from intercam_up
where not (intercam_up is not null)
limit 100000;


select *
  from intercam_up
where ano != 2025
limit 100000;


select *
  from intercam_up
where trimestre != 4
limit 100000;

select *
  from intercam_up
where valor_transacoes <= 0
limit 100000;

-- tarifa_intercambio - 0 linhas
select *
  from intercam_up
where tarifa_intercambio <= 0 or tarifa_intercambio > 9
limit 100000;

-- quantidade_transacoes - 0 linhas
select *
  from intercam_up
where quantidade_transacoes <= 0
limit 100000;


-- valida produto
select *
  from intercam
where produto < 31 or produto > 38 or produto is null;


select *
  from intercam_up
where modalidade_cartao not in ('P', 'C') or modalidade_cartao is null;


select *
  from intercam_up
where funcao not in ('C', 'D', 'E') or funcao is null;

select *
  from intercam_up
where bandeira is null or bandeira not in (1,2,3,4,5,6,7,8,99)
limit 100000;

select *
  from intercam_up
where forma_captura is null or forma_captura not in (1,2,3,4,5,6);

select *
  from intercam_up
where numero_parcelas is null or numero_parcelas < 1 or numero_parcelas > 12
limit 100000;

select *
  from intercam_up
where codigo_segmento is null or not (codigo_segmento >= 401 and codigo_segmento <= 428);