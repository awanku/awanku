#!/usr/bin/env sh
if [[ "$DB_URL" == '' ]]; then
    echo 'please set database url as DB_URL in environment variable'
    exit 1
fi

psql "${DB_URL}" <<-EOF
drop schema public cascade;
create schema public;
EOF
