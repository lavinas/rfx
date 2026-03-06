-- só pode débito e crédito
CREATE TABLE cadoc_6334.discount (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	-- keys fields
    year numeric(4) NOT NULL, -- ano   -- transactional: year(transaction.period_date)
	quarter numeric(1) NOT NULL, -- trimestre -- transactional: quarter(transaction.period_date)
	function varchar(1) NOT NULL, -- função: 'C' - credito, 'D' - débito -- transactional: transaction.transaction_product (conversão DB - 'D', CR - 'C')
	brand numeric(2) NOT NULL, -- 1 - Visa, 2 - Mastercard, 8 - elo -- transactional: transaction.transaction_brand (conversão V - 1, M - 2, E - 8)
	capture_mode numeric(1) NOT NULL, -- 1 - Cartão tarja, 2 - Cartão chip, 5 - contactless -- transactional: transaction.transaction_capture (conversão TJ - 1, CH - 2, CT - 5)
	installments numeric(2) NOT NULL, -- 1 a 12 -- transactional: transaction.transaction_installments
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
    CONSTRAINT unique_cadoc_6334_desconto UNIQUE (year, quarter, function, brand, capture_mode, installments, segment_code)
);
CREATE INDEX idx_cadoc_6334_desconto_year_quarter ON cadoc_6334.discount (year, quarter);

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

-- interchange
CREATE TABLE IF NOT EXISTS cadoc_6334.interchange (
    id BIGINT PRIMARY KEY,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    year SMALLINT NOT NULL, -- transactional: year(transaction.period_date)
    quarter SMALLINT NOT NULL, -- transactional: quarter(transaction.period_date)
    product_code SMALLINT NOT NULL, -- cadoc_6334.segment_mcc(transaction.bin).product_code
    card_type CHAR(1) NOT NULL, -- -- cadoc_6334.segment_mcc(transaction.bin).card_type
    function CHAR(1) NOT NULL, -- -- função: 'C' - credito, 'D' - débito -- transactional: transaction.transaction_product (conversão DB - 'D', CR - 'C')
    brand SMALLINT NOT NULL, -- 1 - Visa, 2 - Mastercard, 8 - elo -- transactional: transaction.transaction_brand (conversão V - 1, M - 2, E - 8)
    capture_mode SMALLINT NOT NULL, -- 1 - Cartão tarja, 2 - Cartão chip, 5 - contactless -- transactional: transaction.transaction_capture (conversão TJ - 1, CH - 2, CT - 5)
    installments SMALLINT NOT NULL,  -- 1 a 12 -- transactional: transaction.transaction_installments
    segment_code INTEGER NOT NULL, -- tabela código de segmento -- transactional: transaction.establishment_mcc (join com a tabela segment_mcc)
    interchange_fee NUMERIC(7,4) NOT NULL, -- round(transaction.cost_interchange_value / transaction.transaction_amount * 100, 2) -- usar o algoritmo de media incremental 
    transaction_amount NUMERIC(18,2) NOT NULL, -- +=transaction.transaction_amount)
    transaction_quantity INTEGER NOT NULL, -- += 1
	CONSTRAINT unique_cadoc_6334_interchange UNIQUE (year, quarter, product_code, card_type, function, brand, capture_mode, installments, segment_code)
);
CREATE INDEX idx_cadoc_6334_interchange_year_quarter ON cadoc_6334.interchange (year, quarter);

CREATE TABLE IF NOT EXISTS cadoc_6334.bins (
    id bigserial NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
	bin VARCHAR(8) NOT NULL,
	product_code SMALLINT NOT NULL,
	card_type CHAR(1) NOT NULL,
	CONSTRAINT unique_cadoc_6334_bins_bin UNIQUE (bin)
);

CREATE TABLE IF NOT EXISTS cadoc_6334.ranking_establishments (
    id BIGINT PRIMARY KEY,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    year SMALLINT NOT NULL, -- transactional: year(transaction.period_date)
    quarter SMALLINT NOT NULL, -- transactional: quarter(transaction.period_date)
    establishment_code BIGINT NOT NULL, -- transactional: transaction.establishment_code
    function CHAR(1) NOT NULL, -- -- função: 'C' - credito, 'D' - débito -- transactional: transaction.transaction_product (conversão DB - 'D', CR - 'C')
    brand SMALLINT NOT NULL, -- 1 - Visa, 2 - Mastercard, 8 - elo -- transactional: transaction.transaction_brand (conversão V - 1, M - 2, E - 8)
    capture_mode SMALLINT NOT NULL, -- 1 - Cartão tarja, 2 - Cartão chip, 5 - contactless -- transactional: transaction.transaction_capture (conversão TJ - 1, CH - 2, CT - 5)
    installments SMALLINT NOT NULL,  -- 1 a 12 -- transactional: transaction.transaction_installments
    segment_code SMALLINT NOT NULL, tabela código de segmento -- transactional: transaction.establishment_mcc (join com a tabela segment_mcc)
    transaction_amount NUMERIC(18,2) NOT NULL, -- +=transaction.transaction_amount
    transaction_quantity INTEGER NOT NULL, -- += 1
    avg_mcc_fee numeric(4, 2) NULL, -- o mesmo algortimo de media calculado para desconto, porém utilizando apenas a média
    CONSTRAINT unique_cadoc_6334_ranking_establishments UNIQUE (year, quarter, establishment_code, segment_code)
);
CREATE INDEX idx_cadoc_6334_ranking_establishments_year_quarter ON cadoc_6334.ranking_establishments (year, quarter);


CREATE TABLE IF NOT EXISTS cadoc_6334.ranking (
    id BIGINT PRIMARY KEY,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL,
    year SMALLINT NOT NULL, -- transactional: year(transaction.period_date)
    quarter SMALLINT NOT NULL, -- transactional: quarter(transaction.period_date)
    establishment_code BIGINT NOT NULL, -- transactional: transaction.establishment_code
    function CHAR(1) NOT NULL, -- -- função: 'C' - credito, 'D' - débito -- transactional: transaction.transaction_product (conversão DB - 'D', CR - 'C')
    brand SMALLINT NOT NULL, -- 1 - Visa, 2 - Mastercard, 8 - elo -- transactional: transaction.transaction_brand (conversão V - 1, M - 2, E - 8)
    capture_mode SMALLINT NOT NULL, -- 1 - Cartão tarja, 2 - Cartão chip, 5 - contactless -- transactional: transaction.transaction_capture (conversão TJ - 1, CH - 2, CT - 5)
    installments SMALLINT NOT NULL,  -- 1 a 12 -- transactional: transaction.transaction_installments
    segment_code SMALLINT NOT NULL, tabela código de segmento -- transactional: transaction.establishment_mcc (join com a tabela segment_mcc)
    transaction_amount NUMERIC(18,2) NOT NULL, -- +=transaction.transaction_amount
    transaction_quantity INTEGER NOT NULL, -- += 1
    avg_mcc_fee numeric(4, 2) NULL, -- o mesmo algortimo de media calculado para desconto, porém utilizando apenas a média
    CONSTRAINT unique_cadoc_6334_ranking UNIQUE (year, quarter, establishment_code, segment_code)
);
CREATE INDEX idx_cadoc_6334_ranking_year_quarter ON cadoc_6334.ranking (year, quarter);


    
-- segnments table para mapear o código de segmento (segment_code) para o nome do segmento (segment_name)
-- se o segment_code já existe na tabela, verifica se o mcc está concatenado no description, se não estiver, concatena o mcc no description. Exemplo: 'MCC: 4816, 5045, 5065, 5722, 5732, 5734, 7379, 7622, 7629'
-- se o segment_code não existe na tabela, insere o novo segment_code, segment_name e description com o mcc da transação (transaction.establishment_mcc)
CREATE TABLE IF NOT EXISTS cadoc_6334.segments (
    id bigserial NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    segment_code numeric(3) NOT NULL, -- código do segmento transactional: transaction.establishment_mcc (join com a tabela segment_mcc - campo segment_code)
    segment_name varchar(100) NOT NULL, -- nome do segmento transactional: transaction.establishment_mcc (join com a tabela segment_mcc - campo segment_name)
    description varchar(600) NULL, -- concatenar o mcc mcc no texto (transaction.establishment_mcc). Exemplo: 'MCC: 4816, 5045, 5065, 5722, 5732, 5734, 7379, 7622, 7629'
    CONSTRAINT cadoc_6334_segments_pkey PRIMARY KEY (id),
    CONSTRAINT unique_cadoc_6334_segments UNIQUE (segment_code)
);


CREATE TABLE cadoc_6334.conccred (
    id bigserial NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    year numeric(4) NOT NULL,
    quarter numeric(1) NOT NULL,
    brand numeric(2) NOT NULL,
    function varchar(1) NOT NULL,
    -- establishment job + last transaction (transaction)
    number_accredited_establishments numeric(12) NULL, 
    number_active_establishments numeric(12) NULL,
    -- transaction
    transaction_amount numeric(15, 2) NULL,
    transaction_quantity numeric(12) NULL,
    CONSTRAINT cadoc_6334_conccred_pkey PRIMARY KEY (id),
    CONSTRAINT unique_cadoc_6334_conccred UNIQUE (year, quarter, brand, function)
);

CREATE TABLE cadoc_6334.infresta (
    id bigserial NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    year numeric(4) NOT NULL,
    quarter numeric(1) NOT NULL,
    federation_unit varchar(2) NOT NULL,
    establishment_total_quantity numeric(12) NOT NULL,
    establishment_manual_capture_quantity numeric(12) NOT NULL,
    establishment_eletronic_capture_quantity numeric(12) NOT NULL,
    establishment_remote_capture_quantity numeric(12) NOT NULL,
    CONSTRAINT cadoc_6334_infresta_pkey PRIMARY KEY (id),
    CONSTRAINT unique_cadoc_6334_infresta UNIQUE (year, quarter, federation_unit)
);

CREATE TABLE cadoc_6334.infrterm (
    id bigserial NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    year numeric(4) NOT NULL,
    quarter numeric(1) NOT NULL,
    federation_unit varchar(2) NOT NULL,
    pos_total_quantity numeric(12) NOT NULL, -- igual a pos_shared_quantity + pos_chip_quantity
    pos_shared_quantity numeric(12) NOT NULL DEFAULT 0,
    pos_chip_quantity numeric(12) NOT NULL,
    pdv_quantity numeric(12) NOT NULL,
    CONSTRAINT cadoc_6334_infrterm_pkey PRIMARY KEY (id),
    CONSTRAINT unique_cadoc_6334_infrterm UNIQUE (year, quarter, federation_unit)
);

CREATE TABLE cadoc_6334.luccred (
    id bigserial NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    year numeric(4) NOT NULL,
    quarter numeric(1) NOT NULL,
    gross_revenue numeric(15,2) NOT NULL,
    rental_revenue numeric(15,2) NOT NULL,
    others_revenue numeric(15,2) NOT NULL,
    interchange_cost numeric(15,2) NOT NULL,
    marketing_cost numeric(15,2) NOT NULL,
    brand_access_cost numeric(15,2) NOT NULL,
    risk_cost numeric(15,2) NOT NULL,
    processing_cost numeric(15,2) NOT NULL,
    others_cost(15,2) NOT NULL,
    CONSTRAINT cadoc_6334_luccred_pkey PRIMARY KEY(id),
    CONSTRAINT cadoc_6334_luccred_uk UNIQUE(year, quarter)
);

CREATE TABLE cadoc_6334.contact (
    id bigserial NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    contact_type VARCHAR(1) not null, -- 'D' - diretor, 'T' - técnico, 'I' - instituição
    name VARCHAR(100) NOT NULL,
    position VARCHAR(100) NULL,
    phone_number int4 NOT NULL,
    email VARCHAR(100) NOT NULL,
    CONSTRAINTS cadoc_6334_contact PRIMARY KEY(id),
    CONSTRAINT cadoc_6334_luccred_uk(year, quarter, contact_type, name)
);



/*
--- algoritmo do job concred atualizacao-------------------------------------
-- pega todos os estabelecimentos
select establishment_code
       has_visa,
       has_mastercard,
       has_elo,
       has_debit
       has_credit,
       has_manual_capture,
	   has_eletronic_capture,
	   has_remote_capture,
       federation_unit,
       max(t.period_date) last_transaction
  from raw_data.establishment where accreditation_date < '<ultima_data_trimestre> + 1'
left join transaction.transaction t on t.establishment_code = raw_data.establishment.establishment_code
where t.period_date > '<ultima_data_trimestre> - 180' and t.transaction_id = 2
group by 1, 2, 3, 4, 5, 6;
-- loop
-- if has_visa and has_debit
        -- atualiza a conccred da seguinte forma
            -- update cadoc_6334.conccred 
            --    set number_acrredited_establishement += 1,
            --        number_actived_establishemnts += (1 se query.last_transaction >= '<ultima data do trimestre - 180> senão 0)
            -- where year = <ano do relatório> 
            --   and quarter = <trimetre>
            --   and brand = visa
            --   and function = debit  
-- if has_master and had_credit
        -- atualiza a conccred da seguinte forma
            -- update cadoc_6334.conccred 
            --    set number_acrredited_establishement += 1,
            --        number_actived_establishemnts += (1 se query.last_transaction >= '<ultima data do trimestre - 180> senão 0)
            -- where year = <ano do relatório> 
            --   and quarter = <trimetre>
            --   and brand = mastercard
            --   and function = debit
-- if has_elo and had_credit
        -- atualiza a conccred da seguinte forma
            -- update cadoc_6334.conccred 
            --    set number_acrredited_establishement += 1,
            --        number_actived_establishemnts += (1 se query.last_transaction >= '<ultima data do trimestre - 180> senão 0)
            -- where year = <ano do relatório> 
            --   and quarter = <trimetre>
            --   and brand = elo
-- if has_visa and had_credit
        -- atualiza a conccred da seguinte forma
            -- update cadoc_6334.conccred 
            --    set number_acrredited_establishement += 1,
            --        number_actived_establishemnts += (1 se query.last_transaction >= '<ultima data do trimestre - 180> senão 0)
            -- where year = <ano do relatório> 
            --   and quarter = <trimetre>
            --   and brand = visa
            --   and function = credit
-- if has_master and had_debit
        -- atualiza a conccred da seguinte forma
            -- update cadoc_6334.conccred 
            --    set number_acrredited_establishement += 1,
            --        number_actived_establishemnts += (1 se query.last_transaction >= '<ultima data do trimestre - 180> senão 0)
            -- where year = <ano do relatório> 
            --   and quarter = <trimetre>
            --   and brand = mastercard
            --   and function = credit
-- if has_elo and had_debit
        -- atualiza a conccred da seguinte forma    
            -- update cadoc_6334.conccred 
            --    set number_acrredited_establishement += 1,
            --        number_actived_establishemnts += (1 se query.last_transaction >= '<ultima data do trimestre - 180> senão 0)
            -- where year = <ano do relatório> 
            --   and quarter = <trimetre>
            --   and brand = elo
*/




