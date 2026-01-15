create schema apoio;

 SELECT 
    fc.brand_transaction_identifier AS cd_transacao_fin,
    fc.pos_entry_mode AS forma_captura,
    fc.transaction_date AS dt_processamento,
    fc.transaction_amount AS ValorTransacoes,
    CASE 
        WHEN fci.exchange_rate_type_name = 'PERCENT'
            THEN fci.exchange_rate_amount
        WHEN fci.exchange_rate_type_name = 'FIXED'
            THEN (fci.exchange_rate_amount / NULLIF(fc.transaction_amount,0)) * 100
    END AS percentual_desconto,
    CASE 
        WHEN fci.exchange_rate_type_name = 'PERCENT'
            THEN ROUND(fc.transaction_amount * fci.exchange_rate_amount / 100, 2)
        WHEN fci.exchange_rate_type_name = 'FIXED'
            THEN fci.exchange_rate_amount
    END AS taxa_intercambio_valor,
    fc.brand_name AS bandeira,
    fci.plan AS parcela,
    fc.product_name AS tipo_cartao,
    fc.merchant_category_code AS segmento,
    fc.mask_account_number AS bin,
    NULL AS modalidade_cartao,
    NULL AS produto_cartao,
    NULL AS cadoc_item_id,
    CASE
        WHEN "left"(fc.brand_name::text, 1) = 'M'::text THEN ("left"(fc.brand_transaction_identifier::text, 9) || to_char(fc.transaction_date::timestamp with time zone, 'MMDD'::text))::character varying
            ELSE fc.brand_transaction_identifier
        END::text AS chave_join_inter
FROM interc.fin_contract fc
JOIN interc.fin_contract_installment fci
  ON fci.contract_id = fc.id
 AND fci.installment_number = 1
WHERE fc.transaction_date BETWEEN ? AND ?;
-- public.cadoc_6334_transacao_intercambio definição
-- Drop table
-- DROP TABLE public.cadoc_6334_transacao_intercambio;
CREATE TABLE apoio.cadoc_6334_transacao_intercambio (
 cd_transacao_fin varchar(255) NOT NULL,
 bin varchar(255) NULL,
 cadoc_item_id int4 NULL,
 dt_processamento date NULL,
 modalidade_cartao varchar(255) NULL,
 percentual_desconto numeric(38, 2) NULL,
 produto_cartao varchar(255) NULL,
 taxa_intercambio_valor numeric(38, 2) NULL,
 valor_transacao numeric(38, 2) NULL,
 bandeira varchar(255) NULL,
 parcela int4 NULL,
 tipo_cartao varchar(255) NULL,
 forma_captura varchar(255) NULL,
 segmento varchar(255) NULL,
 qttd int8 DEFAULT 1 NULL,
 tipo varchar(1) DEFAULT 'N'::character varying NULL,
 chave_join_inter varchar(255) NULL,
 CONSTRAINT cadoc_6334_transacao_intercambio_pkey PRIMARY KEY (cd_transacao_fin)
);
  
-- intercambio
select  cdc.bandeira,
        cdc.parcela,
        cdc.tipo_cartao,
        cdc.forma_captura,
        cdc.segmento,
        cdc.bin,
        sum(cdc.valor_transacao) as sum_valor_transacoes,
        count(1) as quantidade_transacoes,
        sum(cdc.percentual_desconto) as sum_percentual_desconto,
        avg(cdc.percentual_desconto) as avg_percentual_desconto,
        min(cdc.percentual_desconto) as min_percentual_desconto,
        max(cdc.percentual_desconto) as max_percentual_desconto,
        stddev_samp(cdc.percentual_desconto) as dev_percentual_desconto,
        sum(cdc.taxa_intercambio_valor) as sum_taxa_intercambio_valor,
        avg(cdc.taxa_intercambio_valor) as avg_taxa_intercambio_valor,
        min(cdc.taxa_intercambio_valor) as min_taxa_intercambio_valor,
        max(cdc.taxa_intercambio_valor) as max_taxa_intercambio_valor,
        stddev_samp(cdc.taxa_intercambio_valor) as dev_taxa_intercambio_valor
    from apoio.cadoc_6334_transacao_intercambio cdc
where cdc.dt_processamento BETWEEN '2025-10-01' and '2026-01-01'
group by 1,2,3,4,5,6;
