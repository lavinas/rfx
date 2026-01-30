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
    transaction_capture varchar(3) not null, -- POS, TEF
    -- financial values
    revenue_mdr_value numeric(15,2) null,
    cost_interchange_value numeric(15,2) null
    -- references/control
    high_source_priority int not null, -- 30 - webservice, 20 - intercambio, 10 - gestao (o de maior prioridade)
    -- status fields
    status_id int not null, -- 1 - pendente, 2 - pronto
    status_name varchar(20) not null, -- 1 - pendente, 2 - pronto
    status_count int not null default 3, -- contador de tentativas de processamento
    -- closing fields
    period_date date null -- inicialmente nulo
    period_closing_id bigserial null, -- inicialmente nulo, referencia transaction_period_closing(id)
    foreign key (period_closing_id) references transaction_period_closing(id)
);


-- consolidators table to store consolidation data
create table consolidator (
    id BIGINT PRIMARY KEY,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    name VARCHAR(100) NOT NULL,
    last_updated_at TIMESTAMP NULL,
    last_execution_at TIMESTAMP NULL
);

-- consolidator_control table to track consolidation processes
create table consolidator_control (
    id BIGSERIAL PRIMARY KEY,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    consolidator_id BIGINT NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NULL,
    consolidation_type VARCHAR(50) NOT NULL, -- normal, reprocess
    status VARCHAR(50) NOT NULL,
    FOREIGN KEY (consolidator_id) REFERENCES consolidator(id)
);
