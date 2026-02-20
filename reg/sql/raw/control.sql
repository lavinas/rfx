-- public.extractor_process definição
 
-- Drop table
 
-- DROP TABLE public.extractor_process;
 
CREATE TABLE public.extractor_process (
	id bigserial NOT NULL,
	name varchar(50) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT extractor_process_pkey1 PRIMARY KEY (id)
);

-- public.extractor_process_reprocessing definição
-- Drop table
-- DROP TABLE public.extractor_process_reprocessing;
CREATE TABLE public.extractor_process_reprocessing (
	id bigserial NOT NULL,
	window_lag time NOT NULL,
	extractor_process_id int8 NOT NULL,
	CONSTRAINT extractor_process_reprocessing_pkey PRIMARY KEY (id),
	CONSTRAINT uk_jku6ps4wxuqve9fq43yei65ou UNIQUE (extractor_process_id)
);
-- public.extractor_process_reprocessing chaves estrangeiras
ALTER TABLE public.extractor_process_reprocessing ADD CONSTRAINT fkd4l2wf4rp6ad49l07d4x968yq FOREIGN KEY (extractor_process_id) REFERENCES public.extractor_process(id);

-- public.extractor_main_control definição
-- Drop table
-- DROP TABLE public.extractor_main_control;
CREATE TABLE public.extractor_main_control (
	id bigserial NOT NULL,
	extractor_process_id int8 NOT NULL,
	last_period_start timestamp NULL,
	last_period_end timestamp NULL,
	last_total int4 NULL,
	last_quantity int4 NULL,
	last_processing_start timestamp NULL,
	last_processing_end timestamp NULL,
	last_status varchar(20) NULL,
	last_trace_id varchar(50) NULL,
	status_id int4 NOT NULL,
	status_name varchar(20) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	last_error_message varchar(300) NULL,
	CONSTRAINT extractor_main_control_extractor_process_id_key UNIQUE (extractor_process_id),
	CONSTRAINT extractor_main_control_pkey PRIMARY KEY (id)
);
-- public.extractor_main_control chaves estrangeiras
ALTER TABLE public.extractor_main_control ADD CONSTRAINT extractor_main_control_extractor_process_id_fkey FOREIGN KEY (extractor_process_id) REFERENCES public.extractor_process(id);

-- public.extractor_reprocessing_control definição
-- Drop table
-- DROP TABLE public.extractor_reprocessing_control;
CREATE TABLE public.extractor_reprocessing_control (
	id bigserial NOT NULL,
	extractor_process_id int8 NOT NULL,
	required_period_start date NOT NULL,
	required_period_end date NOT NULL,
	last_period_start timestamp NULL,
	last_period_end timestamp NULL,
	last_processing_start timestamp NULL,
	last_processing_end timestamp NULL,
	last_total int4 NULL,
	last_quantity int4 NULL,
	last_status varchar(20) NULL,
	last_trace_id varchar(50) NULL,
	status_id int4 NOT NULL,
	status_name varchar(20) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	required_trace_id varchar(50) NULL,
	user_name varchar(50) NULL,
	last_error_message varchar(300) NULL,
	CONSTRAINT extractor_reprocessing_control_pkey PRIMARY KEY (id)
);
-- public.extractor_reprocessing_control chaves estrangeiras
ALTER TABLE public.extractor_reprocessing_control ADD CONSTRAINT extractor_reprocessing_control_extractor_process_id_fkey FOREIGN KEY (extractor_process_id) REFERENCES public.extractor_process(id);


-- public.extractor_execution definição
-- Drop table
-- DROP TABLE public.extractor_execution;
CREATE TABLE public.extractor_execution (
	id bigserial NOT NULL,
	control_id int8 NOT NULL,
	execution_type varchar(20) NOT NULL,
	trace_id varchar(50) NOT NULL,
	period_start timestamp NOT NULL,
	period_end timestamp NOT NULL,
	execution_total int4 DEFAULT 0 NOT NULL,
	execution_quantity int4 DEFAULT 0 NOT NULL,
	execution_start timestamp NULL,
	execution_end timestamp NULL,
	status_id int4 NOT NULL,
	status_name varchar(20) NOT NULL,
	error_message varchar(300) NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	file_name varchar(300) NULL,
	required_trace_id varchar(50) NULL,
	CONSTRAINT extractor_execution_pkey1 PRIMARY KEY (id)
);
