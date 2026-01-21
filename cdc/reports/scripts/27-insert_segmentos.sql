SELECT
    'INSERT INTO cadoc_6334_segmento (NomeSegmento, DescricaoSegmento, CodigoSegmento) VALUES (' ||
    quote_literal(left(nome_segmento,50)) || ', ' ||
    quote_literal(left(descricao_segmento,255)) || ', ' ||
    COALESCE(codigo_segmento::text, 'NULL') || ');'
AS insert_stmt
FROM reports.segmentos_ch
limit 100000;


select count(1)
from reports.segmentos_ch;