#!/usr/bin/env sh
if [[ "$DB_URL" == '' ]]; then
    echo 'please set database url as DB_URL in environment variable'
    exit 1
fi

pg_dump --file=database/dump.sql --verbose "${DB_URL}"
