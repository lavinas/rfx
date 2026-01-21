SELECT
    'INSERT INTO cadoc_6334_infrterm (Ano, Trimestre, Uf, QuantidadeTotal, QuantidadePOSCompartilhados, QuantidadePOSLeitoraChip, QuantidadePDV) VALUES ('
    || ano::text || ', '
    || trimestre::text || ', '
    || quote_literal(COALESCE(uf, '')) || ', '
    || (COALESCE(quantidade_pos_compartilhados,0) + COALESCE(quantidade_pos_leitora_chip,0) + COALESCE(quantidade_pdv,0))::text || ', '
    || COALESCE(quantidade_pos_compartilhados,0)::text || ', '
    || COALESCE(quantidade_pos_leitora_chip,0)::text || ', '
    || COALESCE(quantidade_pdv,0)::text
    || ');'
AS insert_stmt
FROM reports.infrterm_ch;


select count(*) from reports.infrterm_ch;