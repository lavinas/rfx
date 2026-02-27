CREATE TABLE raw_data.intercambio_transaction (
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
    -- inserter control
	dt_inserter timestamptz DEFAULT now() NOT NULL,
    -- transactional control old
	-- last_event_sent_at timestamptz NULL,
	-- status int4 DEFAULT 2 NULL,
	-- status_description varchar NULL,
	-- status_name varchar NULL,
    -- transaction control
    transactional_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    transactional_status_date timestamp NULL,
    -- reconciliation control
    reconciliation_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    reconciliation_status_date timestamp NULL,
    CONSTRAINT intercambio_transaction_pkey PRIMARY KEY (cd_transacao_fin)
);

CREATE TABLE raw_data.management_transaction (
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
    -- inserter control
	dt_inserter timestamptz DEFAULT now() NOT NULL,
	-- transactional control old
	-- last_event_sent_at timestamptz NULL,
	-- status int4 DEFAULT 2 NULL,
	-- status_description varchar NULL,
	-- status_name varchar NULL,
    -- transaction control
    transactional_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    transactional_status_date timestamp NULL,
    -- reconciliation control
    reconciliation_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    reconciliation_status_date timestamp NULL,
	CONSTRAINT management_transaction_pkey PRIMARY KEY (cd_transacao_fin)
);

 
CREATE TABLE raw_data.webservice_transaction (
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
    -- inserter control
	dt_inserter timestamptz DEFAULT now() NOT NULL,
    -- transactional control old
	-- last_event_sent_at timestamptz NULL,
	-- status int4 DEFAULT 2 NULL,
	-- status_description varchar NULL,
	-- status_name varchar NULL,
    -- transaction control
    transactional_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    transactional_status_date timestamp NULL,
    -- reconciliation control
    reconciliation_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    reconciliation_status_date timestamp NULL
);


CREATE TABLE raw_data.tc57_transaction (
    -- data
    key1 varchar NULL,
	transaction_brand varchar NULL,  -- visa, master, ello
	transaction_product varchar NULL, -- debito, credito
	transaction_date timestamp NULL, -- data/hora da transacao 
    qtd_parc int4 NULL, -- parcelas
    transaction_amount numeric NULL, -- valor da transacao
    -- transaction control
    transactional_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    transactional_status_date timestamp NULL,
    -- reconciliation control
    reconciliation_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    reconciliation_status_date timestamp NULL,
    -- constraints
    constraint uk_base_id unique (base_id),
    constraint fk_base_id foreign key (base_id) references raw_data.base(id)
);