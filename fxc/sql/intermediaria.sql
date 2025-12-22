

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



