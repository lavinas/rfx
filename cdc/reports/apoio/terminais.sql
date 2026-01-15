 SELECT 
                t18.id_terminal as codigo_terminal,
                t18.cd_pessoa_estabelecimento as codigo_estabelecimento,
                t22.de_tp_terminal as tipo_terminal,
                null as cadoc_item_id
            FROM efc.tb05018 t18
            INNER JOIN efc.tb05022 t22 ON t22.cd_tp_terminal = t18.cd_tp_terminal
            WHERE t18.cd_status = 1;
-- public.cadoc_6334_terminal definição
-- Drop table
-- DROP TABLE public.cadoc_6334_terminal;
CREATE TABLE apoio.cadoc_6334_terminal (
 codigo_terminal int8 NOT NULL,
 codigo_estabelecimento int8 NULL,
 tipo_terminal varchar(255) NULL,
 cadoc_item_id int8 NULL,
 CONSTRAINT terminal_pkey PRIMARY KEY (codigo_terminal)
);
 
-- terminais
select codigo_terminal,
       codigo_estabelecimento,
       tipo_terminal
   from apoio.cadoc_6334_terminal;