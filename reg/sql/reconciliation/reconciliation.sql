
CREATE schema reconciliation;

-- sources table to register the different sources of transactions for reconciliation
CREATE TABLE reconciliation.sources (
    id bigserial PRIMARY KEY,
    name varchar(200) NOT NULL
);

CREATE TABLE reconciliation.dash (
    -- control
    id bigserial PRIMARY KEY,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    -- data
    name varchar(200) not null,
    master_id int not null, -- referencia reconciliation.sources(id)
    slave_id int not null, -- referencia reconciliation.sources(id)
    -- constraints
    constraint fk_source_master foreign key (master_id) references reconciliation.sources(id),
    constraint fk_source_slave foreign key (slave_id) references reconciliation.sources(id)
);

CREATE TABLE reconciliation.dash_itens (
    -- control
    id bigserial PRIMARY KEY,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    dash_id int not null, -- referencia reconciliation.dash(id)
    -- data
    master_value numeric(18,2) not null,
    slave_value numeric(18,2) not null,
    status_id int not null, -- 1 - conciliado, 2 - slave missing, 3 - master missing, 4 - discrepancy
    status_description varchar(20) not null, -- 'Conciliado', 'Não encontrado no <slave source>', 'Não encontrado no <master source>', 'Valor divergente'
    constraint fk_dash foreign key (dash_id) references reconciliation.dash(id)
);

