
-- public.process_execution - tabela que registra a execucao diaria dos processos
-- alteracoes:
--  renomear tabela de process_daily_processing para monitor_execution
--  uk (reference_date, process_id) ao inves de (process_id, reference_date)
--  adicionar campos de status de processamento e status de indicadores
--  remover campo status geral do processamento
--  criar indices para process_id
-- CREATE TABLE public.process_daily_processing (
CREATE TABLE public.monitor_execution ( -- talvez monitoring
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
    -- indicators status details
    indicators_status_id int4 NOT NULL,
    indicators_status_name varchar(100) NOT NULL,
    indicators_remarks text NULL,
    -- constraints
	CONSTRAINT monitor_execution_pkey PRIMARY KEY (id),
	CONSTRAINT uk_monitor_execution UNIQUE (reference_date, process_id),
	CONSTRAINT fk_monitor_execution_process FOREIGN KEY (process_id) REFERENCES public.process(id) ON DELETE CASCADE
);
CREATE INDEX idx_monitor_execution_process ON public.monitor_execution USING btree (process_id);

-- public.monitor_execution_indicator - tabela que registra os indicadores associados a execucao diaria dos processos
-- alteracoes:
--  renomear tabela de process_indicator_processing para monitor_execution_indicator
--  uk (process_indicator_id, reference_date)
-- CREATE TABLE public.process_indicator_processing (
CREATE TABLE public.monitor_execution_indicator (
    -- id
	id bigserial NOT NULL,
    -- reference date
	reference_date date NOT NULL,
    -- indicator values
	origin_value numeric(10, 2) NULL,
	target_value numeric(10, 2) NULL,
    -- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    -- foreign keys
	process_indicator_id int8 NOT NULL,
    monitor_execution_id int8 NOT NULL,
    -- indicator execution status details
	status_id int4 NOT NULL,
	status_name varchar(255) NOT NULL,
	remarks text NULL,
    -- constraints
	CONSTRAINT uk_monitor_execution_indicator UNIQUE (process_indicator_id, reference_date),
	CONSTRAINT monitor_execution_indicator_pkey PRIMARY KEY (id),
	CONSTRAINT fk_monitor_execution_indicator_process_indicator FOREIGN KEY (process_indicator_id) REFERENCES public.process_indicator(id) ON DELETE CASCADE
    CONSTRAINT fk_monitor_execution_indicator_monitor_execution FOREIGN KEY (monitor_execution_id) REFERENCES public.monitor_execution(id) ON DELETE CASCADE
);





