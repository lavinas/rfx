
drop table if exists reports.luccred_ch;
CREATE TABLE IF NOT EXISTS reports.luccred_ch (
    id BIGSERIAL PRIMARY KEY,
    Ano SMALLINT NOT NULL,
    Trimestre SMALLINT NOT NULL,
    ReceitaTaxaDescontoBruta NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    ReceitaAluguelEquipamentosConectividade NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    ReceitaOutras NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    CustoTarifaIntercambio NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    CustoMarketingPropaganda NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    CustoTaxasAcessoBandeiras NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    CustoRisco NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    CustoProcessamento NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    CustoOutros NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT ux_ano_trimestre UNIQUE (Ano, Trimestre)
);


INSERT INTO reports.luccred_ch (
    Ano,
    Trimestre,
    ReceitaTaxaDescontoBruta,
    ReceitaAluguelEquipamentosConectividade,
    ReceitaOutras,
    CustoTarifaIntercambio,
    CustoMarketingPropaganda,
    CustoTaxasAcessoBandeiras,
    CustoRisco,
    CustoProcessamento,
    CustoOutros
) VALUES (
2025,
4,
9017307.88,
1733861.31,
2305460.42,
4456252.49,
202671.95,
1613244.38,
0.00,
468808.17,
6201337.03
);