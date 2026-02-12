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
    avg_mcc_fee numeric(4, 2) NULL, 
	min_mcc_fee numeric(4, 2) NULL, -- se esta criando é: round(transaction.revenue_mdr_value / transaction.transaction_amount) * 100, 2) -- se esta atualizando é o min (min_mcc_fee, round(transaction.revenue_mdr_value / transaction.transaction_amount) * 100, 2))
	max_mcc_fee numeric(4, 2) NULL, -- se esta criando é: round(transaction.revenue_mdr_value / transaction.transaction_amount) * 100, 2) -- se esta atualizando é o max (max_mcc_fee, round(transaction.revenue_mdr_value / transaction.transaction_amount) * 100, 2))
	stdev_mcc_fee numeric(4, 2) NULL,
    -- aux fields
    sqrdiff_mcc_fee numeric(4, 2) NULL, 
	CONSTRAINT cadoc_6334_desconto_pkey PRIMARY KEY (id),
    CONSTRAINT unique_cadoc_6334_desconto UNIQUE (year, quarter, function, brand, forma_captura, numero_parcelas, segment_code)
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
