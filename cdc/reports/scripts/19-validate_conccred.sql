select * from reports.conccred_ch;

select bandeira, funcao, count(1)
  from reports.conccred_ch
 group by 1,2
having count(1) != 1;


select sum(quantidade_estabelecimentos_totais) from reports.infresta_ch;

-- estabelecimentos
select max(quantidade_estabelecimentos_credenciados) estab
  from reports.conccred_ch
union all
select sum(quantidade_estabelecimentos_totais) estab
  from reports.infresta_ch;

-- totais
select funcao,
       bandeira,
       sum(valor_transacoes) valor_total,
       sum(quantidade_transacoes) qtde_total,
       'concred' tipo
  from reports.conccred_ch
group by 1, 2
union ALL
select funcao,
       bandeira,
       sum(valor_transacoes) valor_total,
         sum(quantidade_transacoes) qtde_total,
        'descontos' tipo
from reports.descontos_ch
group by 1, 2
order by funcao, bandeira, tipo;