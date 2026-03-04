CREATE TABLE raw_data.intercambio_transaction (
    -- data
	cd_transacao_fin varchar NOT NULL,
	key1 varchar NULL,  -- chave transformada para poder cruzar entre sistemas
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
	transaction_nsu varchar NULL,
	authorization_code varchar NULL,
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
CREATE INDEX idx_intercambio_transaction_transactional_status_id ON raw_data.intercambio_transaction (transactional_status_id);
CREATE INDEX idx_intercambio_transaction_reconciliation_status_id ON raw_data.intercambio_transaction (reconciliation_status_id);
CREATE INDEX idx_dt_processamento ON raw_data.intercambio_transaction (dt_processamento);


CREATE TABLE raw_data.management_transaction (
	-- data
	cd_transacao_fin varchar NOT NULL,
	key1 varchar NULL,  -- chave transformada para poder cruzar entre sistemas
	-- cnt_privado varchar NULL,
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
	transac_id varchar NULL,
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
CREATE INDEX idx_management_transaction_transactional_status_id ON raw_data.management_transaction (transactional_status_id);
CREATE INDEX idx_management_transaction_reconciliation_status_id ON raw_data.management_transaction (reconciliation_status_id);
CREATE INDEX idx_management_dt_processamento ON raw_data.management_transaction (dt_processamento);

 
CREATE TABLE raw_data.webservice_transaction (
	-- data
	ref_num_bnd varchar NULL,
	key1 varchar NULL,  -- chave transformada para poder cruzar entre sistemas
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
	CONSTRAINT webservice_transaction_pkey PRIMARY KEY (ref_num_fis)
);
CREATE INDEX idx_webservice_transaction_transactional_status_id ON raw_data.webservice_transaction (transactional_status_id);
CREATE INDEX idx_webservice_transaction_reconciliation_status_id ON raw_data.webservice_transaction (reconciliation_status_id);
CREATE INDEX idx_webservice_transaction_date ON raw_data.webservice_transaction (transaction_date);


CREATE TABLE raw_data.tc57_transaction (
    -- data
	cd_transacao_fin varchar NOT NULL,
    key1 varchar NULL,  -- amanha alinhamos
	transaction_brand varchar NULL,  -- TC57-TCR0-TR2.campo20 ('E' - Elo, 'V' - Visa, 'M' - Master)    
	transaction_product varchar NULL, -- TC57-TCR0-TR2.campo33 ('D' - Debito, '' - Credito)
	transaction_date timestamp NULL, -- TC57-TCR0_TR2.campo10 TC57-TCR0_TR2.campo22 + TC57-TCR0-TR2.campo23
    qtd_parc int4 NULL, -- Elo: TC57-TCR1-EL.campo31; Visa: TC57-TCRD-IP.campo07 (se não existir = 1); Master: TC57-TCR5-TR2.campo47 (se '' = 1)
    transaction_amount numeric NULL, -- TC57-TCR0-TR2.campo11
	establishment_code varchar NULL, --TC57-TCR0-TR1.campo20 (armazenar na variavel cada vez que coloca ele)
	bin varchar NULL, -- TC57-TCR0-TR2.campo21 (primeiros 8 dígitos)
    -- inserter control
	dt_inserter timestamptz DEFAULT now() NOT NULL,
    -- transaction control
    transactional_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    transactional_status_date timestamp NULL,
    -- reconciliation control
    reconciliation_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    reconciliation_status_date timestamp NULL,
    -- constraints
	CONSTRAINT tc57_transaction_pkey PRIMARY KEY (cd_transacao_fin)
);
CREATE INDEX idx_tc57_transaction_transactional_status_id ON raw_data.tc57_transaction (transactional_status_id);
CREATE INDEX idx_tc57_transaction_reconciliation_status_id ON raw_data.tc57_transaction (reconciliation_status_id);
CREATE INDEX idx_tc57_transaction_date ON raw_data.tc57_transaction (transaction_date);


CREATE TABLE raw_data.pix_transaction (
    -- data
	key_pix varchar NULL,
	transaction_date timestamp NULL, -- data/hora da transacao 
    transaction_amount numeric NULL, -- valor da transacao
	transaction_nsu varchar NULL,
	authorization_code varchar NULL,
	transac_id varchar NULL,
	originator_bank_qr_code varchar NULL,
	establishment_code varchar NULL,
	-- inserter control
	document_type int4 NULL, -- 0 - cpf, 1 - cnpj 
    -- inserter control
	dt_inserter timestamptz DEFAULT now() NOT NULL,
    -- transaction control
    transactional_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    transactional_status_date timestamp NULL,
    -- reconciliation control
    reconciliation_status_id int4 DEFAULT 0 NOT NULL, -- 0 - pending, 1 - translated, 2 - recognized, 3 - error
    reconciliation_status_date timestamp NULL,
    -- constraints
    CONSTRAINT pix_transaction_pkey PRIMARY KEY (key_pix)
);
CREATE INDEX idx_pix_transaction_transactional_status_id ON raw_data.pix_transaction (transactional_status_id);
CREATE INDEX idx_pix_transaction_reconciliation_status_id ON raw_data.pix_transaction (reconciliation_status_id);
CREATE INDEX idx_pix_transaction_date ON raw_data.pix_transaction (transaction_date);

