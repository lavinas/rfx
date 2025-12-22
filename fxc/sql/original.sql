-- DROP SEQUENCE public.batch_job_execution_seq;

CREATE SEQUENCE public.batch_job_execution_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.batch_job_seq;

CREATE SEQUENCE public.batch_job_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.batch_step_execution_seq;

CREATE SEQUENCE public.batch_step_execution_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_daily_processing_id_seq;

CREATE SEQUENCE public.process_daily_processing_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_error_history_id_seq;

CREATE SEQUENCE public.process_error_history_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_error_id_seq;

CREATE SEQUENCE public.process_error_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_event_call_message_id_seq;

CREATE SEQUENCE public.process_event_call_message_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_event_call_status_id_seq;

CREATE SEQUENCE public.process_event_call_status_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_event_error_id_seq;

CREATE SEQUENCE public.process_event_error_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_event_id_seq;

CREATE SEQUENCE public.process_event_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_event_source_id_seq;

CREATE SEQUENCE public.process_event_source_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_event_status_history_id_seq;

CREATE SEQUENCE public.process_event_status_history_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_id_seq;

CREATE SEQUENCE public.process_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_indicator_history_id_seq;

CREATE SEQUENCE public.process_indicator_history_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_indicator_id_seq;

CREATE SEQUENCE public.process_indicator_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_indicator_processing_event_id_seq;

CREATE SEQUENCE public.process_indicator_processing_event_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_indicator_processing_id_seq;

CREATE SEQUENCE public.process_indicator_processing_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_message_history_id_seq;

CREATE SEQUENCE public.process_message_history_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_message_id_seq;

CREATE SEQUENCE public.process_message_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_time_limit_history_id_seq;

CREATE SEQUENCE public.process_time_limit_history_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.process_time_limit_id_seq;

CREATE SEQUENCE public.process_time_limit_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.report_gestao_id_seq;

CREATE SEQUENCE public.report_gestao_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.report_intercambio_id_seq;

CREATE SEQUENCE public.report_intercambio_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.report_tc57_id_seq;

CREATE SEQUENCE public.report_tc57_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;
-- DROP SEQUENCE public.report_ws_id_seq;

CREATE SEQUENCE public.report_ws_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;-- public.batch_job_instance definição

-- Drop table

-- DROP TABLE public.batch_job_instance;

CREATE TABLE public.batch_job_instance (
	job_instance_id int8 NOT NULL,
	"version" int8 NULL,
	job_name varchar(100) NOT NULL,
	job_key varchar(32) NOT NULL,
	CONSTRAINT batch_job_instance_pkey PRIMARY KEY (job_instance_id),
	CONSTRAINT job_inst_un UNIQUE (job_name, job_key)
);


-- public.process definição

-- Drop table

-- DROP TABLE public.process;

CREATE TABLE public.process (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	process_name varchar(100) NOT NULL,
	flow_id int8 DEFAULT 1 NOT NULL,
	flow_name varchar(100) DEFAULT 'Clearance'::character varying NOT NULL,
	source_id int8 DEFAULT 1 NOT NULL,
	source_name varchar(100) DEFAULT 'FIS'::character varying NOT NULL,
	process_description text DEFAULT 'Processamento'::text NOT NULL,
	CONSTRAINT process_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_process_flow_id ON public.process USING btree (flow_id);
CREATE INDEX idx_process_source_id ON public.process USING btree (source_id);


-- public.report_gestao definição

-- Drop table

-- DROP TABLE public.report_gestao;

CREATE TABLE public.report_gestao (
	id bigserial NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	agenda_status varchar(255) NULL,
	arranjo_pagamento varchar(30) NULL,
	cap_amount numeric(18, 2) NULL,
	cap_qty int4 NULL,
	cap_status bool NULL,
	cd_banco int2 NULL,
	cd_mcc int4 NULL,
	cd_pessoa int8 NULL,
	chave_join varchar(255) NULL,
	descricao_mcc varchar(150) NULL,
	dt_agenda date NULL,
	dt_efetivacao_pagamento date NULL,
	dt_processamento timestamp NULL,
	fin_amount numeric(18, 2) NULL,
	fin_qty int4 NULL,
	fin_status bool NULL,
	gross_amount numeric(18, 2) NULL,
	id_terminal varchar(16) NULL,
	installment_gross_amount numeric(18, 2) NULL,
	installment_net_amount numeric(16, 6) NULL,
	installment_number int2 NULL,
	installment_tax_amount numeric(16, 6) NULL,
	liquidacao_status varchar(255) NULL,
	liquidacao_tipo varchar(255) NULL,
	net_amount numeric(18, 2) NULL,
	nm_banco varchar(100) NULL,
	nm_fantasia varchar(255) NULL,
	razao_social varchar(150) NULL,
	ref_num_bnd varchar(512) NULL,
	schedule_qty int4 NULL,
	schedule_status bool NULL,
	tax_amount numeric(18, 2) NULL,
	tax_mdr_applied numeric(18, 2) NULL,
	tax_rav_ec numeric(18, 2) NULL,
	tp_pessoa bpchar(1) NULL,
	CONSTRAINT report_gestao_pkey PRIMARY KEY (id)
);


-- public.report_intercambio definição

-- Drop table

-- DROP TABLE public.report_intercambio;

CREATE TABLE public.report_intercambio (
	id bigserial NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	bandeira varchar(255) NULL,
	bin varchar(255) NULL,
	chave_join varchar(255) NULL,
	dt_fis date NULL,
	due_date date NULL,
	emissor varchar(255) NULL,
	exchange_amount numeric(18, 2) NULL,
	exchange_rate_amount numeric(10, 4) NULL,
	exchange_rate_type int4 NULL,
	installment_amount numeric(18, 2) NULL,
	installment_net_amount numeric(18, 2) NULL,
	installment_number int4 NULL,
	origin_file_name varchar(50) NULL,
	outgoing_file_name varchar(50) NULL,
	outgoing_installment_amount numeric(18, 2) NULL,
	outgoing_qty int4 NULL,
	outgoing_transacation_amount numeric(18, 2) NULL,
	payment_scheduled_date date NULL,
	qtd_parc int4 NULL,
	ref_num_bnd varchar(255) NULL,
	release_date date NULL,
	schedule_qty int4 NULL,
	schedule_status bool NULL,
	scheduled_date date NULL,
	status_name varchar(20) NULL,
	term_id varchar(16) NULL,
	term_loc varchar(8) NULL,
	tipo_apresentacao varchar(255) NULL,
	tran_amount numeric(18, 2) NULL,
	tran_qty int4 NULL,
	tran_status bool NULL,
	transaction_date date NULL,
	CONSTRAINT report_intercambio_pkey PRIMARY KEY (id)
);


-- public.report_tc57 definição

-- Drop table

-- DROP TABLE public.report_tc57;

CREATE TABLE public.report_tc57 (
	id bigserial NOT NULL,
	ref_num_fis varchar(20) NULL,
	tc57_file_name varchar(2000) NULL,
	ref_num_bnd varchar(20) NULL,
	chave_join varchar NULL,
	bandeira varchar(20) NULL,
	tipo_transacao varchar(200) NULL,
	dt_fis timestamp NULL,
	dt_pos timestamp NULL,
	term_loc int8 NULL,
	term_id varchar(8) NULL,
	tran_amount numeric(19, 2) NULL,
	qtd_parc int4 NULL,
	bin varchar(8) NULL,
	destino varchar(6) NULL,
	chave_join2 varchar NULL,
	chave_join_intercambio varchar NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT report_tc57_pkey PRIMARY KEY (id)
);


-- public.report_ws definição

-- Drop table

-- DROP TABLE public.report_ws;

CREATE TABLE public.report_ws (
	id bigserial NOT NULL,
	ref_num_fis varchar(12) NULL,
	ref_num_bnd text NULL,
	chave_join text NULL,
	bandeira varchar(200) NULL,
	tipo_transacao varchar(200) NULL,
	dt_fis timestamp NULL,
	dt_pos timestamp NULL,
	term_loc varchar(15) NULL,
	term_id varchar(8) NULL,
	tran_amount numeric(19, 2) NULL,
	qtd_parc int2 NULL,
	bin text NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT report_ws_pkey PRIMARY KEY (id)
);


-- public.batch_job_execution definição

-- Drop table

-- DROP TABLE public.batch_job_execution;

CREATE TABLE public.batch_job_execution (
	job_execution_id int8 NOT NULL,
	"version" int8 NULL,
	job_instance_id int8 NOT NULL,
	create_time timestamp NOT NULL,
	start_time timestamp NULL,
	end_time timestamp NULL,
	status varchar(10) NULL,
	exit_code varchar(2500) NULL,
	exit_message varchar(2500) NULL,
	last_updated timestamp NULL,
	job_configuration_location varchar(2500) NULL,
	CONSTRAINT batch_job_execution_pkey PRIMARY KEY (job_execution_id),
	CONSTRAINT job_inst_exec_fk FOREIGN KEY (job_instance_id) REFERENCES public.batch_job_instance(job_instance_id)
);


-- public.batch_job_execution_context definição

-- Drop table

-- DROP TABLE public.batch_job_execution_context;

CREATE TABLE public.batch_job_execution_context (
	job_execution_id int8 NOT NULL,
	short_context varchar(2500) NOT NULL,
	serialized_context text NULL,
	CONSTRAINT batch_job_execution_context_pkey PRIMARY KEY (job_execution_id),
	CONSTRAINT job_exec_ctx_fk FOREIGN KEY (job_execution_id) REFERENCES public.batch_job_execution(job_execution_id)
);


-- public.batch_job_execution_params definição

-- Drop table

-- DROP TABLE public.batch_job_execution_params;

CREATE TABLE public.batch_job_execution_params (
	job_execution_id int8 NOT NULL,
	type_cd varchar(6) NOT NULL,
	key_name varchar(100) NOT NULL,
	string_val varchar(250) NULL,
	date_val timestamp NULL,
	long_val int8 NULL,
	double_val float8 NULL,
	identifying bpchar(1) NOT NULL,
	CONSTRAINT job_exec_params_fk FOREIGN KEY (job_execution_id) REFERENCES public.batch_job_execution(job_execution_id)
);


-- public.batch_step_execution definição

-- Drop table

-- DROP TABLE public.batch_step_execution;

CREATE TABLE public.batch_step_execution (
	step_execution_id int8 NOT NULL,
	"version" int8 NOT NULL,
	step_name varchar(100) NOT NULL,
	job_execution_id int8 NOT NULL,
	start_time timestamp NOT NULL,
	end_time timestamp NULL,
	status varchar(10) NULL,
	commit_count int8 NULL,
	read_count int8 NULL,
	filter_count int8 NULL,
	write_count int8 NULL,
	read_skip_count int8 NULL,
	write_skip_count int8 NULL,
	process_skip_count int8 NULL,
	rollback_count int8 NULL,
	exit_code varchar(2500) NULL,
	exit_message varchar(2500) NULL,
	last_updated timestamp NULL,
	CONSTRAINT batch_step_execution_pkey PRIMARY KEY (step_execution_id),
	CONSTRAINT job_exec_step_fk FOREIGN KEY (job_execution_id) REFERENCES public.batch_job_execution(job_execution_id)
);


-- public.batch_step_execution_context definição

-- Drop table

-- DROP TABLE public.batch_step_execution_context;

CREATE TABLE public.batch_step_execution_context (
	step_execution_id int8 NOT NULL,
	short_context varchar(2500) NOT NULL,
	serialized_context text NULL,
	CONSTRAINT batch_step_execution_context_pkey PRIMARY KEY (step_execution_id),
	CONSTRAINT step_exec_ctx_fk FOREIGN KEY (step_execution_id) REFERENCES public.batch_step_execution(step_execution_id)
);


-- public.process_daily_processing definição

-- Drop table

-- DROP TABLE public.process_daily_processing;

CREATE TABLE public.process_daily_processing (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	reference_date date NOT NULL,
	process_id int8 NOT NULL,
	status_id int4 NOT NULL,
	status_name varchar(100) NOT NULL,
	remarks text NULL,
	CONSTRAINT process_daily_processing_pkey PRIMARY KEY (id),
	CONSTRAINT uk_process_daily_processing UNIQUE (process_id, reference_date),
	CONSTRAINT fk_process_daily_processing_process FOREIGN KEY (process_id) REFERENCES public.process(id) ON DELETE CASCADE
);


-- public.process_error definição

-- Drop table

-- DROP TABLE public.process_error;

CREATE TABLE public.process_error (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	error_key varchar(100) NOT NULL,
	generate_call bool NOT NULL,
	message_body text NULL,
	process_id int8 DEFAULT 1 NOT NULL,
	description text NULL,
	CONSTRAINT process_error_pkey PRIMARY KEY (id),
	CONSTRAINT uk_process_error_error_key UNIQUE (error_key),
	CONSTRAINT fkoht3cb97lmqyibhpjqxtxilub FOREIGN KEY (process_id) REFERENCES public.process(id)
);


-- public.process_error_history definição

-- Drop table

-- DROP TABLE public.process_error_history;

CREATE TABLE public.process_error_history (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	field varchar(100) NOT NULL,
	previous_value varchar(255) NOT NULL,
	current_value varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	process_error_id int8 NOT NULL,
	CONSTRAINT process_error_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_error_history_error FOREIGN KEY (process_error_id) REFERENCES public.process_error(id) ON DELETE CASCADE
);


-- public.process_event definição

-- Drop table

-- DROP TABLE public.process_event;

CREATE TABLE public.process_event (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	process_name varchar(100) NOT NULL,
	trace_id varchar(35) NULL,
	status int4 NOT NULL,
	status_name varchar(20) NOT NULL,
	started_at timestamp NOT NULL,
	finished_at timestamp NOT NULL,
	errors_count int4 DEFAULT 0 NOT NULL,
	reprocess_count int4 DEFAULT 0 NOT NULL,
	process_id int8 NOT NULL,
	correlation_id varchar(64) NULL,
	broker_sent_datetime timestamp NULL,
	broker_status varchar(20) DEFAULT 'SENT'::character varying NULL,
	process_status_id int4 NOT NULL,
	process_status_name varchar(100) NOT NULL,
	remarks text NULL,
	process_daily_id int8 NOT NULL,
	CONSTRAINT process_event_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_event_process FOREIGN KEY (process_id) REFERENCES public.process(id) ON DELETE CASCADE,
	CONSTRAINT fk_process_event_process_daily FOREIGN KEY (process_daily_id) REFERENCES public.process_daily_processing(id) ON DELETE CASCADE
);
CREATE INDEX idx_process_event_correlation ON public.process_event USING btree (correlation_id);
CREATE UNIQUE INDEX process_event_trace_id_idx ON public.process_event USING btree (trace_id);


-- public.process_event_call_message definição

-- Drop table

-- DROP TABLE public.process_event_call_message;

CREATE TABLE public.process_event_call_message (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	message_subject varchar(100) NOT NULL,
	message_body text NOT NULL,
	process_event_id int8 NOT NULL,
	CONSTRAINT process_event_call_message_pkey PRIMARY KEY (id),
	CONSTRAINT uk_process_event_call_message UNIQUE (process_event_id),
	CONSTRAINT fk_process_event_call_message_event FOREIGN KEY (process_event_id) REFERENCES public.process_event(id) ON DELETE CASCADE
);


-- public.process_event_call_status definição

-- Drop table

-- DROP TABLE public.process_event_call_status;

CREATE TABLE public.process_event_call_status (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	status_code int4 NOT NULL,
	status_description varchar(500) NOT NULL,
	call_code varchar(100) NULL,
	process_event_id int8 NOT NULL,
	CONSTRAINT process_event_call_status_pkey PRIMARY KEY (id),
	CONSTRAINT uk_process_event_call_status_history UNIQUE (process_event_id),
	CONSTRAINT fk_process_event_call_status_history FOREIGN KEY (process_event_id) REFERENCES public.process_event(id) ON DELETE CASCADE
);


-- public.process_event_error definição

-- Drop table

-- DROP TABLE public.process_event_error;

CREATE TABLE public.process_event_error (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	id_error varchar(100) NOT NULL,
	field_name varchar(100) NOT NULL,
	line_number int4 NOT NULL,
	line varchar(1500) NOT NULL,
	"position" int4 NOT NULL,
	"size" int4 NOT NULL,
	description text NOT NULL,
	process_event_id int8 NOT NULL,
	process_error_id int8 NULL,
	CONSTRAINT process_event_error_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_event_error_error FOREIGN KEY (process_error_id) REFERENCES public.process_error(id) ON DELETE SET NULL,
	CONSTRAINT fk_process_event_error_event FOREIGN KEY (process_event_id) REFERENCES public.process_event(id) ON DELETE CASCADE
);


-- public.process_event_source definição

-- Drop table

-- DROP TABLE public.process_event_source;

CREATE TABLE public.process_event_source (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	"type" varchar(20) NOT NULL,
	"name" varchar(50) NOT NULL,
	lines int4 NOT NULL,
	"path" varchar(256) NOT NULL,
	process_event_id int8 NOT NULL,
	CONSTRAINT process_event_source_pkey PRIMARY KEY (id),
	CONSTRAINT uk_process_event_source UNIQUE (process_event_id),
	CONSTRAINT fk_process_event_source_event FOREIGN KEY (process_event_id) REFERENCES public.process_event(id) ON DELETE CASCADE
);


-- public.process_event_status_history definição

-- Drop table

-- DROP TABLE public.process_event_status_history;

CREATE TABLE public.process_event_status_history (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	status int4 NOT NULL,
	status_name varchar(100) NOT NULL,
	description varchar(255) NULL,
	process_event_id int8 NOT NULL,
	CONSTRAINT process_event_status_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_event_broker_history_status FOREIGN KEY (process_event_id) REFERENCES public.process_event(id) ON DELETE CASCADE
);


-- public.process_indicator definição

-- Drop table

-- DROP TABLE public.process_indicator;

CREATE TABLE public.process_indicator (
	id bigserial NOT NULL,
	"name" varchar(255) NOT NULL,
	process_origin_id int8 NOT NULL,
	process_target_id int8 NOT NULL,
	origin_under_var numeric(5, 4) NULL,
	origin_over_var numeric(5, 4) NULL,
	target_under_var numeric(5, 4) NULL,
	target_over_var numeric(5, 4) NULL,
	message_body text NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	CONSTRAINT process_indicator_pkey PRIMARY KEY (id),
	CONSTRAINT process_indicators_uk UNIQUE (name, process_origin_id, process_target_id),
	CONSTRAINT fk_pi_origin FOREIGN KEY (process_origin_id) REFERENCES public.process(id) ON DELETE CASCADE,
	CONSTRAINT fk_pi_target FOREIGN KEY (process_target_id) REFERENCES public.process(id) ON DELETE CASCADE
);


-- public.process_indicator_history definição

-- Drop table

-- DROP TABLE public.process_indicator_history;

CREATE TABLE public.process_indicator_history (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	field varchar(100) NOT NULL,
	previous_value varchar(255) NOT NULL,
	current_value varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	process_indicator_id int8 NOT NULL,
	CONSTRAINT process_indicator_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_indicator_history_indicator FOREIGN KEY (process_indicator_id) REFERENCES public.process_indicator(id) ON DELETE CASCADE
);


-- public.process_indicator_processing definição

-- Drop table

-- DROP TABLE public.process_indicator_processing;

CREATE TABLE public.process_indicator_processing (
	id bigserial NOT NULL,
	reference_date date NOT NULL,
	origin_value numeric(10, 2) NULL,
	target_value numeric(10, 2) NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	process_indicator_id int8 NOT NULL,
	status_id int4 NOT NULL,
	status_name varchar(255) NOT NULL,
	remarks text NULL,
	CONSTRAINT pip_uk UNIQUE (process_indicator_id, reference_date),
	CONSTRAINT process_indicator_processing_pkey PRIMARY KEY (id),
	CONSTRAINT fk_pip_pi FOREIGN KEY (process_indicator_id) REFERENCES public.process_indicator(id) ON DELETE CASCADE
);


-- public.process_indicator_processing_event definição

-- Drop table

-- DROP TABLE public.process_indicator_processing_event;

CREATE TABLE public.process_indicator_processing_event (
	id bigserial NOT NULL,
	event_id int8 NOT NULL,
	process_indicator_processing_id int8 NOT NULL,
	origin_value numeric(10, 2) NOT NULL,
	target_value numeric(10, 2) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	status_id int4 NOT NULL,
	status_name varchar(255) NOT NULL,
	remarks text NULL,
	CONSTRAINT process_indicator_processing_event_pkey PRIMARY KEY (id),
	CONSTRAINT fk_pipe_event FOREIGN KEY (event_id) REFERENCES public.process_event(id) ON DELETE CASCADE,
	CONSTRAINT fk_pipe_pip FOREIGN KEY (process_indicator_processing_id) REFERENCES public.process_indicator_processing(id) ON DELETE CASCADE
);


-- public.process_message definição

-- Drop table

-- DROP TABLE public.process_message;

CREATE TABLE public.process_message (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	error_body text NOT NULL,
	error_subject varchar(100) NOT NULL,
	timeout_body text NOT NULL,
	timeout_subject varchar(100) NOT NULL,
	process_id int8 NOT NULL,
	CONSTRAINT process_message_pkey PRIMARY KEY (id),
	CONSTRAINT uk_gg3rl1rb3hfymghlmk7pq7ah4 UNIQUE (process_id),
	CONSTRAINT fk1y0kevilbg6hodcsm5otbe5t1 FOREIGN KEY (process_id) REFERENCES public.process(id)
);


-- public.process_message_history definição

-- Drop table

-- DROP TABLE public.process_message_history;

CREATE TABLE public.process_message_history (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	field varchar(100) NOT NULL,
	previous_value varchar(255) NOT NULL,
	current_value varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	process_message_id int8 NOT NULL,
	CONSTRAINT process_message_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_message_history_message FOREIGN KEY (process_message_id) REFERENCES public.process_message(id) ON DELETE CASCADE
);


-- public.process_time_limit definição

-- Drop table

-- DROP TABLE public.process_time_limit;

CREATE TABLE public.process_time_limit (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	periodicity varchar(50) NOT NULL,
	time_limit time NOT NULL,
	process_id int8 NOT NULL,
	CONSTRAINT process_time_limit_pkey PRIMARY KEY (id),
	CONSTRAINT uk_process_time_limit UNIQUE (process_id),
	CONSTRAINT fk_process_time_limit_process FOREIGN KEY (process_id) REFERENCES public.process(id) ON DELETE CASCADE
);


-- public.process_time_limit_history definição

-- Drop table

-- DROP TABLE public.process_time_limit_history;

CREATE TABLE public.process_time_limit_history (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	field varchar(100) NOT NULL,
	previous_value varchar(255) NULL,
	current_value varchar(255) NULL,
	username varchar(255) NOT NULL,
	process_time_limit_id int8 NOT NULL,
	CONSTRAINT process_time_limit_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_time_limit_history_time_limit FOREIGN KEY (process_time_limit_id) REFERENCES public.process_time_limit(id) ON DELETE CASCADE
);