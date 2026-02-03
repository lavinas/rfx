-- intercam
select id, 
       ano, 
       trimestre, 
       produto, 
       modalidade_cartao, 
       funcao,
       bandeira, 
       forma_captura, 
       numero_parcelas, 
       codigo_segmento, 
       replace(round(tarifa_intercambio,2)::text, '.', ',') as tarifa_intercambio,
       replace(round(valor_transacoes,2)::text, '.', ',') as valor_transacoes, 
       quantidade_transacoes 
    from intercam_ch
    limit 1000000;

select count(*) as total_records from intercam_ch;

-- desconto
select id, 
       ano, 
       trimestre, 
       funcao,
       bandeira, 
       forma_captura, 
       numero_parcelas, 
       codigo_segmento,
       replace(round(taxa_desconto_minima,2)::text, '.', ',') as taxa_desconto_minima,
       replace(round(taxa_desconto_maxima,2)::text, '.', ',') as taxa_desconto_maxima,
       replace(round(desvio_padrao_taxa_desconto,2)::text, '.', ',') as desvio_padrao_taxa_desconto,
       replace(round(taxa_desconto_media,2)::text, '.', ',') as taxa_desconto_media,
       replace(round(valor_transacoes,2)::text, '.', ',') as valor_transacoes,
       quantidade_transacoes
    from descontos_ch
    limit 1000000;


select count(*) as total_records from descontos_ch;

-- ranking
select id +56787 as id, 
       ano, 
       trimestre,
       codigo_estabelecimento,
       funcao,
       bandeira,
       forma_captura,
       numero_parcelas,
       codigo_segmento,
       replace(round(taxa_desconto_media, 2)::text, '.', ',') as taxa_desconto_media,
       replace(round(valor_transacoes, 2)::text, '.', ',') as valor_transacoes,
       quantidade_transacoes 
    from ranking_ch
    limit 1000000;

select count(*) as total_records from ranking_ch;


-- luccred
select id,
       ano, 
       trimestre,
       replace(round(receitataxadescontobruta, 2)::text, '.', ',') as receitataxadescontobruta,
       replace(round(receitaaluguelequipamentosconectividade, 2)::text, '.', ',') as receitaaluguelequipamentosconectividade,
       replace(round(receitaoutras, 2)::text, '.', ',') as receitaoutras,
       replace(round(custotarifaintercambio, 2)::text, '.', ',') as custotarifaintercambio,
       replace(round(customarketingpropaganda, 2)::text, '.', ',') as customarketingpropaganda,
       replace(round(custotaxasacessobandeiras, 2)::text, '.', ',') as custotaxasacessobandeiras,
       replace(round(custorisco, 2)::text, '.', ',') as custorisco,
       replace(round(custoprocessamento, 2)::text, '.', ',') as custoprocessamento,
       replace(round(custooutros, 2)::text, '.', ',') as custooutros
  from luccred_ch
  limit 1000000;



  -- conccred
select id, 
       ano, 
       trimestre,
       bandeira,
       funcao,
       quantidade_estabelecimentos_credenciados,
       quantidade_estabelecimentos_ativos,
       replace(round(valor_transacoes, 2)::text, '.', ',') as valor_transacoes,
       quantidade_transacoes
  from conccred_ch
  limit 1000000;


-- infresta
select id,
       ano,
       trimestre,
       uf,
       quantidade_estabelecimentos_captura_manual,
       quantidade_estabelecimentos_captura_remota,
       quantidade_estabelecimentos_captura_eletronica,
       quantidade_estabelecimentos_totais
  from infresta_ch
  limit 1000000;

  -- infrtern
  select id,
        ano,
        trimestre,
        uf,
        quantidade_total,
        quantidade_pos_compartilhados,
        quantidade_pos_leitora_chip,
        quantidade_pdv
    from infrterm_ch
    limit 1000000;


-- segmentos
select id,
       nome_segmento,
       descricao_segmento,
       codigo_segmento
  from segmentos_ch
  limit 1000000;


select id,
       ano,
       trimestre,
       tipocontato,
       nome,
       cargo,
       numerotelefone,
       email
  from contatos_ch;


select funcao,
       replace(sum(valor_transacoes)::text, '.', ',') as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
  from intercam_ch
group by funcao
order by 1 desc;


select bandeira,
       replace(sum(valor_transacoes)::text, '.', ',') as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
  from intercam_ch
group by bandeira
order by 1;


select numero_parcelas,
       replace(sum(valor_transacoes)::text, '.', ',') as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
  from intercam_ch
group by numero_parcelas
order by 1;


select codigo_segmento,
       replace(sum(valor_transacoes)::text, '.', ',') as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
  from intercam_ch
group by codigo_segmento
order by 1;


select forma_captura,
       replace(sum(valor_transacoes)::text, '.', ',') as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
  from intercam_ch
group by forma_captura
order by 1;