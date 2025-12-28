
create schema new;


-- nunuma alteracao estrutural
-- retirar valores default
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
-- nenhuma alteracao estrutural
CREATE TABLE new.process_time_limit (
    -- id
	id bigserial NOT NULL,
    -- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
    -- time limit details
	periodicity varchar(50) NOT NULL,
	time_limit time NOT NULL,
    -- foreign key to process
	process_id int8 NOT NULL,
    -- constraints
	CONSTRAINT process_time_limit_pkey PRIMARY KEY (id),
	CONSTRAINT uk_process_time_limit UNIQUE (process_id),
	CONSTRAINT fk_process_time_limit_process FOREIGN KEY (process_id) REFERENCES new.process(id) ON DELETE CASCADE
);
-- nenhuma alteracao estrutural
CREATE TABLE new.process_time_limit_history (
    -- id
	id bigserial NOT NULL,
    -- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
    -- change details
	field varchar(100) NOT NULL,
	previous_value varchar(255) NULL,
	current_value varchar(255) NULL,
	username varchar(255) NOT NULL,
    -- foreign key to process_time_limit
	process_time_limit_id int8 NOT NULL,
	-- constraints
	CONSTRAINT process_time_limit_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_time_limit_history_time_limit FOREIGN KEY (process_time_limit_id) REFERENCES new.process_time_limit(id) ON DELETE CASCADE
);
-- nenhuma alteracao estrutural
-- correcao do nome de constraint
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
	CONSTRAINT uk_process_error_error_key UNIQUE (error_key),
	CONSTRAINT fk_process_error_process FOREIGN KEY (process_id) REFERENCES new.process(id)
);
-- nenhuma alteracao estrutural
CREATE TABLE new.process_error_history (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- change details
	field varchar(100) NOT NULL,
	previous_value varchar(255) NOT NULL,
	current_value varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	-- foreign key to process_error
	process_error_id int8 NOT NULL,
	-- constraints
	CONSTRAINT process_error_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_error_history_error FOREIGN KEY (process_error_id) REFERENCES new.process_error(id) ON DELETE CASCADE
);
-- mudanca 1
-- normalizacao dos tipos, nomes e corpo de mensagens
-- contraimt unique para process_id e type
CREATE TABLE new.process_message (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	-- message details
	type varchar(50) NOT NULL,
	subject varchar(100) NOT NULL,
	body text NOT NULL,
	-- foreign key to process
	process_id int8 NOT NULL,
	-- constraints
	CONSTRAINT process_message_pkey PRIMARY KEY (id),
	CONSTRAINT uk_process_message UNIQUE (process_id, type),
	CONSTRAINT fk_process_message_process FOREIGN KEY (process_id) REFERENCES new.process(id)
);
-- nenhuma alteracao estrutural
CREATE TABLE new.process_message_history (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- change details
	field varchar(100) NOT NULL,
	previous_value varchar(255) NOT NULL,
	current_value varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	-- foreign key to process_message
	process_message_id int8 NOT NULL,
	-- constraints
	CONSTRAINT process_message_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_message_history_message FOREIGN KEY (process_message_id) REFERENCES new.process_message(id) ON DELETE CASCADE
);
-- mudanca 2
-- process_target_id vira process_id
-- process_origin_id vira process_reference_id
-- retirada targed_under_var e origin_over_var
-- origin_under_var e target_over_var viram under_var e over_var
-- a contraint unique passa a ser em name, process_id, process_reference_id
-- isto e´,
-- a logica de indicador é alterada para ser sempre em relacao a um processo
-- nao a mais a separacao de processo origem e destino
-- o indicador é sempre de um processo e tem como a referencia os valores de outro processo
-- pode-se ter mais de um indicador com nomes diferentes para o mesmo par de processos
CREATE TABLE new.process_indicator (
	-- id
	id bigserial NOT NULL,
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
	process_id int8 NOT NULL, -- antigo target
	process_reference_id int8 NOT NULL, -- antigo origin
	-- constraints
	CONSTRAINT process_indicator_pkey PRIMARY KEY (id),
	CONSTRAINT process_indicators_uk UNIQUE (name, process_id, process_reference_id),
	CONSTRAINT fk_pi_process_id FOREIGN KEY (process_id) REFERENCES new.process(id) ON DELETE CASCADE,
	CONSTRAINT fk_pi_process_reference FOREIGN KEY (process_reference_id) REFERENCES new.process(id) ON DELETE CASCADE
);
-- nenhuma alteracao estrutural
CREATE TABLE new.process_indicator_history (
	-- id
	id bigserial NOT NULL,
	-- control
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	-- change details
	field varchar(100) NOT NULL,
	previous_value varchar(255) NOT NULL,
	current_value varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	-- foreign key to process_indicator
	process_indicator_id int8 NOT NULL,
	-- constraints
	CONSTRAINT process_indicator_history_pkey PRIMARY KEY (id),
	CONSTRAINT fk_process_indicator_history_indicator FOREIGN KEY (process_indicator_id) REFERENCES new.process_indicator(id) ON DELETE CASCADE
);