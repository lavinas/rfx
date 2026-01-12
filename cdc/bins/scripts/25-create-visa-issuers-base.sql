SET search_path TO bins;

drop table if exists visa_issuers_bin;
CREATE TABLE visa_issuers_bin (
    bin VARCHAR(20) primary key,
    id VARCHAR(20) NOT NULL
);
create index visa_issuers_bin_idx on visa_issuers_bin(bin);


drop table visa_issuers_main;
CREATE TABLE visa_issuers_main (
    id VARCHAR(20),
    code VARCHAR(20),
    issuer VARCHAR(100)
);
create index visa_issuers_main_idx on visa_issuers_main(id);