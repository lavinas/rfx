import os
import sys
import psycopg2

DNS = os.getenv("PG_DSN", "dbname=cdc user=root password=root host=localhost port=5435")

def parse_and_insert(filename, dsn=None):
    with psycopg2.connect(DNS) as conn, conn.cursor() as cur:
        with open(filename, "r", encoding="utf-8", errors="replace") as f:
            total_lines = sum(1 for _ in f)
            f.seek(0)
            for idx, line in enumerate(f, start=1):
                if (not line.strip()
                        or line.startswith("VISA INTERCHANGE DIRECTORY")
                        or len(line) < 97):
                    continue
                if idx % 200 == 0:
                    print(f"Processed {idx}/{total_lines}")
                id = line[0:8].strip()
                code = line[8:11].strip()
                issuer = line[17:97].replace("'", "").strip()
                if not id:
                    continue
                cur.execute(
                    "INSERT INTO bins.visa_issuers_main (id, code, issuer) VALUES (%s, %s, %s);",
                    (id, code, issuer),
                )
if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python 27-import-visa-issuers-main.py <filename>")
    else:
        parse_and_insert(sys.argv[1])
