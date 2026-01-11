SET search_path TO bins;
CREATE TABLE elo_bin (
    id BIGINT PRIMARY KEY,
    bin_number INTEGER NOT NULL,
    bin_range_high INTEGER NOT NULL,
    bin_range_low INTEGER NOT NULL,
    created_date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    f01_c001_register_type_identifier CHAR(2) NOT NULL,
    f02_c002_bin_number VARCHAR(32) NOT NULL,
    f03_c003_bank_id VARCHAR(32) NOT NULL,
    f04_c004_allow_debt CHAR(1) NOT NULL,
    f05_c005_allow_credit CHAR(1) NOT NULL,
    f06_c006_bin_range CHAR(1) NOT NULL,
    f07_c007_card_type VARCHAR(64) NOT NULL,
    f08_c008_attribute_indicator CHAR(1) NOT NULL,
    f09_c009_personal_company_indicator CHAR(1) NOT NULL,
    f10_c010_reciever_bank_debt VARCHAR(32) NOT NULL,
    f11_c011_receiver_bank_credit VARCHAR(32) NOT NULL,
    f12_c012_capture_code VARCHAR(64) NOT NULL,
    f13_c013_product_code VARCHAR(64) NOT NULL,
    f14_c014_card_type_code VARCHAR(32) NOT NULL,
    f15_c015_bin_product_type CHAR(1) NOT NULL,
    f16_c018_bin_length INTEGER NOT NULL,
    f17_c019_bin_range_start_concat VARCHAR(32) NOT NULL,
    f18_c020_bin_range_end_concat VARCHAR(32) NOT NULL,
    f19_c021_early_settlement CHAR(1) NOT NULL,
    f21_c016_elo_reserved CHAR(1) NOT NULL,
    line_raw TEXT NOT NULL,
    updated_date TIMESTAMP WITHOUT TIME ZONE
);


CREATE TABLE elo_issuer (
    created_date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    f01_c001_register_type_identifier CHAR(2) NOT NULL,
    f02_c002_bank_id CHAR(4) NOT NULL,
    f03_c003_cnpj CHAR(14) NOT NULL,
    f04_c005_bank_name VARCHAR(128) NOT NULL,
    f05_c006_receiver_bank_code TEXT NOT NULL,
    f06_c007_product_code TEXT NOT NULL,
    f07_c008_reserved1 CHAR(24) NOT NULL,
    f08_c009_reserved2 TEXT NOT NULL,
    f09_c010_reserved3 TEXT NOT NULL,
    line_raw TEXT NOT NULL,
    updated_date TIMESTAMP WITHOUT TIME ZONE
);
