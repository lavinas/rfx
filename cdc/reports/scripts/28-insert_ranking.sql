SELECT
    'INSERT INTO cadoc_6334_ranking (Ano, Trimestre, CodigoEstabelecimento, Funcao, Bandeira, FormaCaptura, NumeroParcelas, CodigoSegmento, ValorTransacoes, QuantidadeTransacoes, TaxaDescontoMedia) VALUES ('
    ||
    COALESCE(ano::text, 'NULL') || ', '
    ||
    COALESCE(trimestre::text, 'NULL') || ', '
    ||
    (CASE WHEN codigo_estabelecimento IS NULL THEN 'NULL'
                ELSE '''' || replace(left(codigo_estabelecimento,8), '''', '''''') || '''' END) || ', '
    ||
    (CASE WHEN funcao IS NULL THEN 'NULL'
                ELSE '''' || replace(funcao::text, '''', '''''') || '''' END) || ', '
    ||
    COALESCE(bandeira::text, 'NULL') || ', '
    ||
    COALESCE(forma_captura::text, 'NULL') || ', '
    ||
    COALESCE(numero_parcelas::text, 'NULL') || ', '
    ||
    COALESCE(codigo_segmento::text, 'NULL') || ', '
    ||
    (CASE WHEN valor_transacoes IS NULL THEN 'NULL'
                ELSE replace(valor_transacoes::text, ',', '.') END) || ', '
    ||
    COALESCE(quantidade_transacoes::text, 'NULL') || ', '
    ||
    (CASE WHEN taxa_desconto_media IS NULL THEN 'NULL'
                ELSE replace(taxa_desconto_media::text, ',', '.') END)
    ||
    ');' AS insert_stmt
FROM reports.ranking_ch
limit 1000000;


select count(1)
from reports.ranking_ch;