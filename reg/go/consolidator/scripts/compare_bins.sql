select count(1)
  from bins_2025_4t.bins a
left join bins.bins b
  on a.bin = b.bin
where b.bin is null

select count(1)
  from bins.bins a
left join bins_2025_4t.bins b
  on a.bin = b.bin
where b.bin is null



select count(1)
  from transaction_v4.transaction a
left join bins.bins b
  on b.bin = a.bin::text
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2
  and b.bin is null;


select a.bin, a.pais, b.pais, a.produto_intermediario, b.produto_intermediario, a.produto_final, b.produto_final
  from bins.bins a
inner join bins_2025_4t.bins b
  on a.bin = b.bin
where a.produto_final != b.produto_final;


select a.bin, a.pais, b.pais, a. a.produto_intermediario, b.produto_intermediario, a.produto_final, b.produto_final, count(1)
  from bins.bins a
inner join bins_2025_4t.bins b
  on a.bin = b.bin
where a.produto_final != b.produto_final
group by 1, 2, 3, 4, 5, 6, 7
ORDER BY 8 DESC;


select a.pais, b.pais, a.bandeira, b.bandeira, a.produto_base, b.produto_base, a.produto_intermediario, b.produto_intermediario, a.produto_final, b.produto_final, count(1)
  from bins.bins_v2 a
inner join bins_2025_4t.bins b
  on a.bin = b.bin
where a.produto_final != b.produto_final
group by 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
ORDER BY 11 DESC
limit 1000000;


create table bins.tmp_diff as
select a.bin, a.produto_final new, b.produto_final old
  from bins.bins_v2 a
inner join bins_2025_4t.bins b
  on a.bin = b.bin
where a.produto_final != b.produto_final

select count(1)
  from bins.tmp_diff


select count(1)
  from transaction_v4.transaction a
inner join bins.tmp_diff b
  on b.bin = a.bin::text
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2;


select count(1)
  from transaction_v4.transaction a
inner join bins.bins_v2 b
  on b.bin = a.bin::text
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2
  and a.transaction_brand = 'V'
  and b.produto_final = 38


select a.transaction_brand, b.produto_final, count(1), sum(a.transaction_amount)
  from transaction_v4.transaction a
inner join bins_2025_4t.bins b
  on b.bin = a.bin::text
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2
group by 1, 2
order by 1, 2;


select a.transaction_brand, b.produto_final, count(1), sum(a.transaction_amount)
  from transaction_v4.transaction a
inner join bins.bins_v2 b
  on b.bin = a.bin::text
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2
group by 1, 2
order by 1, 2;


select *
  from bins_2025_4t.bins
where produto_intermediario  like 'Propr%'


select count(1), sum(a.transaction_amount)
  from transaction_v4.transaction a
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2
  and a.bin in (
  51802286,
  51552855,
  52141935,
  51550173,
  51802216,
  51808519
  );


select count(1), sum(a.transaction_amount)
  from transaction_v4.transaction a
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2
  and a.bin in (
  51802286,
  51552855,
  52141935,
  51550173,
  51802216,
  51808519
  );

select count(1), sum(a.transaction_amount)
  from transaction_v4.transaction a
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2;



drop table bins.tmp_diff;


select count(1)
  from transaction_v4.transaction a
inner join bins.bins_v2 b
  on b.bin = a.bin::text
inner join bins_2025_4t.bins c
  on c.bin = a.bin::text
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2
  and a.transaction_brand = 'M'
  and b.produto_final = 32
  and c.produto_final != 32;



select b.bin, b.pais, b.produto_intermediario, b.produto_final, count(1)
  from transaction_v4.transaction a
inner join bins.bins_v2 b
  on b.bin = a.bin::text
left join bins_2025_4t.bins c
  on c.bin = a.bin::text
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2
  and a.transaction_brand = 'M'
  and b.produto_final = 32
  and c.bin is null
  group  by 1, 2, 3, 4
  order by 5 desc;
