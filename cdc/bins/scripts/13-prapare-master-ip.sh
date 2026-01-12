#!/usr/bin/env sh

# This script replaces all occurrences of 'master.' with 'bins.' in the specified file.
set -eu

if [ "$#" -ne 1 ]; then
    printf 'usage: %s <file>\n' "$0" >&2
    exit 1
fi

sed -i 's/"IP0072T1"/bins\.master_ip0072t1/g' "$1"