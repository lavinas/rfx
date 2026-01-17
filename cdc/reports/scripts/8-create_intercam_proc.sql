CREATE OR REPLACE PROCEDURE reports.create_intercam(
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
CREATE TABLE IF NOT EXISTS reports.intercam_up (
    id BIGINT PRIMARY KEY,
    sync_status SMALLINT NOT NULL,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    ano SMALLINT NOT NULL,
    trimestre SMALLINT NOT NULL,
    produto SMALLINT NOT NULL,
    modalidade_cartao CHAR(1) NOT NULL,
    funcao CHAR(1) NOT NULL,
    bandeira SMALLINT NOT NULL,
    forma_captura SMALLINT NOT NULL,
    numero_parcelas SMALLINT NOT NULL,
    codigo_segmento INTEGER NOT NULL,
    tarifa_intercambio NUMERIC(7,4) NOT NULL,
    valor_transacoes NUMERIC(18,2) NOT NULL,
    quantidade_transacoes INTEGER NOT NULL
);
    -- delete bandeira records
    delete from reports.intercam_up
     where bandeira = target_bandeira;

    -- get proportions
    select round(target_tpv / sum(valor_transacoes), 15),
           round(cast(target_qtd as NUMERIC) / cast(sum(quantidade_transacoes) as numeric), 15)
      into prop_valor_transacoes, 
           prop_quantidade_transacoes
      from reports.intercam
     where bandeira = target_bandeira;
    -- insert adjusted records
    insert into reports.intercam_up
    select a.id, 
           a.sync_status, 
           a.created_at, 
           a.updated_at, 
           a.ano, 
           a.trimestre,
           a.produto, 
           a.modalidade_cartao, 
           a.funcao, 
           a.bandeira, 
           a.forma_captura,
           a.numero_parcelas, 
           a.codigo_segmento, 
           a.tarifa_intercambio,
           round(a.valor_transacoes * prop_valor_transacoes, 2) as valor_transacoes,
           round(a.quantidade_transacoes * prop_quantidade_transacoes, 0) as quantidade_transacoes
        from reports.intercam a
         where a.bandeira = target_bandeira;
    -- calculate remaining difference
    select target_qtd - sum(a.quantidade_transacoes), 
           target_tpv - sum(a.valor_transacoes)
      into resto_quantidade_transacoes, resto_valor_transacoes
      from reports.intercam_up a
     where a.bandeira = target_bandeira;
    -- select the biggest record to adjust
    update reports.intercam_up
       set quantidade_transacoes = quantidade_transacoes + resto_quantidade_transacoes,
           valor_transacoes = valor_transacoes + resto_valor_transacoes
     where id = (select id
                   from reports.intercam_up
                  where bandeira = target_bandeira
                  order by quantidade_transacoes desc
                  limit 1);
END;
$$ LANGUAGE plpgsql;