#!/bin/bash

docker exec tcms_database bash -c "psql -U tcms -d tcms < fixtures/load-fixtures.sql"
