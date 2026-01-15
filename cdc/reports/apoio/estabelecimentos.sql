    SELECT
              t09.cd_pessoa as codigo_estabelecimento,
              t09.dt_cadastro as data_credenciamento,
              t09.nm_pessoa as razao_social,
              ultima.dt_processamento as data_ultima_transacao,
              t27.nu_cnpj  AS cnpj,
              t25.nu_cpf   AS cpf,
              t22.cd_uf    AS uf,
              (MAX(CASE WHEN t04.cd_tp_cartao = 2 THEN 1 ELSE 0 END) > 0) AS tem_credito,
              (MAX(CASE WHEN t04.cd_tp_cartao = 1 THEN 1 ELSE 0 END) > 0) AS tem_debito,
              (MAX(CASE WHEN t52.produto_intercambio = 'MASTERCARD' THEN 1 ELSE 0 END) > 0) AS tem_mastercard,
              (MAX(CASE WHEN t52.produto_intercambio = 'VISA' THEN 1 ELSE 0 END) > 0) AS tem_visa,
              (MAX(CASE WHEN t52.produto_intercambio = 'ELO' THEN 1 ELSE 0 END) > 0) AS tem_elo,
              t01.cd_mcc,
              NULL::text AS segmento,
              FALSE AS captura_manual,
              TRUE  AS captura_eletronica,
              FALSE AS captura_remota,
              NULL AS cadoc_item_id
            FROM efc.tb01009 t09
            INNER JOIN  efc.tb01023 t23 ON t23.cd_pessoa_estabelecimento = t09.cd_pessoa
            INNER JOIN  efc.tb01001 t01 ON t01.cd_atividade = t23.cd_atividade
            LEFT JOIN  efc.tb01025 t25 ON t25.cd_pessoa = t23.cd_pessoa_estabelecimento
            LEFT JOIN  efc.tb01027 t27 ON t27.cd_pessoa = t23.cd_pessoa_estabelecimento
            INNER JOIN  efc.tb12008 t08 ON t08.cd_pessoa_estabelecimento = t09.cd_pessoa
            INNER JOIN  efc.tb12004 t04 ON t04.cd_produto_cartao_adquirente = t08.cd_produto_cartao_adquirente
            INNER JOIN  efc.tb12052 t52 ON t52.cd_produto_bandeira = t04.cd_produto_bandeira
            INNER JOIN  efc.tb01022 t22 ON t22.cd_pessoa = t23.cd_pessoa_estabelecimento
            LEFT JOIN (
              SELECT DISTINCT ON (cd_pessoa_estabelecimento)
                cd_pessoa_estabelecimento,
                dt_processamento
              FROM  efc.tb12006
              ORDER BY cd_pessoa_estabelecimento, dt_processamento DESC
            ) ultima ON ultima.cd_pessoa_estabelecimento = t09.cd_pessoa
            GROUP BY
              t09.cd_pessoa,
              t09.dt_cadastro,
              ultima.dt_processamento,
              t27.nu_cnpj,
              t25.nu_cpf,
              t22.cd_uf,
              t01.cd_mcc
            ;

 -- public.cadoc_6334_estabelecimento definição

-- Drop table

-- DROP TABLE public.cadoc_6334_estabelecimento;

CREATE TABLE apoio.cadoc_6334_estabelecimento (
 codigo_estabelecimento int8 NOT NULL,
 data_credenciamento timestamp NULL,
 data_ultima_transacao timestamp NULL,
 razao_social varchar(255) NULL,
 cnpj varchar(255) NULL,
 cpf varchar(255) NULL,
 uf varchar(255) NULL,
 tem_debito bool NULL,
 tem_credito bool NULL,
 tem_visa bool NULL,
 tem_mastercard bool NULL,
 tem_elo bool NULL,
 mcc varchar(255) NULL,
 segmento int4 NULL,
 captura_manual bool NULL,
 captura_eletronica bool NULL,
 captura_remota bool NULL,
 cadoc_item_id int8 NULL,
 CONSTRAINT cadoc_6334_estabelecimento_pkey PRIMARY KEY (codigo_estabelecimento)
 -- CONSTRAINT cadoc_6334_estabelecimento_cadoc_item_id_fkey FOREIGN KEY (cadoc_item_id) REFERENCES public.cadoc_6334_item(id)
);

-- estabelecimentos
select codigo_estabelecimento,
       data_credenciamento,
       data_ultima_transacao,
       uf,
       tem_debito,
       tem_credito,
       tem_visa,
       tem_mastercard,
       tem_elo,
       mcc,
       segmento,
       captura_manual,
       captura_eletronica,
       captura_remota
  from apoio.cadoc_6334_estabelecimento;