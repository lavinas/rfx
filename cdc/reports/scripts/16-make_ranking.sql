-- cria tabela reports.ranking_ch
drop table if exists reports.ranking_ch;
CREATE TABLE IF NOT EXISTS reports.ranking_ch (
    id BIGSERIAL PRIMARY KEY,
    sync_status SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    ano SMALLINT NOT NULL,
    trimestre SMALLINT NOT NULL,
    codigo_estabelecimento VARCHAR(20) NOT NULL,
    funcao CHAR(1) NOT NULL,
    bandeira SMALLINT NOT NULL,
    forma_captura SMALLINT NOT NULL,
    numero_parcelas SMALLINT NOT NULL,
    codigo_segmento INTEGER NOT NULL,
    valor_transacoes NUMERIC(18,2) NOT NULL,
    quantidade_transacoes INTEGER NOT NULL,
    taxa_desconto_media NUMERIC(5,2) NOT NULL
);

-- atualiza gestao com segmento correto
update apoio.gestao
   set segmento = csmc.segment
  from apoio.segmentos csmc
 where csmc.mcc_init::int = apoio.gestao.mcc::int
   and apoio.gestao.segmento is null;



-- criando schema ranking
create schema ranking;

-- seleciona tpv e clientes por segmento
drop table if exists ranking.ranking_group;
CREATE TABLE ranking.ranking_group AS
SELECT
    ctg.codigo_estabelecimento,
    ctg.segmento as segmento,
    SUM(ctg.sum_valor_transacoes) AS valor_transacao
FROM apoio.gestao ctg
GROUP BY 1, 2;

-- seleciona os 15 primeiros de cada segmento
drop table if exists ranking.ranking_15;
-- 15 primeiros
CREATE TABLE ranking.ranking_15 AS
SELECT
    codigo_estabelecimento,
    segmento,
    valor_transacao,
    rn,
    '15primeiros'
FROM (
    SELECT
        codigo_estabelecimento,
        segmento,
        valor_transacao,
        ROW_NUMBER() OVER (PARTITION BY segmento ORDER BY valor_transacao DESC) AS rn
    FROM ranking.ranking_group
) ranked
WHERE rn <= 15;

-- seleciona os 200 últimos de cada segmento
drop table if exists ranking.ranking_200;
CREATE TABLE ranking.ranking_200 AS
SELECT
    codigo_estabelecimento,
    segmento,
    valor_transacao,
    '200ultimos'
FROM (
    SELECT
        codigo_estabelecimento,
        segmento,
        valor_transacao,
        ROW_NUMBER() OVER (PARTITION BY segmento ORDER BY valor_transacao ASC) AS rn
    FROM ranking.ranking_group
) ranked
WHERE rn <= 200;


-- inserindo ranking dos 15 primeiros na tabela intermediária
truncate table reports.ranking_ch;
INSERT INTO reports.ranking_ch (
    ano,
    trimestre,
    codigo_estabelecimento,
    funcao,
    bandeira,
    forma_captura,
    numero_parcelas,
    codigo_segmento,
    valor_transacoes,
    quantidade_transacoes,
    taxa_desconto_media
)
SELECT 
    2025 AS ano,
    4 AS trimestre,
    ttt.codigo_estabelecimento::varchar AS codigo_estabelecimento,
    ctg.funcao AS funcao,
    ctg.bandeira::numeric AS bandeira,
    CASE 
        WHEN ctg.forma_captura = '7' THEN 5
        WHEN ctg.forma_captura = '2' THEN 1
        ELSE 2
    END AS forma_captura,
    ctg.numero_parcelas AS numero_parcelas,
    ttt.segmento::int AS codigo_segmento,
    SUM(ctg.sum_valor_transacoes) AS valor_transacoes,
    SUM(ctg.quantidade_transacoes) AS quantidade_transacoes,
    ROUND(SUM(ctg.avg_percentual_desconto / 100 * ctg.sum_valor_transacoes) / SUM(ctg.sum_valor_transacoes) * 100, 2) AS taxa_desconto_media
FROM ranking.ranking_15 ttt
INNER JOIN apoio.gestao ctg 
    ON ctg.codigo_estabelecimento = ttt.codigo_estabelecimento
    and ctg.segmento = ttt.segmento
GROUP BY 
    1,2,3,4,5,6,7,8;

-- inserindo ranking dos 200 últimos na tabela intermediária
INSERT INTO reports.ranking_ch (
    ano,
    trimestre,
    codigo_estabelecimento,
    funcao,
    bandeira,
    forma_captura,
    numero_parcelas,
    codigo_segmento,
    valor_transacoes,
    quantidade_transacoes,
    taxa_desconto_media
)
SELECT 
    2025 AS ano,
    4 AS trimestre,
    'group200' AS codigo_estabelecimento,
    ctg.funcao AS funcao,
    ctg.bandeira::numeric AS bandeira,
    CASE 
        WHEN ctg.forma_captura = '7' THEN 5
        WHEN ctg.forma_captura = '2' THEN 1
        ELSE 2
    END AS forma_captura,
    ctg.numero_parcelas AS numero_parcelas,
    ttt.segmento::int AS codigo_segmento,
    SUM(ctg.sum_valor_transacoes) AS valor_transacoes,
    SUM(ctg.quantidade_transacoes) AS quantidade_transacoes,
    ROUND(SUM(ctg.avg_percentual_desconto / 100 * ctg.sum_valor_transacoes) / SUM(ctg.sum_valor_transacoes) * 100, 2) AS taxa_desconto_media
FROM ranking.ranking_200 ttt
INNER JOIN apoio.gestao ctg 
    ON ctg.codigo_estabelecimento = ttt.codigo_estabelecimento
    and ctg.segmento = ttt.segmento
GROUP BY 
    1,2,3,4,5,6,7,8;
