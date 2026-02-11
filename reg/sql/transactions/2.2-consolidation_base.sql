-- só pode débito e crédito
CREATE TABLE cadoc_6334.desconto (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	-- keys fields
    year numeric(4) NOT NULL, -- ano   -- transactional: year(transaction.period_date)
	quarter numeric(1) NOT NULL, -- trimestre -- transactional: quarter(transaction.period_date)
	function varchar(1) NOT NULL, -- função: 'C' - credito, 'D' - débito -- transactional: transaction.transaction_product (conversão DB - 'D', CR - 'C')
	brand numeric(2) NOT NULL, -- 1 - Visa, 2 - Mastercard, 8 - elo -- transactional: transaction.transaction_brand (conversão V - 1, M - 2, E - 8)
	forma_captura numeric(1) NOT NULL, -- 1 - Cartão tarja, 2 - Cartão chip, 5 - contactless -- transactional: transaction.transaction_capture (conversão TJ - 1, CH - 2, CT - 5)
	numero_parcelas numeric(2) NOT NULL, -- 1 a 12 -- transactional: transaction.transaction_installments
	segment_code numeric(3) NOT NULL, -- tabela código de segmento -- transactional: transaction.establishment_mcc (join com a tabela segment_mcc)
    -- values fields
	transaction_amount numeric(15, 2) NULL,
	transaction_quantity numeric(12) NULL,
    -- avg fields
    avg_mcc_fee numeric(4, 2) NULL, -- se esta criando é: round(transaction.revenue_mdr_value / transaction.transaction_amount) * 100, 2) -- se esta atualizando é round((round(transaction.revenue_mdr_value / transaction.transaction_amount) * 100, 2) + avg_mcc_fee * transaction_count) / (transaction_count + 1), 2)
	min_mcc_fee numeric(4, 2) NULL, -- se esta criando é: round(transaction.revenue_mdr_value / transaction.transaction_amount) * 100, 2) -- se esta atualizando é o min (min_mcc_fee, round(transaction.revenue_mdr_value / transaction.transaction_amount) * 100, 2))
	max_mcc_fee numeric(4, 2) NULL, -- se esta criando é: round(transaction.revenue_mdr_value / transaction.transaction_amount) * 100, 2) -- se esta atualizando é o max (max_mcc_fee, round(transaction.revenue_mdr_value / transaction.transaction_amount) * 100, 2))
	stdev_mcc_fee numeric(4, 2) NULL, -- colocar zero por enquanto
	CONSTRAINT cadoc_6334_desconto_pkey PRIMARY KEY (id),
    CONSTRAINT unique_cadoc_6334_desconto UNIQUE (year, quarter, function, brand, forma_captura, numero_parcelas, codigo_segmento)
);

-- tabela de segmentos para mapear o código MCC para segmento
CREATE TABLE cadoc_6334.segment_mcc (
    id bigserial NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    mcc_code numeric(4) NOT NULL, -- código MCC
    segment_code numeric(3) NOT NULL, -- código do segmento
    segment_name varchar(100) NOT NULL, -- nome do segmento
    CONSTRAINT cadoc_6334_segment_mcc_pkey PRIMARY KEY (id),
    CONSTRAINT unique_cadoc_6334_segment_mcc UNIQUE (mcc_code)
);

 create table transaction (
    -- control fields
    id bigserial primary key,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    -- keys fields 
    key1 varchar(50) not null,
    key2 varchar(50) null,
    key3 varchar(50) null,
    -- establishment fields
    establishment_code numeric(50) null,
    establishment_mcc numeric(4) null,
    establishment_terminal_code numeric(50) null,
    -- authorization fields
    bin numeric(11) not null,
    authorization_code varchar(20) null,
    transaction_nsu varchar(20) not null, 
    -- transaction fields
    transaction_date date not null,
    transaction_amount numeric(15,2) not null,
    transaction_installments numeric(2) not null, -- 1 to 12
    transaction_installments_type varchar(10) not null, -- 'loja', 'emissor'
    transaction_brand varchar(2) not null, -- 'V' Visa, 'M' Mastercard, 'E' Elo
    transaction_product varchar(2) not null, -- 'DB' debit, 'CR' credit, 'PX' Pix, 'PP' Pre-pago
    transaction_capture varchar(3) not null, -- 'TJ' - Tarja, 'CH' - Chip, 'CT' - Contactless, 'NP' - Não presencial, 'RC' - Recorrente
    -- financial values
    revenue_mdr_value numeric(15,2) null,
    cost_interchange_value numeric(15,2) null,
    -- references/control
    high_source_priority int not null, -- 30 - webservice, 20 - intercambio, 10 - gestao (o de maior prioridade)
    -- status fields
    status_id int not null, -- 1 - pendente, 2 - pronto
    status_name varchar(20) not null, -- 1 - pendente, 2 - pronto
    status_count int not null default 3, -- contador de tentativas de processamento
    -- closing fields
    period_date date null, -- inicialmente nulo
    period_closing_id bigint null, -- inicialmente nulo, referencia transaction_period_closing(id)
    foreign key (period_closing_id) references transaction_period_closing(id)
);
