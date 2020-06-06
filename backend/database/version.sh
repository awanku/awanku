#!/usr/bin/env sh
if [[ "$DB_URL" == '' ]]; then
    echo 'please set database url as DB_URL in environment variable'
    exit 1
fi
go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate -database "${DB_URL}" -source file://./database/migrations version
