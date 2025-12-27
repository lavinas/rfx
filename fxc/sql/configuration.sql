-- Active: 1766004629598@@127.0.0.1@5433@flx


-- public.process - tabela principal que lista os processos que serao monitorados
-- alteracoes:
--   retirar default das classificacoes e description
--   criar indices para process_name
CREATE TABLE public.process (
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
CREATE INDEX idx_process_flow_id ON public.process USING btree (flow_id);
CREATE INDEX idx_process_source_id ON public.process USING btree (source_id);
CREATE INDEX idx_process_name ON public.process USING btree (process_name);

-- public.process_time_limit - tabela que define os limites de tempo para cada processo
-- alteracoes: 
--  adicionar index  process_id
CREATE TABLE public.process_time_limit (
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
	CONSTRAINT fk_process_time_limit_process FOREIGN KEY (process_id) REFERENCES public.process(id) ON DELETE CASCADE
);
CREATE INDEX idx_process_time_limit_process ON public.process_time_limit USING btree (process_id);

-- public.process_time_limit_history - tabela que registra o historico de alteracoes nos limites de tempo dos processos
-- alteracoes:
--  adicionar index process_time_limit_id e created_at
CREATE TABLE public.process_time_limit_history (
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
	CONSTRAINT fk_process_time_limit_history_time_limit FOREIGN KEY (process_time_limit_id) REFERENCES public.process_time_limit(id) ON DELETE CASCADE
);
CREATE INDEX idx_process_time_limit_history_process_time_limit ON public.process_time_limit_history USING btree (process_time_limit_id);
CREATE INDEX idx_process_time_limit_history_created_at ON public.process_time_limit_history USING btree (created_at);

-- public.process_error - tabela que define os possiveis erros associados aos processos
-- alteracoes: nome da constraint uk_process_error_error_key corrigido
CREATE TABLE public.process_error (
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
	CONSTRAINT fk_process_error_process FOREIGN KEY (process_id) REFERENCES public.process(id)
);
CREATE INDEX idx_process_error_process ON public.process_error USING btree (process_id);

-- public.process_error_history - tabela que registra o historico de alteracoes nos erros dos processos
-- alteracoes: 
--   adicionar index process_error_id e created_at
CREATE TABLE public.process_error_history (
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
	CONSTRAINT fk_process_error_history_error FOREIGN KEY (process_error_id) REFERENCES public.process_error(id) ON DELETE CASCADE
);
CREATE INDEX idx_process_error_history_process_error ON public.process_error_history USING btree (process_error_id);
CREATE INDEX idx_process_error_history_created_at ON public.process_error_history USING btree (created_at);

-- public.process_message - tabela que define as mensagens de erro e timeout para cada processo
-- alteracoes:
--  campo type adicionado para categorizar a mensagem
--  os campos subject e body foram adicionados para armazenar o conteudo da mensagem em substituição aos campos individuais
--  alteracoes nos nomes das constraints para seguir o padrao
--  adicionar constraint unica para process_id e type
--  criar index para process_id
CREATE TABLE public.process_message (
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
	CONSTRAINT fk_process_message_process FOREIGN KEY (process_id) REFERENCES public.process(id)
);
CREATE INDEX idx_process_message_process ON public.process_message USING btree (process_id);

-- public.process_message_history - tabela que registra o historico de alteracoes nas mensagens dos processos
-- alteracoes:
--  adicionar index process_message_id e created_at
CREATE TABLE public.process_message_history (
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
	CONSTRAINT fk_process_message_history_message FOREIGN KEY (process_message_id) REFERENCES public.process_message(id) ON DELETE CASCADE
);
CREATE INDEX idx_process_message_history_process_message ON public.process_message_history USING btree (process_message_id);
CREATE INDEX idx_process_message_history_created_at ON public.process_message_history USING btree (created_at);

-- public.process_indicator - tabela que define os indicadores de monitoramento de um processo
-- alteracoes:
-- deixar nx1 com a tabela de processos - ou seja, cada processo tem n indicadores, mas cada indicador pertence a um unico processo
--  adicionar campo description
-- somente campos under_var e over_var em relacao ao processo de comparacao
-- process_id, process_reference_id ao inves de process_origin_id, process_target_id
CREATE TABLE public.process_indicator (
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
	CONSTRAINT fk_pi_process_id FOREIGN KEY (process_id) REFERENCES public.process(id) ON DELETE CASCADE,
	CONSTRAINT fk_pi_process_reference FOREIGN KEY (process_reference_id) REFERENCES public.process(id) ON DELETE CASCADE
);
CREATE INDEX idx_process_indicator_process ON public.process_indicator USING btree (process_id);
CREATE INDEX idx_process_indicator_process_reference ON public.process_indicator USING btree (process_reference_id);

-- public.process_indicator_history - tabela que registra o historico de alteracoes nos indicadores dos processos
-- alteracoes:
--  adicionar index process_indicator_id e created_at
CREATE TABLE public.process_indicator_history (
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
	CONSTRAINT fk_process_indicator_history_indicator FOREIGN KEY (process_indicator_id) REFERENCES public.process_indicator(id) ON DELETE CASCADE
);
CREATE INDEX idx_process_indicator_history_process_indicator ON public.process_indicator_history USING btree (process_indicator_id);
CREATE INDEX idx_process_indicator_history_created_at ON public.process_indicator_history USING btree (created_at);


