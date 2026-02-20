CREATE TABLE raw_data.intercambio_transaction (
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
	dt_inserter timestamptz DEFAULT now() NOT NULL,
	last_event_sent_at timestamptz NULL,
	status int4 DEFAULT 2 NULL,
	status_description varchar NULL,
	status_name varchar NULL,
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
	dt_inserter timestamptz DEFAULT now() NOT NULL,
	last_event_sent_at timestamptz NULL,
	status int4 DEFAULT 2 NULL,
	status_description varchar NULL,
	status_name varchar NULL,
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
	dt_inserter timestamptz DEFAULT now() NOT NULL,
	last_event_sent_at timestamptz NULL,
	status int4 DEFAULT 2 NULL,
	status_description varchar NULL,
	status_name varchar NULL,
	CONSTRAINT webservice_transaction_pkey PRIMARY KEY (ref_num_fis)
);