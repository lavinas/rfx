CREATE OR REPLACE PROCEDURE reports.create_desconto(
    target_bandeira INT,
    target_tpv NUMERIC,
    target_qtd INT
) AS $$
DECLARE
-- variables
   prop_valor_transacoes NUMERIC;
   prop_quantidade_transacoes NUMERIC;
   resto_valor_transacoes NUMERIC;
   resto_quantidade_transacoes INT;
BEGIN
    -- create table if not exists
    CREATE TABLE IF NOT EXISTS reports.descontos_ch (
        id BIGINT PRIMARY KEY,
        sync_status SMALLINT NOT NULL,
        created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
        updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
        ano SMALLINT NOT NULL,
        trimestre SMALLINT NOT NULL,
        funcao CHAR(1) NOT NULL,
        bandeira SMALLINT NOT NULL,
        forma_captura SMALLINT NOT NULL,
        numero_parcelas SMALLINT NOT NULL,
        codigo_segmento INTEGER NOT NULL,
        taxa_desconto_media NUMERIC(5,2) NOT NULL,
        taxa_desconto_minima NUMERIC(5,2) NOT NULL,
        taxa_desconto_maxima NUMERIC(5,2) NOT NULL,
        desvio_padrao_taxa_desconto NUMERIC(6,3) NOT NULL,
        valor_transacoes NUMERIC(18,2) NOT NULL,
        quantidade_transacoes INTEGER NOT NULL
    );
    -- delete bandeira records
    delete from reports.descontos_ch
     where bandeira = target_bandeira;

    -- get proportions
    select round(target_tpv / sum(valor_transacoes), 15),
           round(cast(target_qtd as NUMERIC) / cast(sum(quantidade_transacoes) as numeric), 15)
      into prop_valor_transacoes, 
           prop_quantidade_transacoes
      from reports.descontos
     where bandeira = target_bandeira;
    -- insert adjusted records
    insert into reports.descontos_ch
    select a.id, 
           a.sync_status, 
           a.created_at, 
           a.updated_at, 
           a.ano, 
           a.trimestre,
           a.funcao, 
           a.bandeira, 
           a.forma_captura,
           a.numero_parcelas, 
           a.codigo_segmento, 
           a.taxa_desconto_media,
           a.taxa_desconto_minima,
           a.taxa_desconto_maxima,
           a.desvio_padrao_taxa_desconto,
           round(a.valor_transacoes * prop_valor_transacoes, 2) as valor_transacoes,
           round(a.quantidade_transacoes * prop_quantidade_transacoes, 0) as quantidade_transacoes
        from reports.descontos a
         where a.bandeira = target_bandeira;
    -- calculate remaining difference
    select target_qtd - sum(a.quantidade_transacoes), 
           target_tpv - sum(a.valor_transacoes)
      into resto_quantidade_transacoes, resto_valor_transacoes
      from reports.descontos_ch a
     where a.bandeira = target_bandeira;
    -- select the biggest record to adjust
    update reports.descontos_ch
       set quantidade_transacoes = quantidade_transacoes + resto_quantidade_transacoes,
           valor_transacoes = valor_transacoes + resto_valor_transacoes
     where id = (select id
                   from reports.descontos_ch
                  where bandeira = target_bandeira
                  order by quantidade_transacoes desc
                  limit 1);
END;
$$ LANGUAGE plpgsql;