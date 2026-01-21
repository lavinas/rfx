SELECT
    'INSERT INTO cadoc_6334_conccred (Ano, Trimestre, Bandeira, Funcao, QuantidadeEstabelecimentosCredenciados, QuantidadeEstabelecimentosAtivos, ValorTransacoes, QuantidadeTransacoes) VALUES (' ||
    ano::text || ', ' ||
    trimestre::text || ', ' ||
    COALESCE(bandeira::text, '0') || ', ' ||
    quote_literal(funcao) || ', ' ||
    COALESCE(quantidade_estabelecimentos_credenciados::text, '0') || ', ' ||
    COALESCE(quantidade_estabelecimentos_ativos::text, '0') || ', ' ||
    COALESCE(to_char(valor_transacoes, 'FM999999999990.00'), '0.00') || ', ' ||
    COALESCE(quantidade_transacoes::text, '0') ||
    ');'
FROM reports.conccred_ch
ORDER BY id;

select count(1) from reports.conccred_ch;