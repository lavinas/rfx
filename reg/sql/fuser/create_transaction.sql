CREATE TABLE transaction(
    id BIGSERIAL NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    key1 varchar(50) NOT NULL,
    key2 varchar(50),
    key3 varchar(50),
    establishment_code numeric,
    establishment_nature numeric,
    establishment_mcc numeric,
    establishment_terminal_code numeric,
    bin numeric,
    authorization_code varchar(20),
    transaction_nsu varchar(20),
    transaction_date date,
    transaction_secondary_date date,
    transaction_amount numeric(15,2),
    transaction_secondary_amount numeric(15,2),
    transaction_installments numeric,
    transaction_installments_type varchar(10),
    transaction_brand varchar(2),
    transaction_product varchar(2),
    transaction_capture varchar(3),
    revenue_mdr_value numeric(15,2),
    cost_interchange_value numeric(15,2),
    high_source_priority integer,
    status_id integer,
    status_name varchar(20),
    status_count integer DEFAULT 3,
    period_date date,
    period_closing_id bigint,
    transac_id varchar(50),
    reference_id bigint,
    PRIMARY KEY(id),
    CONSTRAINT fk_reference_id FOREIGN KEY (reference_id) REFERENCES transaction_v4.transaction(id)
);
CREATE UNIQUE INDEX idx_transaction_key1 ON transaction_v4.transaction USING btree (key1);
CREATE INDEX idx_transaction_key2 ON transaction_v4.transaction USING btree (key2);
CREATE INDEX idx_transaction_secondary_transaction_date ON transaction_v4.transaction USING btree (transaction_secondary_date);
CREATE INDEX idx_transaction_transaction_date ON transaction_v4.transaction USING btree (transaction_date);
CREATE INDEX transaction_transaction_date_status_id_idx ON transaction_v4.transaction USING btree (transaction_date, status_id);
CREATE INDEX idx_transaction_reference_id ON transaction_v4.transaction USING btree (reference_id);


-- create table transaction_v4.transaction_reference
CREATE TABLE transaction_reference (
    transaction_id BIGSERIAL NOT NULL,
    reference_type VARCHAR(50) NOT NULL,
    reference_id VARCHAR(50) NOT NULL,
    PRIMARY KEY(transaction_id, reference_type, reference_id)
);

CREATE INDEX idx_transaction_reference_transaction_id ON transaction_reference USING btree (transaction_id);