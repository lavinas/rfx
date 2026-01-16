import argparse
import csv
from pathlib import Path
import psycopg2
from psycopg2 import sql
from psycopg2.extras import execute_values
from decimal import Decimal, InvalidOperation, ROUND_HALF_UP

dsn_default = "postgresql://root:root@localhost:5435/cdc"
schema_default = "reports"
table_default = "reports_tables"

def format_value(value):
    if value == "":
        return None
    if isinstance(value, str):
        stripped = value.strip()
        if "." in stripped:
            integer_part, _, fractional_part = stripped.partition(".")
            if len(fractional_part) > 2:
                try:
                    decimal_value = Decimal(stripped)
                except InvalidOperation:
                    return value
                rounded = decimal_value.quantize(Decimal("0.01"), rounding=ROUND_HALF_UP)
                return format(rounded, "f")
        return value

    if isinstance(value, (float, Decimal)):
        decimal_value = Decimal(str(value))
        if decimal_value.as_tuple().exponent < -2:
            return decimal_value.quantize(Decimal("0.01"), rounding=ROUND_HALF_UP)
        return value

    return value


def insert_batch(cursor, schema_name, table_name, columns, rows):
    rows = [
        tuple(format_value(value) for value in row)
        for row in rows
    ]

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


def load_csv_to_table(dsn, schema_name, table_name, csv_path):
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

        with path.open(newline="", encoding="utf-8") as count_file:
            count_reader = csv.reader(count_file)
            next(count_reader, None)
            total_lines = sum(1 for row in count_reader if any(row))

        processed = 0

        with psycopg2.connect(dsn) as connection:
            with connection.cursor() as cursor:
                batch = []
                for row in reader:
                    if not any(row):
                        continue
                    if len(row) != len(columns):
                        raise ValueError(f"Linha com número inválido de colunas: {row}")
                    batch.append(row)
                    processed += 1
                    if total_lines and (processed % 500 == 0 or processed == total_lines):
                        print(f"Processed {processed}/{total_lines}")
                    if batch:
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
    load_csv_to_table(args.dsn, args.schema, args.table, args.file)


if __name__ == "__main__":
    main()