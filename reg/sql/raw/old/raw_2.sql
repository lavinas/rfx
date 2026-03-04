CREATE TABLE raw_data.management_transaction (
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
	dt_inserter timestamptz DEFAULT now() NOT NULL,
	last_event_sent_at timestamptz NULL,
	status int4 DEFAULT 2 NULL,
	status_description varchar NULL,
	status_name varchar NULL,
	transac_id varchar NULL,
	CONSTRAINT management_transaction_pkey PRIMARY KEY (cd_transacao_fin)
);


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
	transaction_nsu varchar NULL,
	authorization_code varchar NULL,
	CONSTRAINT intercambio_transaction_pkey PRIMARY KEY (cd_transacao_fin)
);