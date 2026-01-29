
-- vide config
CREATE TABLE new.process (
    -- id
	id bigserial NOT NULL, -- 10
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
-- vide config
CREATE TABLE new.process_error (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- error details
	error_key varchar(100) NOT NULL,
	generate_call bool NOT NULL,
	message_body text NULL,
	-- foreign key to process
	process_id int8 DEFAULT 1 NOT NULL,
	-- description
	description text NULL,
	--error details
	CONSTRAINT process_error_pkey PRIMARY KEY (id),
	-- constraints
	-- CONSTRAINT fk_process_error_process FOREIGN KEY (process_id) REFERENCES new.process(id)
	CONSTRAINT uk_process_error_error_key UNIQUE (error_key)
);
-- vide config
CREATE TABLE new.process_indicator (
	-- id
	id bigserial NOT NULL, -- 1
	-- control
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	-- indicator details
	name varchar(255) NOT NULL,
	description varchar(255) NULL,
	under_var numeric(5, 4) NULL,
	over_var numeric(5, 4) NULL,
	message_body text NOT NULL,
	-- foreign key to process
	process_id int8 NOT NULL, -- antigo target  100 - outgoing-master-credito
	process_reference_id int8 NOT NULL, -- antigo origin 10 - tc57-master-credito
	-- constraints
	CONSTRAINT process_indicator_pkey PRIMARY KEY (id),
	--CONSTRAINT fk_pi_process_id FOREIGN KEY (process_id) REFERENCES new.process(id) ON DELETE CASCADE,
	--CONSTRAINT fk_pi_process_reference FOREIGN KEY (process_reference_id) REFERENCES new.process(id) ON DELETE CASCADE,
	CONSTRAINT process_indicators_uk UNIQUE (name, process_id, process_reference_id)
);
-- inicio das tabelas de execucao
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
    -- foreign key to processI
	process_id int8 NOT NULL,
    -- processing status details
	processing_status_id int4 NOT NULL, -- 0 - N/A, 1 - timeout, 2 - error, 4 - ok
	processing_status_name varchar(100) NOT NULL, -- 'N/A' 'timeout', 'error', 'ok'
	processing_remarks text NULL, -- mensagem corresponente
	-- processing time limits
	processing_expected_limit time NULL,
	processing_actual_limit time NULL,
    -- indicators status details
    indicators_status_id int4 NOT NULL, -- 0 - N/A, 2 - error, 3 - waiting indicators, 4 - ok
    indicators_status_name varchar(100) NOT NULL, -- 'N/A', 'error', 'waiting indicators', 'ok'
    indicators_remarks text NULL, -- mensagem corresponente
	-- indicators counts
	indicators_count int8 NOT NULL, -- total de indicadores esperados - nao subtrai 
    indicators_ok int8 NOT NULL, -- total de indicadores ok (nao soma de novo se não mudou de status)
    indicators_error int8 NOT NULL, -- total de indicadores com erro (nao soma de novo se não mudou de status)
    -- constraints
	CONSTRAINT monitoring_pkey PRIMARY KEY (id),
	CONSTRAINT uk_monitoring UNIQUE (reference_date, process_id),
	CONSTRAINT fk_monitoring_process FOREIGN KEY (process_id) REFERENCES new.process(id) ON DELETE CASCADE
);
-- antiga tabela process_indicator_processing
CREATE TABLE new.monitoring_indicator (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- foreign key to monitoring
	monitoring_id int8 NOT NULL,
	-- indicator status details
	indicator_status_id int4 NOT NULL, -- 0 - waiting_compartion, 2 - error, 3 - waiting indicators, 4 - ok
	indicator_status_name varchar(100) NOT NULL,-- 'error', 'waiting indicators', 'ok'
	indicator_remarks text NULL,
	-- received values
	process_indicator_id int8 NOT NULL,
	indicator_value numeric(20, 6) NULL, -- antigo target_value
	reference_value numeric(20, 6) NULL, -- antigo origin_value
	-- constraints
	CONSTRAINT monitoring_indicator_pkey PRIMARY KEY (id),
	CONSTRAINT uk_monitoring_indicator UNIQUE (monitoring_id, process_indicator_id),
	CONSTRAINT fk_monitoring_indicator_process_indicator FOREIGN KEY (process_indicator_id) REFERENCES new.process_indicator(id) ON DELETE CASCADE,
	CONSTRAINT fk_monitoring_indicator_monitoring FOREIGN KEY (monitoring_id) REFERENCES new.monitoring(id) ON DELETE CASCADE
);

-- antiga tabela process_event
CREATE TABLE new.monitoring_event (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- link to external traceability
	trace_id varchar(35) NOT NULL,
	-- result status
	event_type varchar(20) NOT NULL, -- e.g., 'file', 'indicator'
	event_status_id int4 NOT NULL, -- 0 - waiting-comparation, 1 - timeout (just execution), 2 - error, 3 - waiting indicator , 4 - ok
	event_status_name varchar(20) NOT NULL, -- e.g., 'timeout', 'error', 'waiting indicators', 'ok'
	remarks text NULL,
	-- exists open_call
	open_call bool DEFAULT false NOT NULL, 
	-- foreign keys
	monitoring_id int8 NOT NULL,
	-- constraints
	CONSTRAINT monitoring_event_pkey PRIMARY KEY (id),
	CONSTRAINT fk_monitoring_event_monitoring FOREIGN KEY (monitoring_id) REFERENCES new.monitoring(id) ON DELETE CASCADE
);

-- nova tabela especializada da tabela process_event para timeouts
CREATE TABLE new.monitoring_event_timeout (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- foreign key to monitoring_event
	monitoring_event_id int8 NOT NULL,
	-- timeout specific details
	expected_time time NOT NULL,
	-- constraints
	CONSTRAINT monitoring_event_timeout_pkey PRIMARY KEY (id),
	CONSTRAINT uk_monitoring_event_timeout UNIQUE (monitoring_event_id),
	CONSTRAINT fk_monitoring_event_timeout FOREIGN KEY (monitoring_event_id) REFERENCES new.monitoring_event(id) ON DELETE CASCADE
);

-- nova tabela especializada da tabela process_event para monitoramentos, porem tambem antiga tabela process_event_source
CREATE TABLE new.monitoring_event_execution (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- foreign key to monitoring_event
	monitoring_event_id int8 NOT NULL,
	-- execution specific details
	started_at timestamp NOT NULL,
	finished_at timestamp NOT NULL,
	errors_count int4 NOT NULL,
	-- source details
	source_type varchar(20) NOT NULL,
	source_name varchar(50) NOT NULL,
	source_lines int4 NOT NULL,
	source_path varchar(256) NOT NULL,
	-- constraints
	CONSTRAINT monitoring_event_execution_pkey PRIMARY KEY (id),
	CONSTRAINT uk_monitoring_event_execution UNIQUE (monitoring_event_id),
	CONSTRAINT fk_monitoring_event_execution FOREIGN KEY (monitoring_event_id) REFERENCES new.monitoring_event(id) ON DELETE CASCADE
)

-- nova tabela especializada da tabela process_event para indicadores
CREATE TABLE new.monitoring_event_indicator (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- foreign key to monitoring_event
	monitoring_event_id int8 NOT NULL,
	-- indicator specific details
	monitoring_indicator_id int8 NOT NULL,
    process_id int8 NOT NULL, 
	indicator_value numeric(20, 6) NOT NULL,
	-- constraints
	CONSTRAINT monitoring_event_indicator_pkey PRIMARY KEY (id),
	CONSTRAINT fk_monitoring_event_indicator_process FOREIGN KEY (process_id) REFERENCES new.process(id) ON DELETE CASCADE,
	CONSTRAINT uk_monitoring_event_indicator UNIQUE (monitoring_event_id),
	CONSTRAINT fk_monitoring_event_indicator_indicator FOREIGN KEY (monitoring_indicator_id) REFERENCES new.monitoring_indicator(id) ON DELETE CASCADE,
	CONSTRAINT fk_monitoring_event_indicator FOREIGN KEY (monitoring_event_id) REFERENCES new.monitoring_event(id) ON DELETE CASCADE
);

-- antiga tabela process_event_error
CREATE TABLE new.monitoring_event_execution_error (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- error details
	id_error varchar(100) NOT NULL,
	field_name varchar(100) NOT NULL,
	line_number int4 NOT NULL,
	line varchar(1500) NOT NULL,
	position int4 NOT NULL,
	size int4 NOT NULL,
	description text NOT NULL,
	-- foreign key to monitoring_event
	monitoring_event_execution_id int8 NOT NULL,
	process_error_id int8 NULL,
	-- constraints
	CONSTRAINT monitoring_event_error_pkey PRIMARY KEY (id),
	CONSTRAINT fk_monitoring_event_error_process_error FOREIGN KEY (process_error_id) REFERENCES new.process_error(id) ON DELETE SET NULL,
	CONSTRAINT fk_monitoring_event_error FOREIGN KEY (monitoring_event_execution_id) REFERENCES new.monitoring_event_execution(id) ON DELETE CASCADE
);
-- antigas tabelas process_event_call_message e process_event_call_status
CREATE TABLE new.monitoring_event_open_call (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- message details
	message_subject varchar(100) NOT NULL,
	message_body text NOT NULL,
	-- sent control
	reprocess_count int4 DEFAULT 0 NOT NULL,
	-- sent status
	status_code int4 NOT NULL, -- old: new, error, released, sent, archived, reprocessed // new: 1 - new, 2 - sent, 3 - sent error, 4 - call opened, 5 - call error, 6 - error archived
	-- o registro é criado: new
	-- envio pra fila com sucesso: new -> sent
	-- erro no envio pra fila: sent -> sent error
	-- retentativa de envio para fila com sucesso: sent error -> sent
	-- retentativa de envio para fila com erro: sent error -> sent error
	-- ticketer chamado aberto com sucesso: sent -> call opened
	-- ticketer chamado com erro: sent -> call error
	-- expiradas tentativas de reenvios: call error -> error archived ou sent error -> error archived
	status_description varchar(500) NOT NULL,
	call_code varchar(100) NULL,
	-- foreign key to monitoring_event
	monitoring_event_id int8 NOT NULL,
	-- constraints
	CONSTRAINT monitoring_event_open_call_pkey PRIMARY KEY (id),
	CONSTRAINT uk_monitoring_event_open_call UNIQUE (monitoring_event_id),
	CONSTRAINT fk_monitoring_event_open_call FOREIGN KEY (monitoring_event_id) REFERENCES new.monitoring_event(id) ON DELETE CASCADE
);