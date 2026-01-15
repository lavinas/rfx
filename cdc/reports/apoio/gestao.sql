 SELECT t06.cd_transacao_fin,
                   t02.cnt_privado,
                   t06.dt_processamento ,
                   t06.nu_valor AS ValorTransacao,
                   CASE
                       WHEN t52.cd_produto_bandeira IN (31,32) THEN '08'
                       WHEN t52.cd_produto_bandeira IN (26,29) THEN '02'
                       WHEN t52.cd_produto_bandeira IN (27,30) THEN '01'
                       ELSE 'N'
                   END AS Bandeira,
                   t06.cd_pessoa_estabelecimento AS CodigoEstabelecimento,
                   t01.cd_mcc  AS mcc ,
                   null  AS segmento  ,
                   t02.tp_entrada as FormaCaptura,
                   CASE t04.cd_tp_cartao
                       WHEN 2 THEN 'C'
                       WHEN 1 THEN 'D'
                       ELSE 'N'
                   END AS Funcao,
                   t06.nu_parcelas AS NumeroParcelas,
                   COALESCE(taxas.TaxaDescontoTotal, 0) as DescontoValor,
                   ROUND((COALESCE(taxas.TaxaDescontoTotal, 0) / NULLIF(t06.nu_valor,0)) * 100, 2) AS PercentualDesconto,
                  null AS cadocItemId 
            FROM efc.tb12006 t06
            INNER JOIN efc.tb04002 t02 on t02.id_sra = t06.id_sra and t02.cd_transacao_sra = t06.cd_transacao_sra
            INNER JOIN efc.tb01023 t23 ON t23.cd_pessoa_estabelecimento = t06.cd_pessoa_estabelecimento
            INNER join efc.tb01001 t01 on t01.cd_atividade  = t23.cd_atividade   
            INNER JOIN efc.tb12004 t04 ON t04.cd_produto_cartao_adquirente = t06.cd_produto_cartao_adquirente
            INNER JOIN efc.tb12052 t52 ON t52.cd_produto_bandeira = t04.cd_produto_bandeira
            LEFT JOIN (
                SELECT cd_transacao_fin, SUM(nu_valor_taxa) AS TaxaDescontoTotal
                FROM efc.tb12002
                GROUP BY cd_transacao_fin
            ) taxas ON taxas.cd_transacao_fin = t06.cd_transacao_fin
            WHERE t06.dt_processamento >= ? 
            AND t06.dt_processamento < ?
            AND t06.tp_origem = 1
            AND t06.cd_rede_adquirente = 1;   -- public.cadoc_6334_transacao_gestao definição

-- Drop table

-- DROP TABLE apoio.cadoc_6334_transacao_gestao;

CREATE TABLE apoio.cadoc_6334_transacao_gestao (
 cd_transacao_fin int8 NOT NULL,
 dt_processamento timestamp NULL,
 valor_transacao numeric(38, 2) NULL,
 bandeira varchar(255) NULL,
 codigo_estabelecimento int8 NULL,
 mcc varchar(4) NULL,
 segmento int4 NULL,
 forma_captura varchar(255) NULL,
 funcao varchar(255) NULL,
 numero_parcelas int4 NULL,
 percentual_desconto numeric(38, 2) NULL,
 taxa_desconto_total numeric(38, 2) NULL,
 cadoc_item_id int8 NULL,
 cnt_privado varchar(255) NULL,
 codigo_segmento varchar(255) NULL,
 qttd int8 DEFAULT 1 NULL,
 tipo varchar(1) DEFAULT 'N'::character varying NULL,
 CONSTRAINT transacao_financeira_gestao_pkey PRIMARY KEY (cd_transacao_fin)
 -- CONSTRAINT cadoc_transacao_item_fk FOREIGN KEY (cadoc_item_id) REFERENCES apoio.cadoc_6334_item(id)
);


select dt_processamento,
       bandeira,
       mcc,
       segmento,
       forma_captura,
       funcao,
       numero_parcelas,
       codigo_segmento,
       sum(valor_transacao) as sum_valor_transacoes,
       count(1) as quantidade_transacoes,
       sum(percentual_desconto) as sum_percentual_desconto,
       avg(percentual_desconto) as avg_percentual_desconto,
       min(percentual_desconto) as min_percentual_desconto,
       max(percentual_desconto) as max_percentual_desconto,
       stddev_samp(percentual_desconto) as dev_percentual_desconto,
       sum(taxa_desconto_total) as sum_taxa_desconto_total,
       avg(taxa_desconto_total) as avg_taxa_desconto_total,
       min(taxa_desconto_total) as min_taxa_desconto_total,
       max(taxa_desconto_total) as max_taxa_desconto_total,
       stddev_samp(taxa_desconto_total) as dev_taxa_desconto_total
   from apoio.cadoc_6334_transacao_gestao
group by 1,2,3,4,5,6,7,8;


-- gestao
select bandeira,
       codigo_estabelecimento,
       mcc,
       segmento,
       forma_captura,
       funcao,
       numero_parcelas,
       codigo_segmento,
       sum(valor_transacao) as sum_valor_transacoes,
       count(1) as quantidade_transacoes,
       sum(percentual_desconto) as sum_percentual_desconto,
       avg(percentual_desconto) as avg_percentual_desconto,
       min(percentual_desconto) as min_percentual_desconto,
       max(percentual_desconto) as max_percentual_desconto,
       stddev_samp(percentual_desconto) as dev_percentual_desconto,
       sum(taxa_desconto_total) as sum_taxa_desconto_total,
       avg(taxa_desconto_total) as avg_taxa_desconto_total,
       min(taxa_desconto_total) as min_taxa_desconto_total,
       max(taxa_desconto_total) as max_taxa_desconto_total,
       stddev_samp(taxa_desconto_total) as dev_taxa_desconto_total
   from apoio.cadoc_6334_transacao_gestao
where dt_processamento BETWEEN '2025-10-01' and '2026-01-01'
group by 1,2,3,4,5,6,7,8;