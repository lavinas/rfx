SET search_path TO bins;

drop table if exists bins.visa_bin_ardef;
CREATE TABLE bins.visa_bin_ardef (
    id bigint NOT NULL,
    created_date TIMESTAMP,
    effective_date TIMESTAMP NULL,
    f08_issuer_identifier VARCHAR(20),
    f10_account_number_length INT,
    f14_domain VARCHAR(2),
    f15_region VARCHAR(5),
    f16_country VARCHAR(5),
    f29_product_id VARCHAR(10),
    f30_combo_card VARCHAR(5),
    f35_account_funding_source VARCHAR(5),
    f37_travel_account_data VARCHAR(20),
    f40_product_subtype VARCHAR(10),
    key_range_high INT,
    key_range_low INT
);

create index idx_visa_bin_ardef_on_key_range_low
    on bins.visa_bin_ardef (key_range_low, key_range_high); 