select count(1)
  from reports.descontos_ch a
left join reports.segmentos_ch b
    on a.codigo_segmento = b.codigo_segmento
where b.id is null;


select count(1)
  from reports.intercam_ch a
left join reports.segmentos_ch b
    on a.codigo_segmento = b.codigo_segmento
where b.id is null;


select count(1)
  from reports.intercam_ch a
left join reports.ranking_ch b
    on a.codigo_segmento = b.codigo_segmento
where b.id is null;


select codigo_segmento,
       count(1)
  from reports.segmentos_ch
group by 1
having count(1) > 1;


select * from reports.segmentos_ch
where codigo_segmento is null;


select * from reports.segmentos_ch;