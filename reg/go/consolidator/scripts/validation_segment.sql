

drop table if exists cadoc_6334_v2.tmp_segment_group;

create table cadoc_6334_v2.tmp_segment_group as
select extract(year from a.transaction_date) as year,
       extract(quarter from a.transaction_date) as quarter,
       coalesce(b.segment,  423) as segment,
       a.establishment_mcc mcc
  from transaction_v4.transaction a
left join apoio.segmentos b
  on a.establishment_mcc >= b.mcc_init
  and a.establishment_mcc <= b.mcc_end
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-01-06'
  and a.status_id = 2
group by 1, 2, 3, 4
order by 1, 2, 3, 4


select * from cadoc_6334_v2.tmp_segment_group where segment = 413;

drop table if exists cadoc_6334_v2.tmp_segmento;

create table cadoc_6334_v2.tmp_segmento as
select year,
       quarter,
       b.description segment_name,
       concat('MCC: ', string_agg(distinct a.mcc::text, ', ' ORDER BY a.mcc::text)) AS segment_description,
       b.id segment_code
  from cadoc_6334_v2.tmp_segment_group a
left join apoio.segmentos_descricao b
    on a.segment = b.id
group by 1, 2, 5, 3;



select count(1) from cadoc_6334_v2.tmp_segmento;


select count(1) from cadoc_6334_v2.segmento


select count(1) from cadoc_6334_v2.segmento a
left join cadoc_6334_v2.tmp_segmento b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.segment_code = b.segment_code
  and a.segment_name = b.segment_name
where b.segment_code is null;


select count(1) from cadoc_6334_v2.tmp_segmento a
left join cadoc_6334_v2.segmento b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.segment_code = b.segment_code
  and a.segment_name = b.segment_name
where b.segment_code is null;


select a.segment_code, a.segment_name, a.segment_description, b.segment_description
  from cadoc_6334_v2.segmento a
left join cadoc_6334_v2.tmp_segmento b
  on a.year = b.year
  and a.quarter = b.quarter
  and a.segment_code = b.segment_code
  and a.segment_name = b.segment_name
where a.segment_description != b.segment_description
