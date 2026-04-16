-- Active: 1774368236280@@192.168.100.78@5436@dev_regulat@transaction_v3
-- total management
-- 8849392
select count(distinct key1)
  from raw_data_v2.intercambio_transaction a
 where a.dt_processamento >= '2026-01-01'
   and a.dt_processamento < '2026-04-01';

-- 8849392
-- total management
select count(1) 
  from transaction_v3.transaction
 where transaction_date >= '2026-01-01'
   and transaction_date < '2026-04-01'
   and status_id in (0, 2);

-- 8855649
select count(distinct COALESCE(a.key1,  gen_random_uuid()::text))
  from raw_data_v2.management_transaction a
 where a.dt_processamento >= '2026-01-01'
   and a.dt_processamento < '2026-04-01';

-- 8855649
select count(1) 
  from transaction_v3.transaction
 where transaction_secondary_date >= '2026-01-01'
   and transaction_secondary_date < '2026-04-01'
   and status_id in (1, 2);



--
select count(1)
  from transaction_v3.transaction
where transaction_date >= '2026-01-01'
  and transaction_date < '2026-04-01'
  and status_id = 0;


create table tmp_intercam as
select *
 from transaction_v3.transaction
where transaction_date >= '2026-01-01'
  and transaction_date < '2026-04-01'
  and status_id = 0;


create table tmp_management as
select *
 from transaction_v3.transaction
where transaction_date >= '2026-01-01'
  and transaction_date < '2026-04-01'
  and status_id = 1;



select count(1)
  from tmp_intercam a
left join tmp_management b
       on a.key2 = b.key2
where b.key2 is null;

select count(1)
  from tmp_intercam a
left join tmp_management b
       on a.key2 = b.key2
where b.key2 is null;

select count(1)
  from tmp_intercam a
left join tmp_management b
       on a.key2 = b.key2
where b.key2 is null;


select * from tmp_intercam;

select * from tmp_management;



valor + bin + autorizacao


select count(1)
  from tmp_intercam a 
left join tmp_management b
       on a.transaction_amount = b.transaction_amount
      and a.bin = b.bin
      and a.authorization_code = b.authorization_code
      and a.establishment_code = b.establishment_code
      and a.transaction_date >= b.transaction_date - interval '3 day'
      and a.transaction_date <= b.transaction_date + interval '3 day'
where b.transaction_amount is null;



select a.id, count(1)
  from tmp_intercam a 
inner join tmp_management b
       on a.transaction_amount = b.transaction_amount
      and a.bin = b.bin
      and a.authorization_code = b.authorization_code
      and a.establishment_code = b.establishment_code
      and a.transaction_date >= b.transaction_date - interval '3 day'
      and a.transaction_date <= b.transaction_date + interval '3 day'
group by a.id
having count(1) > 1;


select a.id, b.id
  from tmp_intercam a 
inner join tmp_management b
       on a.transaction_amount = b.transaction_amount
      and a.bin = b.bin
      and a.authorization_code = b.authorization_code
      and a.establishment_code = b.establishment_code
      and a.transaction_date >= b.transaction_date - interval '3 day'
      and a.transaction_date <= b.transaction_date + interval '3 day'



select * from transaction_v3.transaction
where id in (2266379, 2487453);

2067925	2298810

MG_59894723

select * from raw_data_v2.management_transaction where cd_transacao_fin = '59894723';

select * from raw_data_v2.intercambio_transaction where key1 = 'V466002514410518';



select count(1)
  from raw_data_v2.management_transaction a
where key1 is null or key1 = '';