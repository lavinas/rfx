import os
import sys
import psycopg2

DSN = os.getenv("PG_DSN", "dbname=cdc user=root password=root host=localhost port=5435")

def parse_bins(input_file):
    with open(input_file, "r", encoding="utf-8") as counter:
        total_lines = sum(1 for _ in counter)
    print(f"Total lines to process: {total_lines}")
    processed = 0
    with psycopg2.connect(DSN) as conn, conn.cursor() as cur, open(input_file, "r", encoding="utf-8") as f:
        for line in f:
            processed += 1
            if processed % 200 == 0 or processed == total_lines:
                print(f"Processed {processed}/{total_lines}")
            if len(line) < 49 or not line[:8].strip().isdigit():
                continue
            id_val = line[0:8].strip()
            bin_val = line[41:49].strip()
            if id_val and bin_val:
                cur.execute(
                    "INSERT INTO bins.visa_issuers_bin (id, bin) VALUES (%s, %s) ON CONFLICT DO NOTHING",
                    (id_val, bin_val),
                )

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python parse_bins.py <input_file>")
    else:
        parse_bins(sys.argv[1])