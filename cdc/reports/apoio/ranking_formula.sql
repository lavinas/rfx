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
 
select * from  public.cadoc_6334_transacao_gestao ctg limit 5;
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
 
create table tmp_tanking_todos as
select *
  from  tmp_ranking_15
    union all
select *
  from tmp_ranking_200;
 