-- create descontos_up table
CREATE TABLE IF NOT EXISTS reports.descontos_up (
    id BIGINT NULL,
    sync_status SMALLINT NULL,
    created_at TIMESTAMP(3) WITHOUT TIME ZONE NULL,
    updated_at TIMESTAMP(3) WITHOUT TIME ZONE NULL,
    ano SMALLINT NULL,
    trimestre SMALLINT NULL,
    funcao CHAR(1) NOT NULL,
    bandeira SMALLINT NOT NULL,
    forma_captura SMALLINT NOT NULL,
    numero_parcelas SMALLINT NOT NULL,
    codigo_segmento INTEGER NOT NULL,
    taxa_desconto_media NUMERIC(5,2) NULL,
    taxa_desconto_minima NUMERIC(5,2) NULL,
    taxa_desconto_maxima NUMERIC(5,2) NULL,
    desvio_padrao_taxa_desconto NUMERIC(6,3) NULL,
    valor_transacoes NUMERIC(18,2) NOT NULL,
    quantidade_transacoes INTEGER NOT NULL
);
-- populate descontos_up table from intercam_up
truncate TABLE reports.descontos_up;
insert into reports.descontos_up (
    funcao,
    bandeira,
    forma_captura,
    numero_parcelas,
    codigo_segmento,
    valor_transacoes,
    quantidade_transacoes
)
select funcao,
       bandeira,
       forma_captura,
       numero_parcelas,
       codigo_segmento,
       sum(valor_transacoes) as valor_transacoes,
       sum(quantidade_transacoes) as quantidade_transacoes
    from reports.intercam_up
group by 1,2,3,4,5;

-- validate descontos_up duplicated
select funcao,
       bandeira,
       forma_captura,
       numero_parcelas,
       codigo_segmento,
       count(1)
  from reports.descontos_up
    group by 1,2,3,4,5
having count(1) > 1;
-- validate descontos_duplicated
select funcao,
       bandeira,
       forma_captura,
       numero_parcelas,
       codigo_segmento,
       count(1)
  from reports.descontos
    group by 1,2,3,4,5
having count(1) > 1;

-- count missing
select count(1)
  from reports.descontos_up a
left join reports.descontos b
    on a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento
 where b.id is null; 

-- count missing
select count(1)
  from reports.descontos a
left join reports.descontos_up b
    on a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento
 where b.funcao is null ; 


select * from reports.descontos_up a

-- update descontos_up with ano and trimestre
update reports.descontos_up a 
   set id = b.id,
       sync_status = b.sync_status,
       created_at = b.created_at,
       updated_at = b.updated_at,
       ano = b.ano,
       trimestre = b.trimestre,
       taxa_desconto_media = b.taxa_desconto_media,
       taxa_desconto_minima = b.taxa_desconto_minima,
       taxa_desconto_maxima = b.taxa_desconto_maxima,
       desvio_padrao_taxa_desconto = b.desvio_padrao_taxa_desconto
  from reports.descontos b
 where a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.codigo_segmento = b.codigo_segmento
   and a.id is null;

update reports.descontos_up a 
   set id = b.id,
       sync_status = b.sync_status,
       created_at = b.created_at,
       updated_at = b.updated_at,
       ano = b.ano,
       trimestre = b.trimestre,
       taxa_desconto_media = b.taxa_desconto_media,
       taxa_desconto_minima = b.taxa_desconto_minima,
       taxa_desconto_maxima = b.taxa_desconto_maxima,
       desvio_padrao_taxa_desconto = b.desvio_padrao_taxa_desconto
  from (
        select funcao,
               bandeira,
               forma_captura,
               numero_parcelas,
               codigo_segmento,
               -1 * min(id) as id,
               min(sync_status) as sync_status,
               min(created_at) as created_at,
               min(updated_at) as updated_at,
               max(ano) as ano,
               max(trimestre) as trimestre,
               sum(taxa_desconto_media * valor_transacoes) / sum(valor_transacoes) as taxa_desconto_media,
               min(taxa_desconto_minima) as taxa_desconto_minima,
               max(taxa_desconto_maxima) as taxa_desconto_maxima,
               sum(desvio_padrao_taxa_desconto * valor_transacoes) / sum(valor_transacoes) as desvio_padrao_taxa_desconto
          from reports.descontos
         group by funcao,
                  bandeira,
                  forma_captura,
                  numero_parcelas,
                  codigo_segmento
       ) b
 where a.funcao = b.funcao
   and a.bandeira = b.bandeira
   and a.forma_captura = b.forma_captura
   and a.numero_parcelas = b.numero_parcelas
   and a.id is null;


   select count(1)
     from reports.descontos_up a
    where a.id is null;

  select *
    from reports.descontos_up
  where id < 0;


 select a.*
     from reports.descontos_up a
    where a.id is null;


-- 1297
 select count(1)
     from reports.descontos_up a;

select round(sum(valor_transacoes * taxa_desconto_media / 100), 2)/sum(valor_transacoes) as valor_descontado
  from reports.descontos_up;


select  case bandeira 
         when 1 then 'Visa'
         when 2 then 'Master'
         when 8 then 'Elo'
         else 'Erro'
       end as bandeira, 
        round(sum(valor_transacoes * taxa_desconto_media / 100), 2)/sum(valor_transacoes) as valor_descontado
  from reports.descontos_up
group by 1
order by 1;