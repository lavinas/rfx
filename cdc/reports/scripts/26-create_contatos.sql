drop table if exists reports.contatos_ch;
CREATE TABLE IF NOT EXISTS reports.contatos_ch (
    id bigserial PRIMARY KEY,
    Ano integer,
    Trimestre integer,
    TipoContato varchar(1),
    Nome varchar(50),
    Cargo varchar(50),
    NumeroTelefone varchar(50),
    Email varchar(50)
);

INSERT INTO reports.contatos_ch (Ano, Trimestre, TipoContato, Nome, Cargo, NumeroTelefone, Email)
VALUES (2025, 4, 'D', 'Diego Silveira de Faria', 'Diretor', '(65) 99683-4775', 'diego@redeflex.com.br');

INSERT INTO reports.contatos_ch (Ano, Trimestre, TipoContato, Nome, Cargo, NumeroTelefone, Email)
VALUES (2025, 4, 'T', 'Werika Oliveira Calassa', 'Gerente de Controladoria', '(65) 99964-2769', 'werika.calassa@redeflex.com.br');

INSERT INTO reports.contatos_ch (Ano, Trimestre, TipoContato, Nome, Cargo, NumeroTelefone, Email)
VALUES (2025, 4, 'T', 'Andrea Da Silva Santos', 'Coordenador Fiscal', '(65) 99901-7193', 'andrea.dasantos@redeflex.com.br');

INSERT INTO reports.contatos_ch (Ano, Trimestre, TipoContato, Nome, Cargo, NumeroTelefone, Email)
VALUES (2025, 4, 'I', '', '', '', 'atendimento.bacen@redeflex.com.br');