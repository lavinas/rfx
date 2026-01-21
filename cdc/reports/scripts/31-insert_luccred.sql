SELECT
    'INSERT INTO cadoc_6334_lucrcred (Ano, Trimestre, ReceitaTaxaDescontoBruta, ReceitaAluguelEquipamentosConectividade, ReceitaOutras, CustoTarifaIntercambio, CustoMarketingPropaganda, CustoTaxasAcessoBandeiras, CustoRisco, CustoProcessamento, CustoOutros) VALUES ('
    || coalesce(ano::text,'NULL') || ', '
    || coalesce(trimestre::text,'NULL') || ', '
    || coalesce(receitataxadescontobruta::text,'NULL') || ', '
    || coalesce(receitaaluguelequipamentosconectividade::text,'NULL') || ', '
    || coalesce(receitaoutras::text,'NULL') || ', '
    || coalesce(custotarifaintercambio::text,'NULL') || ', '
    || coalesce(customarketingpropaganda::text,'NULL') || ', '
    || coalesce(custotaxasacessobandeiras::text,'NULL') || ', '
    || coalesce(custorisco::text,'NULL') || ', '
    || coalesce(custoprocessamento::text,'NULL') || ', '
    || coalesce(custooutros::text,'NULL') || ');'
AS insert_sql
FROM reports.luccred_ch
ORDER BY id;
