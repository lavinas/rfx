
-- public.process_execution - tabela que registra a execucao diaria dos processos
-- alteracoes:
--  renomear tabela de process_daily_processing para monitoring
--  uk (reference_date, process_id) ao inves de (process_id, reference_date)
--  adicionar campos de status de processamento e status de indicadores
--  remover campo status geral do processamento
--  criar indices para process_id
-- CREATE TABLE public.process_daily_processing (
CREATE TABLE public.monitoring (
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
	CONSTRAINT monitoring_pkey PRIMARY KEY (id),
	CONSTRAINT uk_monitoring UNIQUE (reference_date, process_id),
	CONSTRAINT fk_monitoring_process FOREIGN KEY (process_id) REFERENCES public.process(id) ON DELETE CASCADE
);
CREATE INDEX idx_monitoring_process_id ON public.monitoring USING btree (process_id);
CREATE INDEX idx_monitoring_reference_date ON public.monitoring USING btree (reference_date);


-- public.monitor_execution_indicator - tabela que registra os indicadores associados a execucao diaria dos processos
-- alteracoes:
--  renomear tabela de process_indicator_processing para monitor_execution_indicator
--  uk (process_indicator_id, reference_date)
--  aponta para o monitoring.id - é referente a um item de monitoramento
-- CREATE TABLE public.process_indicator_processing (
CREATE TABLE public.monitoring_indicator (
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
	CONSTRAINT fk_monitoring_indicator_process_indicator FOREIGN KEY (process_indicator_id) REFERENCES public.process_indicator(id) ON DELETE CASCADE,
    CONSTRAINT fk_monitoring_indicator_monitoring FOREIGN KEY (monitoring_id) REFERENCES public.monitoring(id) ON DELETE CASCADE
);
CREATE INDEX idx_monitoring_indicator_process_indicator ON public.monitoring_indicator USING btree (process_indicator_id);
CREATE INDEX idx_monitoring_indicator_monitoring ON public.monitoring_indicator USING btree (monitoring_id);

