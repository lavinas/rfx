-- Active: 1766518799113@@127.0.0.1@5434@reg

-- public.gestao_transacao definição
 
-- Drop table
 
-- DROP TABLE public.gestao_transacao;
 
CREATE TABLE public.gestao_transacao (
	id bigserial NOT NULL,
	cd_transacao_fin varchar(200) NOT NULL,
	cnt_privado varchar(50) NULL,
	dt_processamento date NOT NULL,
	valor_transacao numeric(18, 2) NOT NULL,
	bandeira varchar(50) NULL,
	codigo_estabelecimento varchar(50) NULL,
	mcc varchar(10) NULL,
	segmento varchar(50) NULL,
	forma_captura varchar(50) NULL,
	funcao int4 NULL,
	numero_parcelas int4 NULL,
	desconto_valor numeric(18, 2) NULL,
	percentual_desconto numeric(18, 6) NULL,
	cadoc_item_id varchar(50) NULL,
	status int4 DEFAULT 1 NOT NULL,
	last_event_sent_at timestamp NULL,
	dt_inserter timestamp DEFAULT now() NOT NULL,
	status_name varchar(50) NULL,
	status_description text NULL,
	CONSTRAINT gestao_transacao_pkey PRIMARY KEY (id),
	CONSTRAINT gestao_transacao_uk UNIQUE (cd_transacao_fin, cnt_privado)
);
CREATE INDEX idx_gestao_transacao_status_last_event ON public.gestao_transacao USING btree (status, last_event_sent_at);
 
 
-- public.intercambio_transacao definição
 
-- Drop table
 
-- DROP TABLE public.intercambio_transacao;
 
CREATE TABLE public.intercambio_transacao (
	id bigserial NOT NULL,
	cd_transacao_fin varchar(200) NULL,
	forma_captura varchar(50) NULL,
	dt_processamento date NULL,
	valor_transacoes numeric(18, 2) NULL,
	percentual_desconto numeric(18, 6) NULL,
	taxa_intercambio_valor numeric(18, 6) NULL,
	bandeira varchar(50) NULL,
	parcela varchar(20) NULL,
	tipo_cartao varchar(50) NULL,
	segmento varchar(10) NULL,
	bin varchar(50) NULL,
	status int4 DEFAULT 1 NOT NULL,
	last_event_sent_at timestamp NULL,
	dt_inserter timestamp NULL,
	status_name varchar(50) NULL,
	status_description text NULL,
	CONSTRAINT intercambio_transacao_pkey PRIMARY KEY (id)
);
 
 
-- public.webservice_transacao definição
 
-- Drop table
 
-- DROP TABLE public.webservice_transacao;
 
CREATE TABLE public.webservice_transacao (
	id bigserial NOT NULL,
	ref_num_fis varchar(50) NOT NULL,
	ref_num_bnd varchar(100) NOT NULL,
	key1 varchar(150) NOT NULL,
	transaction_brand varchar(100) NULL,
	transaction_product varchar(100) NULL,
	transaction_date timestamp NOT NULL,
	dt_pos timestamp NULL,
	establishment_terminal_code varchar(50) NULL,
	term_id varchar(50) NULL,
	transaction_amount numeric(18, 2) NOT NULL,
	transaction_installments int4 NULL,
	bin varchar(30) NULL,
	status int4 DEFAULT 1 NOT NULL,
	last_event_sent_at timestamp NULL,
	dt_inserter timestamp DEFAULT now() NOT NULL,
	status_name varchar(50) NULL,
	status_description text NULL,
	CONSTRAINT webservice_transacao_pkey PRIMARY KEY (id),
	CONSTRAINT webservice_transacao_uk UNIQUE (ref_num_fis, ref_num_bnd)
);