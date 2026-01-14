create schema reports;

-- descontos
CREATE TABLE IF NOT EXISTS reports.descontos (
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

-- intercam
CREATE TABLE IF NOT EXISTS reports.intercam (
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


CREATE TABLE IF NOT EXISTS reports.ranking (
    id BIGINT PRIMARY KEY,
    sync_status SMALLINT NOT NULL,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    ano SMALLINT NOT NULL,
    trimestre SMALLINT NOT NULL,
    codigo_estabelecimento VARCHAR(20) NOT NULL,
    funcao CHAR(1) NOT NULL,
    bandeira SMALLINT NOT NULL,
    forma_captura SMALLINT NOT NULL,
    numero_parcelas SMALLINT NOT NULL,
    codigo_segmento INTEGER NOT NULL,
    valor_transacoes NUMERIC(18,2) NOT NULL,
    quantidade_transacoes INTEGER NOT NULL,
    taxa_desconto_media NUMERIC(5,2) NOT NULL
);