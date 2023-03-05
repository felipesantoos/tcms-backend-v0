# How to run the project?

1. Install docker;
2. Go to the project root directory;
3. Run in the CLI:
    ```shell
    docker compose rm -sf
    docker compose up --build
    docker exec tcms_database bash -c "psql -U tcms -d tcms < fixtures/load-fixtures.sql"
    ```
