DROP TABLE IF EXISTS reports.infrterm_ch;
CREATE TABLE IF NOT EXISTS reports.infrterm_ch (
    id BIGSERIAL PRIMARY KEY,
    sync_status SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    ano SMALLINT NOT NULL,
    trimestre SMALLINT NOT NULL,
    uf CHAR(2) NOT NULL,
    quantidade_total INTEGER NOT NULL,
    quantidade_pos_compartilhados INTEGER NOT NULL,
    quantidade_pos_leitora_chip INTEGER NOT NULL,
    quantidade_pdv INTEGER NOT NULL
);


INSERT INTO reports.infrterm_ch (
    ano,
    trimestre,
    uf,
    quantidade_total,
    quantidade_pos_compartilhados,
    quantidade_pos_leitora_chip,
    quantidade_pdv
)
select 2025,
       4,
       b.uf,
       count(1) quantidade_total,
       0 quantidade_pos_compartilhados,
       sum(case when a.tipo_terminal = 'POS' then 1 else 0 end) quantidade_pos_leitora_chip,
       sum(case when a.tipo_terminal = 'TEF' then 1 else 0 end) quantidade_pdv
  from apoio.terminais a
left join apoio.estabelecimentos b
  on a.codigo_estabelecimento = b.codigo_estabelecimento
 where b.codigo_estabelecimento is not null
group by 1,2,3;



select count(1)
  from reports.infrterm_ch
where quantidade_total != quantidade_pdv + quantidade_pos_leitora_chip + quantidade_pos_compartilhados;