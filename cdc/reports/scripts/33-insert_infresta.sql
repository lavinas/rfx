SELECT
    'INSERT INTO cadoc_6334_infresta ("Ano","Trimestre","Uf","QuantidadeEstabelecimentosTotais","QuantidadeEstabelecimentosCapturaManual","QuantidadeEstabelecimentosCapturaEletronica","QuantidadeEstabelecimentosCapturaRemota") VALUES (' ||
    ano::text || ', ' ||
    trimestre::text || ', ' ||
    quote_literal(uf) || ', ' ||
    quantidade_estabelecimentos_totais::text || ', ' ||
    quantidade_estabelecimentos_captura_manual::text || ', ' ||
    quantidade_estabelecimentos_captura_eletronica::text || ', ' ||
    quantidade_estabelecimentos_captura_remota::text || ');'
AS insert_sql
FROM reports.infresta_ch
limit 100000;


select count(1) from reports.infresta_ch;