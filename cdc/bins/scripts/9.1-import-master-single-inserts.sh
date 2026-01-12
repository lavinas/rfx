#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
SQL_FILE="$SCRIPT_DIR/../files/9.1-base-master_single_message.sql"

if [[ ! -f "$SQL_FILE" ]]; then
    echo "SQL file not found: $SQL_FILE" >&2
    exit 1
fi

PGPASSWORD="root" PGOPTIONS="--search_path=bins" psql -h localhost -p 5435 -U root -d cdc -f "$SQL_FILE" > result.log 2>&1
