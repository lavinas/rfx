SELECT
    'INSERT INTO cadoc_6334_desconto (Ano, Trimestre, Funcao, Bandeira, FormaCaptura, NumeroParcelas, CodigoSegmento, TaxaDescontoMedia, TaxaDescontoMinima, TaxaDescontoMaxima, DesvioPadraoTaxaDesconto, ValorTransacoes, QuantidadeTransacoes) VALUES (' ||
    COALESCE(ano::text, 'NULL') || ', ' ||
    COALESCE(trimestre::text, 'NULL') || ', ' ||
    COALESCE(quote_literal(trim(funcao)), 'NULL') || ', ' ||
    COALESCE(bandeira::text, 'NULL') || ', ' ||
    COALESCE(forma_captura::text, 'NULL') || ', ' ||
    COALESCE(numero_parcelas::text, 'NULL') || ', ' ||
    COALESCE(codigo_segmento::text, 'NULL') || ', ' ||
    COALESCE(taxa_desconto_media::text, 'NULL') || ', ' ||
    COALESCE(taxa_desconto_minima::text, 'NULL') || ', ' ||
    COALESCE(taxa_desconto_maxima::text, 'NULL') || ', ' ||
    COALESCE(desvio_padrao_taxa_desconto::text, 'NULL') || ', ' ||
    COALESCE(valor_transacoes::text, 'NULL') || ', ' ||
    COALESCE(quantidade_transacoes::text, 'NULL') || ');'
AS insert_stmt
FROM reports.descontos_ch
LIMIT 1000000;


select count(*) from reports.descontos_ch;