
-- base table for raw data, with common fields and indexes for transactional and reconciliation sent flags
CREATE TABLE raw_data.base (
    id bigserial PRIMARY KEY,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    transactional_sent boolean not null default false,
    reconciliation_sent boolean not null default false
);
create index idx_base_transactional_sent on raw_data.base (transactional_sent);
create index idx_base_reconciliation_sent on raw_data.base (reconciliation_sent);

-- raw tables for each source, with foreign key to base table and unique constraint on source-specific transaction
CREATE TABLE raw_data.intercambio_transaction (
    -- id
    id bigserial PRIMARY KEY,
    base_id bigint not null,
    -- data
	key1 varchar NULL,
	cd_transacao_fin varchar NOT NULL,
	forma_captura varchar NULL,
	dt_processamento timestamptz NULL,
	valor_transacoes numeric NULL,
	percentual_desconto numeric NULL,
	taxa_intercambio_valor numeric NULL,
	bandeira varchar NULL,
	parcela varchar NULL,
	tipo_cartao varchar NULL,
	segmento varchar NULL,
	bin varchar NULL,
    -- constraints
    constraint uk_base_id unique (base_id),
    constraint uq_cd_transacao_fin unique (cd_transacao_fin),
    constraint fk_base_id foreign key (base_id) references raw_data.base(id)
);

CREATE TABLE raw_data.management_transaction (
    -- base
    id bigserial PRIMARY KEY,
    base_id bigint not null,
    -- data
	key1 varchar NULL,
	cd_transacao_fin varchar NOT NULL,
	cnt_privado varchar NULL,
	dt_processamento timestamp NULL,
	valor_transacao numeric NULL,
	bandeira varchar NULL,
	cd_pessoa_estabelecimento int8 NULL,
	mcc varchar NULL,
	forma_captura varchar NULL,
	funcao varchar NULL,
	numero_parcelas int4 NULL,
	desconto_valor numeric NULL,
	percentual_desconto numeric NULL,
    -- constraints
    constraint uk_base_id unique (base_id),
    constraint uq_cd_transacao_fin unique (cd_transacao_fin),
    constraint fk_base_id foreign key (base_id) references raw_data.base(id)
);

-- raw table for webservice transactions, with unique constraint on ref_num_fis and foreign key to base table
CREATE TABLE raw_data.webservice_transaction (
    -- base
    id bigserial PRIMARY KEY,
    base_id bigint not null,
    -- data
	ref_num_bnd varchar NULL,
	key1 varchar NULL,
	ref_num_fis varchar NOT NULL,
	transaction_brand varchar NULL,
	transaction_product varchar NULL,
	transaction_date timestamp NULL,
	dt_pos timestamp NULL,
	establishment_terminal_code varchar NULL,
	term_id varchar NULL,
	transaction_amount numeric NULL,
	qtd_parc int4 NULL,
	bin varchar NULL,
    -- constraints
    constraint uk_base_id unique (base_id),
    constraint uq_ref_num_fis unique (ref_num_fis),
    constraint fk_base_id foreign key (base_id) references raw_data.base(id)
);

-- raw table for TC57 transactions, with unique constraint on base_id and foreign key to base table
CREATE TABLE raw_data.tc57_transaction (
    -- base
    id bigserial PRIMARY KEY,
    base_id bigint not null,
    -- data
    key1 varchar NULL,
	transaction_brand varchar NULL,
	transaction_product varchar NULL,
	transaction_date timestamp NULL,
    qtd_parc int4 NULL,
    transaction_amount numeric NULL,
    -- constraints
    constraint uk_base_id unique (base_id),
    constraint fk_base_id foreign key (base_id) references raw_data.base(id)
);