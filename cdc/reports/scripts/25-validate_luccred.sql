

-- verificao da margem de erro dos descontos aplicados
select sum(valor),
       sum(minimo),
       sum(maximo)
  from (
select receitataxadescontobruta valor,
       0 minimo,
       0 maximo
  from reports.luccred_ch
union all
select 0 valor,
       round(sum((a.taxa_desconto_media - 0.1 * a.desvio_padrao_taxa_desconto) * a.valor_transacoes /100), 2) minimo,
       round(sum((a.taxa_desconto_media + 0.1 * a.desvio_padrao_taxa_desconto) * a.valor_transacoes /100), 2) maximo
  from reports.descontos a
) as resultados;


select sum(valor),
       sum(minimo),
       sum(maximo)
  from (
select CustoTarifaIntercambio valor,
       0 minimo,
       0 maximo
  from reports.luccred_ch
union all
select 0 valor,
   sum(round((a.tarifa_intercambio - 0.03) * a.valor_transacoes /100, 2)) minimo,
   sum(round((a.tarifa_intercambio + 0.03) * a.valor_transacoes /100, 2)) maximo
  from reports.intercam a
) as resultados;

-- 114.315,59
select ReceitaTaxaDescontoBruta + ReceitaAluguelEquipamentosConectividade + ReceitaOutras -
       CustoTarifaIntercambio - CustoMarketingPropaganda - CustoTaxasAcessoBandeiras -
       CustoRisco - CustoProcessamento - CustoOutros as LucroBruto
  from reports.luccred_ch;
