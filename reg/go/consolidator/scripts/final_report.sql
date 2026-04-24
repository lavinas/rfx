
-- extract intercam data for final report
select id, 
       year as ano, 
       quarter as trimestre, 
       product_code as produto,
       card_type as modalidade_cartao,
       function as funcao,
       brand as bandeira,
       capture_mode as forma_captura,
       installments as numero_parcelas,
       segment_code as codigo_segmento,
       replace(interchange_fee::text, '.', ',') as "tarifa_intercambio (%)",
       replace(transaction_amount::text, '.', ',') as valor_transacoes,
       transaction_quantity as quantidade_transacoes
  from cadoc_6334_v2.intercam
where year = 2026 and quarter = 1
limit 10000000;

-- intercam totals to compare with excel report
select round(sum(interchange_fee*transaction_amount / 100) / sum(transaction_amount) * 100, 2)  as tarifa_intercambio,
       sum(transaction_amount) as valor_transacoes,
       sum(transaction_quantity) as quantidade_transacoes,
       count(1) linhas
  from cadoc_6334_v2.intercam
where year = 2026 and quarter = 1;


-- extract desconto data for final report
select id, 
       year as ano, 
       quarter as trimestre, 
       function as funcao,
       brand as bandeira,
       capture_mode as forma_captura,
       installments as numero_parcelas,
       segment_code as codigo_segmento,
       replace(min_mdr_fee::text, '.', ',') as "taxa_desconto_minima",
       replace(max_mdr_fee::text, '.', ',') as "taxa_desconto_maxima",
       replace(stdev_mdr_fee::text, '.', ',') as "desvio_padrao_taxa_desconto",
       replace(avg_mdr_fee::text, '.', ',') as "taxa_desconto_media",
       replace(transaction_amount::text, '.', ',') as valor_transacoes,
       transaction_quantity as quantidade_transacoes
  from cadoc_6334_v2.desconto
  where year = 2026 and quarter = 1
limit 10000000;

-- desconto totals to compare with excel report
select round(sum(avg_mdr_fee*transaction_amount / 100) / sum(transaction_amount) * 100, 2)  as taxa_desconto_media,
       sum(transaction_amount) as valor_transacoes,
       sum(transaction_quantity) as quantidade_transacoes,
       count(1) linhas
  from cadoc_6334_v2.desconto
where year = 2026 and quarter = 1;

-- extract ranking data for final report
select id, 
       year as ano,
       quarter as trimestre,
       replace(establishment_code::text, '-1', 'group200') as codigo_estabelecimento,
       function as funcao,
       brand as bandeira,
       capture_mode as forma_captura,
       installments as numero_parcelas,
       segment_code as codigo_segmento,
       replace(avg_mcc_fee::text, '.', ',') as "taxa_mcc_media",
       replace(transaction_amount::text, '.', ',') as valor_transacoes,
       transaction_quantity as quantidade_transacoes
    from cadoc_6334_v2.ranking_filtered
limit 10000000;

-- ranking totals to compare with excel report
select round(sum(avg_mcc_fee*transaction_amount / 100) / sum(transaction_amount) * 100, 2)  as taxa_mcc_media,
       sum(transaction_amount) as valor_transacoes,
       sum(transaction_quantity) as quantidade_transacoes,
       count(1) linhas
  from cadoc_6334_v2.ranking_filtered
where year = 2026 and quarter = 1;


-- extract luccred data for final report
select id,
       year as ano,
       quarter as trimestre,
       replace(gross_revenue::text, '.', ',') as receita_taxa_desconto_bruta,
       replace(rental_revenue::text, '.', ',') as receita_aluguel_equipamentos_conectividade,
       replace(others_revenue::text, '.', ',') as receita_outras,
       replace(interchange_cost::text, '.', ',') as custo_tarifa_intercambio,
       replace(marketing_cost::text, '.', ',') as custo_marketing_propaganda,
       replace(brand_access_cost::text, '.', ',') as custo_taxas_acesso_bandeiras,
       replace(risk_cost::text, '.', ',') as custo_risco,
       replace(processing_cost::text, '.', ',') as custo_processamento,
       replace(others_cost::text, '.', ',') as custos_outros
    from cadoc_6334_v2.luccred
where year = 2026 and quarter = 1
limit 10000000;


-- luccred totals to compare with excel report
select gross_revenue + rental_revenue + others_revenue - interchange_cost - marketing_cost - brand_access_cost - risk_cost - processing_cost - others_cost as resultado
  from cadoc_6334_v2.luccred
where year = 2026 and quarter = 1;


-- extract conccred data for final report
select id,
       year as ano,
       quarter as trimestre,
       brand as bandeira,
       function as funcao,
       number_accredited_establishments as credenciados,
       number_active_establishments as ativos,
       replace(transaction_amount::text, '.', ',') as valor_transacoes,
       transaction_quantity as quantidade_transacoes
    from cadoc_6334_v2.conccred
where year = 2026 and quarter = 1
limit 10000000;

-- conccred totals to compare with excel report
select max(number_accredited_establishments) as credenciados,
       max(number_active_establishments) as ativos,
       sum(transaction_amount) as valor_transacoes,
       sum(transaction_quantity) as quantidade_transacoes,
       count(1) linhas
  from cadoc_6334_v2.conccred
where year = 2026 and quarter = 1;



-- extract infresta data for final report
select id,
       year as ano,
       quarter as trimestre,
       federation_unit as uf,
       establishment_manual_capture_quantity as captura_manual,
       establishment_remote_capture_quantity as captura_remota,
       establishment_eletronic_capture_quantity as captura_eletronica,
       establishment_total_quantity as totais
    from cadoc_6334_v2.infresta
where year = 2026 and quarter = 1;

-- infresta totals to compare with excel report
select sum(establishment_manual_capture_quantity) as captura_manual,
       sum(establishment_remote_capture_quantity) as captura_remota,
       sum(establishment_eletronic_capture_quantity) as captura_eletronica,
       sum(establishment_total_quantity) as totais,
       count(1) linhas
  from cadoc_6334_v2.infresta
where year = 2026 and quarter = 1;


-- extract infrterm data for final report
select id,
         year as ano,
         quarter as trimestre,
         federation_unit as uf,
         pos_total_quantity pos,
         pos_shared_quantity pos_compartilhado,
         pos_chip_quantity pos_leitor_chip,
         pdv_quantity pdv
  from cadoc_6334_v2.infrterm a
where year = 2026 and quarter = 1;


-- infrterm totals to compare with excel report
select sum(pos_total_quantity) pos,
       sum(pos_shared_quantity) pos_compartilhado,
       sum(pos_chip_quantity) pos_leitor_chip,
       sum(pdv_quantity) pdv,
       count(1) linhas
  from cadoc_6334_v2.infrterm a
where year = 2026 and quarter = 1;


select id,
       segment_name as nome,
       segment_description as descricao,
       segment_code as codigo
  from cadoc_6334_v2.segmento a
where year = 2026 and quarter = 1
  order by 4
    limit 10000000;


-- extract contato data for final report
select id,
       year as ano,
       quarter as trimestre,
       contact_type as tipo_contato,
       name as nome,
       position as cargo,
       phone_number as telefone,
       email as email
  from cadoc_6334_v2.contatos a
where year = 2026 and quarter = 1
order by 1
limit 10000000;



-- detalhes

-- funcao

select function as funcao, 
       replace(sum(transaction_amount)::text, '.', ',') as valor,
       sum(transaction_quantity) as quantidade
  from cadoc_6334_v2.intercam
where year = 2026 and quarter = 1
group by 1
order by 1 desc;

select function as funcao, 
       replace(sum(transaction_amount)::text, '.', ',') as valor,
       sum(transaction_quantity) as quantidade
  from cadoc_6334_v2.desconto
where year = 2026 and quarter = 1
group by 1
order by 1 desc;


-- bandeira

select brand as bandeira, 
       replace(sum(transaction_amount)::text, '.', ',') as valor,
       sum(transaction_quantity) as quantidade
  from cadoc_6334_v2.intercam
where year = 2026 and quarter = 1
group by 1
order by 1;

select brand as bandeira, 
       replace(sum(transaction_amount)::text, '.', ',') as valor,
       sum(transaction_quantity) as quantidade
  from cadoc_6334_v2.desconto
where year = 2026 and quarter = 1
group by 1
order by 1;



-- parcela

select installments as parcela, 
       replace(sum(transaction_amount)::text, '.', ',') as valor,
       sum(transaction_quantity) as quantidade
  from cadoc_6334_v2.intercam
where year = 2026 and quarter = 1
group by 1
order by 1;

select installments as parcela, 
       replace(sum(transaction_amount)::text, '.', ',') as valor,
       sum(transaction_quantity) as quantidade
  from cadoc_6334_v2.desconto
where year = 2026 and quarter = 1
group by 1
order by 1;


-- segmento

select segment_code as segmento, 
       replace(sum(transaction_amount)::text, '.', ',') as valor,
       sum(transaction_quantity) as quantidade
  from cadoc_6334_v2.intercam
where year = 2026 and quarter = 1
group by 1
order by 1;

select segment_code as segmento, 
       replace(sum(transaction_amount)::text, '.', ',') as valor,
       sum(transaction_quantity) as quantidade
  from cadoc_6334_v2.desconto
where year = 2026 and quarter = 1
group by 1
order by 1;


-- captura

select capture_mode as segmento, 
       replace(sum(transaction_amount)::text, '.', ',') as valor,
       sum(transaction_quantity) as quantidade
  from cadoc_6334_v2.intercam
where year = 2026 and quarter = 1
group by 1
order by 1;

select capture_mode as segmento, 
       replace(sum(transaction_amount)::text, '.', ',') as valor,
       sum(transaction_quantity) as quantidade
  from cadoc_6334_v2.desconto
where year = 2026 and quarter = 1
group by 1
order by 1;


-- validacoes finais

select a.brand, 
       a.function, 
       round(sum(case when product_code = 38 then transaction_quantity else 0 end) / sum(transaction_quantity) * 100, 2) as percentual_transacoes_produto_38,
       count(1) 
  from cadoc_6334_v2.intercam a
group by 1, 2
order by 3 desc;

