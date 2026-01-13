import argparse
import csv
from pathlib import Path
import psycopg2
from psycopg2 import sql
from psycopg2.extras import execute_values

#!/usr/bin/env python3

dsn_default = "postgresql://root:root@localhost:5435/cdc"
schema_default = "reports"
table_default = "reports_tables"

def insert_batch(cursor, schema_name, table_name, columns, rows):
    table_identifier = (
        sql.SQL(".").join([sql.Identifier(schema_name), sql.Identifier(table_name)])
        if schema_name
        else sql.Identifier(table_name)
    )
    query = sql.SQL("INSERT INTO {table} ({fields}) VALUES %s").format(
        table=table_identifier,
        fields=sql.SQL(",").join(sql.Identifier(col) for col in columns),
    )
    try:
        execute_values(cursor, query, rows)
    except psycopg2.Error as exc:
        sample_row = rows[0] if rows else None
        raise RuntimeError(
            f"Erro ao inserir lote na tabela {schema_name}.{table_name} "
            f"(tamanho do lote: {len(rows)}, primeira linha: {sample_row})"
        ) from exc


def load_csv_to_table(dsn, schema_name, table_name, csv_path, batch_size=1000):
    path = Path(csv_path)
    if not path.is_file():
        raise FileNotFoundError(f"Arquivo CSV não encontrado: {path}")

    with path.open(newline="", encoding="utf-8") as csvfile:
        reader = csv.reader(csvfile)
        try:
            headers = next(reader)
        except StopIteration:
            raise ValueError("CSV vazio, sem cabeçalho")

        columns = [header.strip() for header in headers]
        if not all(columns):
            raise ValueError("Cabeçalho contém nomes de coluna vazios")

        with psycopg2.connect(dsn) as connection:
            with connection.cursor() as cursor:
                batch = []
                for row in reader:
                    if not any(row):
                        continue
                    if len(row) != len(columns):
                        raise ValueError(f"Linha com número inválido de colunas: {row}")
                    batch.append(row)
                    if len(batch) >= batch_size:
                        insert_batch(cursor, schema_name, table_name, columns, batch)
                        batch.clear()

                if batch:
                    insert_batch(cursor, schema_name, table_name, columns, batch)


def parse_args():
    parser = argparse.ArgumentParser(
        description="Importa um arquivo CSV para uma tabela PostgreSQL."
    )
    parser.add_argument("--dsn", required=False, help="String de conexão DSN do PostgreSQL")
    parser.add_argument("--schema", required=False, help="Nome do schema de destino")
    parser.add_argument("--table", required=False, help="Nome da tabela de destino")
    parser.add_argument("--file", required=True, help="Caminho para o arquivo CSV")
    parser.add_argument(
        "--batch-size",
        type=int,
        default=1000,
        help="Quantidade de linhas por lote (padrão: 1000)",
    )
    if not parser.parse_known_args()[0].dsn:
        parser.set_defaults(dsn=dsn_default)
    if not parser.parse_known_args()[0].table:
        parser.set_defaults(table=table_default)
    if not parser.parse_known_args()[0].schema:
        parser.set_defaults(schema=schema_default)
    return parser.parse_args()


def main():
    args = parse_args()
    load_csv_to_table(args.dsn, args.schema, args.table, args.file, args.batch_size)


if __name__ == "__main__":
    main()