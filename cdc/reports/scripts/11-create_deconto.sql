-- create descontos_up table
CREATE TABLE IF NOT EXISTS reports.descontos_up (
    id BIGINT NULL,
    sync_status SMALLINT NULL,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NULL,
    ano SMALLINT NULL,
    trimestre SMALLINT NULL,
    funcao CHAR(1) NOT NULL,
    bandeira SMALLINT NOT NULL,
    forma_captura SMALLINT NOT NULL,
    numero_parcelas SMALLINT NOT NULL,
    codigo_segmento INTEGER NOT NULL,
    taxa_desconto_media NUMERIC(5,2) NULL,
    taxa_desconto_minima NUMERIC(5,2) NULL,
    taxa_desconto_maxima NUMERIC(5,2) NULL,
    desvio_padrao_taxa_desconto NUMERIC(6,3) NULL,
    valor_transacoes NUMERIC(18,2) NOT NULL,
    quantidade_transacoes INTEGER NOT NULL
);
-- populate descontos_up table from intercam_up
truncate TABLE reports.descontos_up;
insert into reports.descontos_up (
    funcao,
    bandeira,
    forma_captura,
    numero_parcelas,
    codigo_segmento,
    valor_transacoes,
    quantidade_transacoes
)
select funcao,
       bandeira,
       forma_captura,
       numero_parcelas,
       codigo_segmento,
       sum(valor_transacoes) as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
    from reports.intercam_up
group by 1,2,3,4,5;

-- validate descontos_up duplicated
select funcao,
       bandeira,
       forma_captura,
       numero_parcelas,
       codigo_segmento,
       count(1)
  from reports.descontos_up
    group by 1,2,3,4,5
having count(1) > 1;
-- validate descontos_duplicated
select funcao,
       bandeira,
       forma_captura,
       numero_parcelas,
       codigo_segmento,
       count(1)
  from reports.descontos
    group by 1,2,3,4,5
having count(1) > 1;

-- count missing
select count(1)
  from reports.descontos_up a
left join reports.descontos b
    on a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento
 where b.id is null; 

-- count missing
select count(1)
  from reports.descontos a
left join reports.descontos_up b
    on a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento
 where b.funcao is null ; 

