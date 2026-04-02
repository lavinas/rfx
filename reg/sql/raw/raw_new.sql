-- Active: 1766518799113@@127.0.0.1@5434@reg@public

CREATE SCHEMA IF NOT EXISTS raw_data;

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
    extractor_execution_id int8 NOT NULL,
    CONSTRAINT intercambio_transaction_pkey PRIMARY KEY (cd_transacao_fin),
    CONSTRAINT fk_extractor_execution FOREIGN KEY (extractor_execution_id) REFERENCES public.extractor_execution(id)
);
CREATE INDEX idx_dt_processamento ON raw_data.intercambio_transaction (dt_processamento);
CREATE INDEX idx_extractor_execution_id ON raw_data.intercambio_transaction (extractor_execution_id);


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
    extractor_execution_id int8 NOT NULL,
	CONSTRAINT management_transaction_pkey PRIMARY KEY (cd_transacao_fin),
    CONSTRAINT fk_extractor_execution FOREIGN KEY (extractor_execution_id) REFERENCES public.extractor_execution(id)
);
CREATE INDEX idx_management_dt_processamento ON raw_data.management_transaction (dt_processamento);
CREATE INDEX idx_management_extractor_execution_id ON raw_data.management_transaction (extractor_execution_id);

 
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
    extractor_execution_id int8 NOT NULL,
    -- reconciliation control
	CONSTRAINT webservice_transaction_pkey PRIMARY KEY (ref_num_fis),
    CONSTRAINT fk_extractor_execution FOREIGN KEY (extractor_execution_id) REFERENCES public.extractor_execution(id)
);
CREATE INDEX idx_webservice_transaction_date ON raw_data.webservice_transaction (transaction_date);
CREATE INDEX idx_webservice_extractor_execution_id ON raw_data.webservice_transaction (extractor_execution_id);


CREATE TABLE raw_data.establishment (
	establishment_code int8 NOT NULL,
	accreditation_date timestamp NULL,
    company_name varchar(255) NULL,
	trading_name varchar(255) NULL,
	cnpj varchar(255) NULL,
	cpf varchar(255) NULL,
	mcc_code varchar(255) NULL,
	address varchar(255) NULL,
	cep varchar(20) NULL,
	city_ibge_code int8,
	federation_unit varchar(2) NULL,
	contact_name varchar(255) NULL,
	contact_phone int4 NULL,
	contact_email varchar(255) NULL,
	has_debit bool NULL,
	has_credit bool NULL,
	has_visa bool NULL,
	has_mastercard bool NULL,
	has_elo bool NULL,
	has_manual_capture bool NULL,
	has_eletronic_capture bool NULL,
	has_remote_capture bool NULL,
	is_subacquirer bool NULL,
    -- inserter control
	dt_inserter timestamptz DEFAULT now() NOT NULL,
    extractor_execution_id int8 NOT NULL,
    -- transactional control
	CONSTRAINT establishment_pkey PRIMARY KEY (establishment_code),
    CONSTRAINT fk_extractor_execution FOREIGN KEY (extractor_execution_id) REFERENCES public.extractor_execution(id)
);
CREATE INDEX idx_establishment_extractor_execution_id ON raw_data.establishment (extractor_execution_id);

create table raw_data.terminals (
	terminal_code varchar(10) NOT NULL,
	establishment_code int8 NOT NULL,
	terminal_type varchar(3), -- 'POS' ou 'TEF'
   -- inserter control
	dt_inserter timestamptz DEFAULT now() NOT NULL,
    extractor_execution_id int8 NOT NULL,
    -- transactional control
	CONSTRAINT terminals_transaction_pkey PRIMARY KEY (terminal_code),
    CONSTRAINT fk_extractor_execution FOREIGN KEY (extractor_execution_id) REFERENCES public.extractor_execution(id)
);
CREATE INDEX idx_terminals_extractor_execution_id ON raw_data.terminals (extractor_execution_id);