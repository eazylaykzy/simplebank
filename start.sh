#!/bin/sh

set -e

echo "run db migration"
. /app/app.env
echo "DB source: $DB_SOURCE"
#source /app/app.env # this highlights error
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"