-- Active: 1766518799113@@127.0.0.1@5434@reg@transaction_v4
-- total management
-- 8849392
-- 8830493
select count(distinct key1)
  from raw_data_v2.intercambio_transaction a
 where a.dt_processamento >= '2026-01-01'
   and a.dt_processamento < '2026-04-01';

-- 8849392
-- 8830493
-- total management
select count(1) 
  from transaction_v4.transaction
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
  from transaction_v4.transaction
 where transaction_secondary_date >= '2026-01-01'
   and transaction_secondary_date < '2026-04-01'
   and status_id in (1, 2);


select status_id, status_name, count(1)
  from transaction_v4.transaction
 where transaction_date >= '2026-01-01'
   and transaction_date < '2026-04-01'
 group by 1, 2;

--
select count(1)
  from transaction_v4.transaction
where transaction_date >= '2026-01-01'
  and transaction_date < '2026-04-01'
  and status_id = 0;


create table tmp_intercam as
select *
 from transaction_v4.transaction
where transaction_date >= '2026-01-01'
  and transaction_date < '2026-04-01'
  and status_id = 0;


create table tmp_management as
select *
 from transaction_v4.transaction
where transaction_date >= '2026-01-01'
  and transaction_date < '2026-04-01'
  and status_id = 1;

select count(1)
  from tmp_intercam a
left join tmp_management b
       on a.key2 = b.key2
where b.key2 is null;


select a.transaction_date, b.transaction_date, a.bin, b.bin
  from tmp_intercam a
left join tmp_management b
  on b.transaction_amount = a.transaction_amount
  and COALESCE(b.bin, '') = COALESCE(a.bin, '')
  and b.authorization_code = a.authorization_code
  and b.establishment_code = a.establishment_code
where b.id is not null;


select count(1)
  from transaction_v4.transaction a
where bin is null;

select count(1)
  from transaction_v4.transaction a
where authorization_code is null;


select count(1)
  from transaction_v4.transaction a
where establishment_code is null;


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
--      and a.bin = b.bin
      and a.authorization_code = b.authorization_code
--      and a.establishment_code = b.establishment_code
--      and a.transaction_date >= b.transaction_date - interval '3 day'
--      and a.transaction_date <= b.transaction_date + interval '3 day'
where b.transaction_amount is not null;



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


select md5(a.transaction_amount::text || a.bin::text || a.authorization_code::text || a.establishment_code::text) as hash_intercam
  from transaction_v3.transaction as a;



update transaction_v3.transaction
   set key2 = md5(transaction_amount::text || bin::text || authorization_code::text || establishment_code::text);



commit;


select count(1)
  from transaction_v3.transaction a
inner join transaction_v3.transaction b
        on a.transaction_amount = b.transaction_amount
       and a.bin = b.bin
       and a.authorization_code = b.authorization_code
       and a.establishment_code = b.establishment_code
       and a.transaction_date >= b.transaction_date - interval '3 day'
       and a.transaction_date <= b.transaction_date + interval '3 day'
       and b.status_id = 1
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 0;


select count(1)
  from transaction_v3.transaction a
inner join transaction_v3.transaction b
        on a.key2 = b.key2
       and b.status_id = 1
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 0;


select count(1)
  from transaction_v3.transaction a
inner join transaction_v3.transaction b
        on b.key2 = a.key2
       and b.transaction_date >= a.transaction_date - interval '3 day'
       and b.transaction_date <= a.transaction_date + interval '3 day'
       and b.status_id = 1
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 0;


select count(1)
  from transaction_v3.transaction a
left join transaction_v3.transaction b
        on b.key2 = a.key2
       and b.transaction_date >= a.transaction_date - interval '3 day'
       and b.transaction_date <= a.transaction_date + interval '3 day'
       and b.status_id = 1
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 0
  and b.key2 is null;


  select count(1)
    from transaction_v4.transaction a
  where transaction_date >= '2026-01-01'
    and transaction_date < '2026-04-01'
    and status_id = 0
    and not exists (
        select 1
          from transaction_v4.transaction b
         where b.key2 = a.key2
           and b.transaction_date >= a.transaction_date - interval '3 day'
           and b.transaction_date <= a.transaction_date + interval '3 day'
           and b.status_id = 1
    );




    select count(1)
      from transaction_v4.transaction a
    where reference_id is not null;


select *
  from transaction_v4.transaction
 where reference_id is not null;



select count(1)
  from transaction_v4.transaction a
inner join transaction_v4.transaction b
        on b.reference_id = a.id
      and b.status_id = 3
where a.reference_id is not null
  and a.status_id = 2;





select count(1)
  from transaction_v4.transaction a
inner join transaction_v4.transaction b
        on b.reference_id = a.id
      and b.status_id = 3
where a.reference_id is not null
  and a.status_id = 2;


update transaction_v4.transaction a
   set reference_id = b.id
  from transaction_v4.transaction b
 where b.reference_id = a.id
   and b.status_id = 3
    and a.status_id = 2
    and a.reference_id is not null;


select count(1)
  from transaction_v4.transaction a
inner join transaction_v4.transaction b
        on b.id = a.reference_id
      and b.status_id = 3
where a.reference_id is not null
  and a.status_id = 2;


select count(1)
  from transaction_v4.transaction a
inner join transaction_v4.transaction b
        on b.id = a.reference_id
      and b.status_id = 2
where a.reference_id is not null
  and a.status_id = 3;



select count(1)
  from transaction_v4.transaction a
where transaction_date >= '2026-01-01'
  and transaction_date < '2026-04-01'
  and status_id = 0;


select *
  from transaction_v4.transaction
 where transaction_date >= '2026-01-01'
   and transaction_date < '2026-04-01'
   and status_id = 0;


select max (a.transaction_date - b.transaction_date), min (a.transaction_date - b.transaction_date)
  from transaction_v4.transaction a
 inner join transaction_v4.transaction b
        on b.id = a.reference_id
where a.transaction_date >= '2026-01-01'
  and a.transaction_date < '2026-04-01'
  and a.status_id = 2
  and a.reference_id is not null;
