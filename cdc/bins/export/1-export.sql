select CONCAT('INSERT INTO cadoc_6334.bins (bin, product_code, card_type) VALUES (''', bin::text, ''', ', produto_final::text, ', ''', modalidade_final::text, ''');') as bin_list
  from bins.bins
limit 1000000;



select * from bins.bins;


select count(1) from bins.bins;

select count(1) from cadoc_6334.bins;