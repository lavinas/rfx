-- intercambio_transaction
CREATE TABLE raw_data.intercambio_transaction (
    key1 varchar,
    cd_transacao_fin varchar PRIMARY KEY,
    forma_captura varchar,
    dt_processamento timestamptz,
    valor_transacoes numeric,
    percentual_desconto numeric,
    taxa_intercambio_valor numeric,
    bandeira varchar,
    parcela varchar,
    tipo_cartao varchar,
    segmento varchar,
    bin varchar,
    modalidade_cartao varchar,
    produto_cartao varchar,
    cadoc_item_id bigint,
    dt_inserter timestamptz DEFAULT now() NOT NULL,
	last_event_sent_at timestamptz NULL
);
 
-- webservice_transaction
CREATE TABLE raw_data.webservice_transaction (
    ref_num_bnd varchar,
    key1 varchar,
    ref_num_fis varchar PRIMARY KEY,
    transaction_brand varchar,
    transaction_product varchar,
    transaction_date timestamp,
    dt_pos timestamp,
    establishment_terminal_code varchar,
    term_id varchar,
    transaction_amount numeric,
    qtd_parc integer,
    bin varchar,
    dt_inserter timestamptz DEFAULT now() NOT NULL,
	last_event_sent_at timestamptz NULL
);
 
-- management_transaction
CREATE TABLE raw_data.management_transaction (
    key1 varchar,
    cd_transacao_fin varchar PRIMARY KEY,
    cnt_privado varchar,
    dt_processamento timestamp,
    valor_transacao numeric,
    bandeira varchar,
    cd_pessoa_estabelecimento bigint,
    mcc varchar,
    segmento varchar,
    forma_captura varchar,
    funcao varchar,
    numero_parcelas integer,
    desconto_valor numeric,
    percentual_desconto numeric,
    cadoc_item_id bigint,
    dt_inserter timestamptz DEFAULT now() NOT NULL,
	last_event_sent_at timestamptz NULL
);

-- pix_transaction
CREATE TABLE raw_data.pix_transaction (
	loja varchar(255) NULL,
	terminal varchar(255) NULL,
	nsu varchar(255) NULL,
	valoroperacao float8 NULL,
	tipotecn int4 NULL,
	formapagamento int4 NULL,
	c int4 NULL,
	nuparcela int4 NULL,
	tipopessoa varchar(50) NULL,
	transacaosplit varchar(255) NULL,
	datatransacao varchar(50) NULL,
	horatransacao varchar(50) NULL,
	pan varchar(255) NULL,
	mcc varchar(50) NULL,
	cpfcnpjpj varchar(50) NULL,
	cep varchar(20) NULL,
	uf varchar(5) NULL,
	localidade varchar(255) NULL,
	nmlogradouro varchar(255) NULL,
	nulogradouro varchar(50) NULL,
	nmcomplemento varchar(255) NULL,
	nmbairro varchar(255) NULL,
	nmpessoa varchar(255) NULL,
	fonecont varchar(50) NULL,
	emailcont varchar(255) NULL,
	datacredenciamento varchar(50) NULL,
	nomefantasia varchar(255) NULL,
	codibge varchar(20) NULL,
	psp varchar(255) NULL,
	dataoperacao varchar(50) NULL,
	tipoterminal varchar(255) NULL
);