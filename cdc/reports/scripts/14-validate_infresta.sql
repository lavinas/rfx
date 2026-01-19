select * from reports.infresta;

update reports.infresta set trimestre = 4;


drop table if exists reports.infresta_ch;
CREATE TABLE IF NOT EXISTS reports.infresta_ch (
    id BIGINT NULL,
    sync_status SMALLINT NULL,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NULL,
    ano SMALLINT NOT NULL,
    trimestre SMALLINT NOT NULL,
    uf CHAR(2) NOT NULL,
    quantidade_estabelecimentos_totais INTEGER NOT NULL,
    quantidade_estabelecimentos_captura_manual INTEGER NOT NULL,
    quantidade_estabelecimentos_captura_eletronica INTEGER NOT NULL,
    quantidade_estabelecimentos_captura_remota INTEGER NOT NULL
);


INSERT INTO reports.infresta_ch (ano, 
                                 trimestre, 
                                 uf,
                                 quantidade_estabelecimentos_totais, 
                                 quantidade_estabelecimentos_captura_manual, 
                                 quantidade_estabelecimentos_captura_eletronica,
                                 quantidade_estabelecimentos_captura_remota)
select 2025,
         4,
         uf,
         count(1),
         0,
         count(1),
         0
      from apoio.estabelecimentos
    where data_credenciamento < '2026-01-01'
      and not (tem_visa = false and tem_mastercard = false and tem_elo = false)
      and not (tem_debito = false and tem_credito = false)
    group by uf;
