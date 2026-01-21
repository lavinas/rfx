CREATE TABLE public.tmp_ranking AS
SELECT
    ctg.codigo_estabelecimento,
    min(csmc.segment) as segment,
    SUM(ctg.valor_transacao) AS valor_transacao
FROM public.cadoc_6334_transacao_gestao ctg
JOIN public.cadoc_6334_segmento_mcc csmc
    ON csmc.mcc_init::int = ctg.mcc::int
  --  WHERE ctg.tipo <> 'A' OR ctg.tipo IS NULL
GROUP BY 1;
 
CREATE TABLE public.tmp_ranking_15 AS
SELECT
    codigo_estabelecimento,
    segment,
    valor_transacao,
    '15primeiros'
FROM (
    SELECT
        codigo_estabelecimento,
        segment,
        valor_transacao,
        ROW_NUMBER() OVER (PARTITION BY segment ORDER BY valor_transacao DESC) AS rn
    FROM public.tmp_ranking
) ranked
WHERE rn <= 15;
 

SELECT
        codigo_estabelecimento,
        segment,
        valor_transacao,
        ROW_NUMBER() OVER (PARTITION BY segment ORDER BY valor_transacao DESC) AS rn
    FROM public.tmp_ranking;


CREATE TABLE public.tmp_ranking_200 AS
SELECT
    codigo_estabelecimento,
    segment,
    valor_transacao,
    '200ultimos'
FROM (
    SELECT
        codigo_estabelecimento,
        segment,
        valor_transacao,
        ROW_NUMBER() OVER (PARTITION BY segment ORDER BY valor_transacao ASC) AS rn
    FROM public.tmp_ranking
) ranked
WHERE rn <= 200;
 
create table public.tmp_ranking_todos as
select *
  from  public.tmp_ranking_15
    union all
select *
  from public.tmp_ranking_200;
 


INSERT INTO public.cadoc_6334_ranking_intermediaria (
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
-- PRIMEIROS 15
SELECT 
    2025 AS ano,
    2 AS trimestre,
    ctg.codigo_estabelecimento::varchar AS codigo_estabelecimento,
    ctg.funcao AS funcao,
    ctg.bandeira::numeric  AS bandeira,
    CASE 
        WHEN ctg.forma_captura = '7' THEN 5
        WHEN ctg.forma_captura = '2' THEN 1
        ELSE 2
    END AS forma_captura,
    ctg.numero_parcelas AS numero_parcelas,
    ttt.segment AS codigo_segmento,
    SUM(ctg.valor_transacao) AS valor_transacoes,
    SUM(ctg.qttd) AS quantidade_transacoes,
    ROUND(
        SUM(ctg.taxa_desconto_total) / NULLIF(SUM(ctg.valor_transacao), 0),
        4
    ) * 100 AS taxa_desconto_media
FROM public.tmp_tanking_todos ttt
INNER JOIN public.cadoc_6334_transacao_gestao ctg 
    ON ctg.codigo_estabelecimento = ttt.codigo_estabelecimento
WHERE ttt."?column?" = '15primeiros'
GROUP BY 
    1,2,3,4,5,6,7,8

UNION ALL

-- 200 ÚLTIMOS (agrupado em 'group200')
SELECT 
    2025 AS ano,
    2 AS trimestre,
    'group200' AS codigo_estabelecimento,
    ctg.funcao AS funcao,
    ctg.bandeira::numeric  AS bandeira,
    CASE 
        WHEN ctg.forma_captura = '7' THEN 5
        WHEN ctg.forma_captura = '2' THEN 1
        ELSE 2
    END AS forma_captura,
    ctg.numero_parcelas AS numero_parcelas,
    ttt.segment AS codigo_segmento,
    SUM(ctg.valor_transacao) AS valor_transacoes,
    SUM(ctg.qttd) AS quantidade_transacoes,
    ROUND(
        SUM(ctg.taxa_desconto_total) / NULLIF(SUM(ctg.valor_transacao), 0),
        4
    ) * 100 AS taxa_desconto_media
FROM public.tmp_tanking_todos ttt
INNER JOIN public.cadoc_6334_transacao_gestao ctg 
    ON ctg.codigo_estabelecimento = ttt.codigo_estabelecimento
WHERE ttt."?column?" = '200ultimos'
GROUP BY 
    1,2,3,4,5,6,7,8;