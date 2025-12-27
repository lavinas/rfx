CREATE TABLE old.process (
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
CREATE TABLE old.process_error (
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
	CONSTRAINT fkoht3cb97lmqyibhpjqxtxilub FOREIGN KEY (process_id) REFERENCES old.process(id)
);
CREATE TABLE old.process_error_history (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	field varchar(100) NOT NULL,
	previous_value varchar(255) NOT NULL,
	current_value varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	process_error_id int8 NOT NULL,
	CONSTRAINT process_error_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_error_history_error FOREIGN KEY (process_error_id) REFERENCES old.process_error(id) ON DELETE CASCADE
);
CREATE TABLE old.process_indicator (
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
	CONSTRAINT fk_pi_origin FOREIGN KEY (process_origin_id) REFERENCES old.process(id) ON DELETE CASCADE,
	CONSTRAINT fk_pi_target FOREIGN KEY (process_target_id) REFERENCES old.process(id) ON DELETE CASCADE
);
CREATE TABLE old.process_indicator_history (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	field varchar(100) NOT NULL,
	previous_value varchar(255) NOT NULL,
	current_value varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	process_indicator_id int8 NOT NULL,
	CONSTRAINT process_indicator_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_indicator_history_indicator FOREIGN KEY (process_indicator_id) REFERENCES old.process_indicator(id) ON DELETE CASCADE
);
CREATE TABLE old.process_message (
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
	CONSTRAINT fk1y0kevilbg6hodcsm5otbe5t1 FOREIGN KEY (process_id) REFERENCES old.process(id)
);
CREATE TABLE old.process_message_history (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	field varchar(100) NOT NULL,
	previous_value varchar(255) NOT NULL,
	current_value varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	process_message_id int8 NOT NULL,
	CONSTRAINT process_message_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_message_history_message FOREIGN KEY (process_message_id) REFERENCES old.process_message(id) ON DELETE CASCADE
);
CREATE TABLE old.process_time_limit (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	periodicity varchar(50) NOT NULL,
	time_limit time NOT NULL,
	process_id int8 NOT NULL,
	CONSTRAINT process_time_limit_pkey PRIMARY KEY (id),
	CONSTRAINT uk_process_time_limit UNIQUE (process_id),
	CONSTRAINT fk_process_time_limit_process FOREIGN KEY (process_id) REFERENCES old.process(id) ON DELETE CASCADE
);
CREATE TABLE old.process_time_limit_history (
	id bigserial NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	field varchar(100) NOT NULL,
	previous_value varchar(255) NULL,
	current_value varchar(255) NULL,
	username varchar(255) NOT NULL,
	process_time_limit_id int8 NOT NULL,
	CONSTRAINT process_time_limit_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_time_limit_history_time_limit FOREIGN KEY (process_time_limit_id) REFERENCES old.process_time_limit(id) ON DELETE CASCADE
);