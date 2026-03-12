create table reporter.report (
    id bigserial primary key,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    name varchar(200) not null,
    description text null
);







-- fechamento de periodo -> consolidacoes -> finalizacoes -> reports



-- replicacao dados brutos -> transacoes consolidadas -> consolidacao -> geracao do report