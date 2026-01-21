SELECT
    'INSERT INTO cadoc_6334_intercam (Ano, Trimestre, Produto, ModalidadeCartao, Funcao, Bandeira, FormaCaptura, NumeroParcelas, CodigoSegmento, TarifaIntercambio, ValorTransacoes, QuantidadeTransacoes) VALUES (' ||
    COALESCE(ano::text, 'NULL') || ', ' ||
    COALESCE(trimestre::text, 'NULL') || ', ' ||
    COALESCE(produto::text, 'NULL') || ', ' ||
    COALESCE(quote_literal(modalidade_cartao), 'NULL') || ', ' ||
    COALESCE(quote_literal(funcao), 'NULL') || ', ' ||
    COALESCE(bandeira::text, 'NULL') || ', ' ||
    COALESCE(forma_captura::text, 'NULL') || ', ' ||
    COALESCE(numero_parcelas::text, 'NULL') || ', ' ||
    COALESCE(codigo_segmento::text, 'NULL') || ', ' ||
    COALESCE(replace(CAST(tarifa_intercambio AS text), ',', '.'), 'NULL') || ', ' ||
    COALESCE(replace(CAST(valor_transacoes AS text), ',', '.'), 'NULL') || ', ' ||
    COALESCE(quantidade_transacoes::text, 'NULL') || ');'
FROM reports.intercam_ch
limit 1000000;

select count(*) from reports.intercam_ch;