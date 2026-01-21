-- Executar no Postgres: gera INSERTs para a segunda tabela (SQL Server)
SELECT
    'INSERT INTO cadoc_6334_contatos (Ano, Trimestre, TipoContato, Nome, Cargo, Telefone, Email) VALUES ('
    || COALESCE(ano::text, 'NULL') || ', '
    || COALESCE(trimestre::text, 'NULL') || ', '
    || COALESCE(quote_literal(tipocontato), 'NULL') || ', '
    || COALESCE(quote_literal(nome), 'NULL') || ', '
    || COALESCE(quote_literal(cargo), 'NULL') || ', '
    || COALESCE(quote_literal(numerotelefone), 'NULL') || ', '
    || COALESCE(quote_literal(email), 'NULL') || ');'
AS insert_stmt
FROM reports.contatos_ch;


select count(*) from reports.contatos_ch;