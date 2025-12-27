-- Active: 1766004629598@@127.0.0.1@5433@flx


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