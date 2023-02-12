#!/bin/bash

# Run application migrations
if [[ $DATABASE_SCHEMA != "" && $DATABASE_MIGRATIONS_PATH != "" ]]; then
  echo "Searching for migrations..."
  uri="$DATABASE_SCHEMA://$DATABASE_USER:$DATABASE_PASSWORD@$DATABASE_HOST:$DATABASE_PORT/$DATABASE_NAME?sslmode=$DATABASE_SSLMODE"
  migrate -path $DATABASE_MIGRATIONS_PATH -database $uri up
fi

tail -f /dev/null
