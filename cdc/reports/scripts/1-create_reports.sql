-- Active: 1767998905059@@127.0.0.1@5435@cdc
create schema reports;

-- reports files

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

-- estabelecimentos
CREATE TABLE IF NOT EXISTS reports.conccred (
    id BIGINT PRIMARY KEY,
    sync_status SMALLINT NOT NULL,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    ano SMALLINT NOT NULL,
    trimestre SMALLINT NOT NULL,
    bandeira SMALLINT NOT NULL,
    funcao CHAR(1) NOT NULL,
    quantidade_estabelecimentos_credenciados INTEGER NOT NULL,
    quantidade_estabelecimentos_ativos INTEGER NOT NULL,
    valor_transacoes NUMERIC(18,2) NOT NULL,
    quantidade_transacoes INTEGER NOT NULL
);

-- infraestrutura
CREATE TABLE IF NOT EXISTS reports.infresta (
    id BIGINT PRIMARY KEY,
    sync_status SMALLINT NOT NULL,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    ano SMALLINT NOT NULL,
    trimestre SMALLINT NOT NULL,
    uf CHAR(2) NOT NULL,
    quantidade_estabelecimentos_totais INTEGER NOT NULL,
    quantidade_estabelecimentos_captura_manual INTEGER NOT NULL,
    quantidade_estabelecimentos_captura_eletronica INTEGER NOT NULL,
    quantidade_estabelecimentos_captura_remota INTEGER NOT NULL
);


CREATE TABLE IF NOT EXISTS reports.infrterm (
    id BIGINT PRIMARY KEY,
    sync_status SMALLINT NOT NULL,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    ano SMALLINT NOT NULL,
    trimestre SMALLINT NOT NULL,
    uf CHAR(2) NOT NULL,
    quantidade_pos_compartilhados INTEGER NOT NULL,
    quantidade_pos_leitora_chip INTEGER NOT NULL,
    quantidade_pdv INTEGER NOT NULL
);

-- support tables for reports generation
CREATE TABLE IF NOT EXISTS apoio.estabelecimentos (
    codigo_estabelecimento BIGINT PRIMARY KEY,
    data_credenciamento TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    data_ultima_transacao TIMESTAMP(3) WITHOUT TIME ZONE,
    uf CHAR(2) NOT NULL,
    tem_debito BOOLEAN NOT NULL,
    tem_credito BOOLEAN NOT NULL,
    tem_visa BOOLEAN NOT NULL,
    tem_mastercard BOOLEAN NOT NULL,
    tem_elo BOOLEAN NOT NULL,
    mcc VARCHAR(4) NOT NULL,
    segmento VARCHAR(255),
    captura_manual BOOLEAN NOT NULL,
    captura_eletronica BOOLEAN NOT NULL,
    captura_remota BOOLEAN NOT NULL
);

drop table if exists apoio.gestao;
CREATE TABLE IF NOT EXISTS apoio.gestao (
    bandeira SMALLINT NOT NULL,
    codigo_estabelecimento BIGINT NOT NULL,
    mcc VARCHAR(4) NOT NULL,
    segmento VARCHAR(255),
    forma_captura SMALLINT NOT NULL,
    funcao CHAR(1) NOT NULL,
    numero_parcelas SMALLINT NOT NULL,
    codigo_segmento INTEGER,
    sum_valor_transacoes NUMERIC(18,2) NOT NULL,
    quantidade_transacoes INTEGER NOT NULL,
    sum_percentual_desconto NUMERIC(18,2) NOT NULL,
    avg_percentual_desconto NUMERIC(18,2) NOT NULL,
    min_percentual_desconto NUMERIC(18,2) NOT NULL,
    max_percentual_desconto NUMERIC(18,2) NOT NULL,
    dev_percentual_desconto NUMERIC(18,2),
    sum_taxa_desconto_total NUMERIC(18,2) NOT NULL,
    avg_taxa_desconto_total NUMERIC(18,2) NOT NULL,
    min_taxa_desconto_total NUMERIC(18,2) NOT NULL,
    max_taxa_desconto_total NUMERIC(18,2) NOT NULL,
    dev_taxa_desconto_total NUMERIC(18,2)
);

drop table if exists apoio.intercambio;
CREATE TABLE IF NOT EXISTS apoio.intercambio (
    modalidade_cartao VARCHAR(16),
    produto_cartao VARCHAR(16),
    bandeira VARCHAR(32),
    parcela SMALLINT,
    tipo_cartao VARCHAR(16),
    forma_captura CHAR(2),
    segmento VARCHAR(16),
    sum_valor_transacoes NUMERIC(18,2),
    quantidade_transacoes INTEGER,
    sum_percentual_desconto NUMERIC(18,2),
    avg_percentual_desconto NUMERIC(18,2),
    min_percentual_desconto NUMERIC(18,2),
    max_percentual_desconto NUMERIC(18,2),
    dev_percentual_desconto NUMERIC(18,2),
    sum_taxa_intercambio_valor NUMERIC(18,2),
    avg_taxa_intercambio_valor NUMERIC(18,2),
    min_taxa_intercambio_valor NUMERIC(18,2),
    max_taxa_intercambio_valor NUMERIC(18,2),
    dev_taxa_intercambio_valor NUMERIC(18,2)
);

CREATE TABLE IF NOT EXISTS apoio.terminais (
    codigo_terminal BIGINT NOT NULL,
    codigo_estabelecimento BIGINT NOT NULL,
    tipo_terminal VARCHAR(16) NOT NULL
);

CREATE TABLE IF NOT EXISTS apoio.segmentos (
    mcc_init INTEGER NOT NULL,
    mcc_end INTEGER NOT NULL,
    segment INTEGER NOT NULL
);
create index segmentos_mcc_idx on apoio.segmentos(mcc_init, mcc_end);

