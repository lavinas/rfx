-- create table segmentos_descricao
DROP TABLE IF EXISTS apoio.segmentos_descricao;
CREATE TABLE apoio.segmentos_descricao (
    id INTEGER,
    description TEXT
);

truncate table apoio.segmentos_descricao;
-- use 2-import_reports.py to import segmentos_descricao.csv into segmentos_descricao table

-- criar segmentos table
DROP TABLE IF EXISTS reports.segmentos_ch;
CREATE TABLE reports.segmentos_ch(
    id SERIAL NOT NULL,
    sync_status integer DEFAULT 0,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    nome_segmento varchar(50),
    descricao_segmento varchar(250),
    codigo_segmento integer,
    UNIQUE(codigo_segmento),
    UNIQUE(nome_segmento),
    PRIMARY KEY(id)
);


-- insert segmentos
truncate table reports.segmentos_ch;
insert into reports.segmentos_ch (codigo_segmento, nome_segmento, descricao_segmento)
select a.id codigo_segmento,
       a.description nome_segmento,
       concat('MCC: ', string_agg(distinct b.mcc::text, ', ')) AS descricao_segmento
  from apoio.segmentos_descricao a
inner join apoio.gestao b
    on a.id::varchar = b.segmento
group by 1, 2
order by 1;

