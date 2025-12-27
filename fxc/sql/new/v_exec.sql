-- vide config
CREATE TABLE new.process (
    -- id
	id bigserial NOT NULL,
    -- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
    -- name
	process_name varchar(100) NOT NULL,
    -- classifications
	flow_id int8 DEFAULT 1 NOT NULL,
	flow_name varchar(100) NOT NULL,
	source_id int8 NOT NULL,
	source_name varchar(100) NOT NULL,
    -- description
	process_description text NOT NULL,
	-- constraints
    CONSTRAINT process_pkey PRIMARY KEY (id)
);
-- antiga tabela process_daily_processing
-- retirada a coluna process_name por ser redundante
-- separacao dos status de processamento e indicadores
-- adicao dos limites de tempo esperados e reais
-- alteracao da ordem do unique key para ter a data mais abrangente primeiro
CREATE TABLE new.monitoring (
    -- id
	id bigserial NOT NULL,
    -- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
    -- reference date
	reference_date date NOT NULL,
    -- foreign key to process
	process_id int8 NOT NULL,
    -- processing status details
	processing_status_id int4 NOT NULL,
	processing_status_name varchar(100) NOT NULL,
	processing_remarks text NULL,
	-- processing time limits
	processing_expected_limit time NOT NULL,
	processing_actual_limit time NOT NULL,
    -- indicators status details
    indicators_status_id int4 NOT NULL,
    indicators_status_name varchar(100) NOT NULL,
    indicators_remarks text NULL,
    -- constraints
	CONSTRAINT monitoring_pkey PRIMARY KEY (id),
	CONSTRAINT uk_monitoring UNIQUE (reference_date, process_id),
	CONSTRAINT fk_monitoring_process FOREIGN KEY (process_id) REFERENCES public.process(id) ON DELETE CASCADE
);
-- antiga tabela process_event
-- separacao dos status de execução e 
-- continuar daqui

CREATE TABLE new.monitoring_event (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- event details
	trace_id varchar(35) NULL,
	-- execution status details
	status int4 NOT NULL,
	status_name varchar(20) NOT NULL,
	started_at timestamp NOT NULL,
	finished_at timestamp NOT NULL,
	errors_count int4 DEFAULT 0 NOT NULL,
	reprocess_count int4 DEFAULT 0 NOT NULL,
	-- foreign keys
	process_id int8 NOT NULL, -- redundante, mas facilita consultas e relatórios
	correlation_id varchar(64) NULL,
	-- broker details
	broker_sent_datetime timestamp NULL,
	broker_status varchar(20) DEFAULT 'SENT'::character varying NULL,
	-- process status details
	process_status_id int4 NOT NULL,
	process_status_name varchar(100) NOT NULL,
	remarks text NULL,
	-- foreign key to monitoring
	monitoring_id int8 NOT NULL,
	-- constraints
	CONSTRAINT process_event_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_event_process FOREIGN KEY (process_id) REFERENCES new.process(id) ON DELETE CASCADE,
	CONSTRAINT fk_process_event_monitoring FOREIGN KEY (monitoring_id) REFERENCES new.monitoring(id) ON DELETE CASCADE
);


