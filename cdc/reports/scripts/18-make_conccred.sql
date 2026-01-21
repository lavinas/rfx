drop table if exists reports.conccred_ch;
CREATE TABLE IF NOT EXISTS reports.conccred_ch (
    id BIGSERIAL PRIMARY KEY,
    sync_status SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    ano SMALLINT NOT NULL,
    trimestre SMALLINT NOT NULL,
    bandeira SMALLINT NOT NULL,
    funcao CHAR(1) NOT NULL,
    quantidade_estabelecimentos_credenciados INTEGER NOT NULL,
    quantidade_estabelecimentos_ativos INTEGER NOT NULL,
    valor_transacoes NUMERIC(18,2) NOT NULL,
    quantidade_transacoes INTEGER NOT NULL
);


insert into reports.conccred_ch (ano, trimestre, bandeira, funcao, valor_transacoes, quantidade_transacoes, quantidade_estabelecimentos_credenciados, quantidade_estabelecimentos_ativos)
select 2025, 4, 1, 'C', 0, 0, 
       count(1),
       sum(case when data_ultima_transacao >= '2026-01-01'::date - INTERVAL '180 days' then 1 else 0 end)
  from apoio.estabelecimentos
where data_credenciamento < '2026-01-01'
  and (tem_visa = true and tem_credito = true);
insert into reports.conccred_ch (ano, trimestre, bandeira, funcao, valor_transacoes, quantidade_transacoes, quantidade_estabelecimentos_credenciados, quantidade_estabelecimentos_ativos)
select 2025, 4, 1, 'D', 0, 0, 
       count(1),
       sum(case when data_ultima_transacao >= '2026-01-01'::date - INTERVAL '180 days' then 1 else 0 end)
  from apoio.estabelecimentos
where data_credenciamento < '2026-01-01'
  and (tem_visa = true and tem_debito = true);
insert into reports.conccred_ch (ano, trimestre, bandeira, funcao, valor_transacoes, quantidade_transacoes, quantidade_estabelecimentos_credenciados, quantidade_estabelecimentos_ativos)
select 2025, 4, 2, 'C', 0, 0, 
       count(1),
       sum(case when data_ultima_transacao >= '2026-01-01'::date - INTERVAL '180 days' then 1 else 0 end)
  from apoio.estabelecimentos
where data_credenciamento < '2026-01-01'
  and (tem_mastercard = true and tem_credito = true);
insert into reports.conccred_ch (ano, trimestre, bandeira, funcao, valor_transacoes, quantidade_transacoes, quantidade_estabelecimentos_credenciados, quantidade_estabelecimentos_ativos)
select 2025, 4, 2, 'D', 0, 0, 
       count(1),
       sum(case when data_ultima_transacao >= '2026-01-01'::date - INTERVAL '180 days' then 1 else 0 end)
  from apoio.estabelecimentos
where data_credenciamento < '2026-01-01'
  and (tem_mastercard = true and tem_debito = true);
insert into reports.conccred_ch (ano, trimestre, bandeira, funcao, valor_transacoes, quantidade_transacoes, quantidade_estabelecimentos_credenciados, quantidade_estabelecimentos_ativos)
select 2025, 4, 8, 'C', 0, 0, 
       count(1),
       sum(case when data_ultima_transacao >= '2026-01-01'::date - INTERVAL '180 days' then 1 else 0 end)
  from apoio.estabelecimentos
where data_credenciamento < '2026-01-01'
  and (tem_elo = true and tem_credito = true);
insert into reports.conccred_ch (ano, trimestre, bandeira, funcao, valor_transacoes, quantidade_transacoes, quantidade_estabelecimentos_credenciados, quantidade_estabelecimentos_ativos)
select 2025, 4, 8, 'D', 0, 0, 
       count(1),
       sum(case when data_ultima_transacao >= '2026-01-01'::date - INTERVAL '180 days' then 1 else 0 end)
  from apoio.estabelecimentos
where data_credenciamento < '2026-01-01'
  and (tem_elo = true and tem_debito = true);
update reports.conccred_ch concred
   set valor_transacoes = (select sum(valor_transacoes)
                             from reports.descontos_ch
                            where concred.bandeira = bandeira
                              and concred.funcao = funcao),
         quantidade_transacoes = (select sum(quantidade_transacoes)
                                   from reports.descontos_ch
                                  where concred.bandeira = bandeira
                                    and concred.funcao = funcao);
