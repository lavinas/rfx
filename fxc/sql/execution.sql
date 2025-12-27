-- Active: 1766004629598@@127.0.0.1@5433@flx

-- new.process_execution - tabela que registra a execucao diaria dos processos
-- alteracoes:
--  renomear tabela de process_daily_processing para monitoring
--  uk (reference_date, process_id) ao inves de (process_id, reference_date)
--  adicionar campos de status de processamento e status de indicadores
--  remover campo status geral do processamento
--  criar indices para process_id
--  incluir horarios limites esperados e reais de processamento
-- CREATE TABLE new.process_daily_processing (
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
	CONSTRAINT fk_monitoring_process FOREIGN KEY (process_id) REFERENCES new.process(id) ON DELETE CASCADE
);
CREATE INDEX idx_monitoring_process_id ON new.monitoring USING btree (process_id);
CREATE INDEX idx_monitoring_reference_date ON new.monitoring USING btree (reference_date);

-- pyblic.monitor_execution_event - tabela que registra os eventos de execucao dos processos
-- alteracoes:
--  renomear tabela de process_event para monitoring_event
--  apontar para monitoring.id ao inves de process_daily_processing.id
CREATE TABLE new.monitoring_event (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	process_name varchar(100) NOT NULL, -- redundante, mas facilita consultas e relatórios
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
CREATE INDEX idx_process_event_correlation ON new.process_event USING btree (correlation_id);
CREATE UNIQUE INDEX process_event_trace_id_idx ON new.process_event USING btree (trace_id);



-- new.monitor_execution_indicator - tabela que registra os indicadores associados a execucao diaria dos processos
-- alteracoes:
--  renomear tabela de process_indicator_processing para monitor_execution_indicator
--  uk (process_indicator_id, reference_date)
--  aponta para o monitoring.id - é referente a um item de monitoramento
-- CREATE TABLE new.process_indicator_processing (
CREATE TABLE new.monitoring_indicator (
    -- id
	id bigserial NOT NULL,
    -- reference date
	reference_date date NOT NULL, -- redundante, pois já existe na tabela monitoring (se puder, retirar ou senao pode manter)
    -- indicator values
	origin_value numeric(10, 2) NULL,
	target_value numeric(10, 2) NULL,
    -- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    -- foreign keys
	process_indicator_id int8 NOT NULL,
    monitoring_id int8 NOT NULL,
    -- indicator execution status details
	status_id int4 NOT NULL,
	status_name varchar(255) NOT NULL,
	remarks text NULL,
    -- constraints
	CONSTRAINT uk_monitoring_indicator UNIQUE (process_indicator_id, reference_date),
	CONSTRAINT monitoring_indicator_pkey PRIMARY KEY (id),
	CONSTRAINT fk_monitoring_indicator_process_indicator FOREIGN KEY (process_indicator_id) REFERENCES new.process_indicator(id) ON DELETE CASCADE,
    CONSTRAINT fk_monitoring_indicator_monitoring FOREIGN KEY (monitoring_id) REFERENCES new.monitoring(id) ON DELETE CASCADE
);
CREATE INDEX idx_monitoring_indicator_process_indicator ON new.monitoring_indicator USING btree (process_indicator_id);
CREATE INDEX idx_monitoring_indicator_monitoring ON new.monitoring_indicator USING btree (monitoring_id);

--- new.monitoring_indicator_event - tabela que relaciona os eventos de processamento com os indicadores de monitoramento
-- alteracoes:
--  renomear tabela de process_indicator_processing_event para monitoring_indicator_event
--  aponta para monitoring_indicator.id ao inves de process_indicator_processing_id
--  manter event_id apontando para process_event.id
--  criar indices para event_id e monitoring_indicator_id
-- CREATE TABLE new.process_indicator_processing_event (
CREATE TABLE new.monitoring_indicator_event (
	-- id
	id bigserial NOT NULL,
	-- foreign keys
	event_id int8 NOT NULL,
	monitoring_indicator_id int8 NOT NULL,
	-- indicator values
	origin_value numeric(10, 2) NOT NULL, 
	target_value numeric(10, 2) NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	-- status details
	status_id int4 NOT NULL,
	status_name varchar(255) NOT NULL,
	remarks text NULL,
	-- constraints
	CONSTRAINT monitoring_indicator_event_pkey PRIMARY KEY (id),
	CONSTRAINT fk_monitoring_indicator_event FOREIGN KEY (event_id) REFERENCES new.process_event(id) ON DELETE CASCADE,
	CONSTRAINT fk_monitoring_indicator_event_indicator FOREIGN KEY (monitoring_indicator_id) REFERENCES new.monitoring_indicator(id) ON DELETE CASCADE
);
CREATE INDEX idx_monitoring_indicator_event_event ON new.monitoring_indicator_event USING btree (event_id);
CREATE INDEX idx_monitoring_indicator_event_indicator ON new.monitoring_indicator_event USING btree (monitoring_indicator_id);