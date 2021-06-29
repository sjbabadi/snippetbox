# Snippetbox
###  To run Snippetbox:
- Create a database on a running Postgres server
- Create a `.env` file and include the following:
    - `PG_HOST`, `PG_PORT`, `PG_USER`, `PG_PWD`, `DB_NAME`, and `PORT`
- Install the [Postgres driver for Go](https://github.com/lib/pq)
- Run `psql -d $DB_NAME < db/setup.sql` from the root directory of snippetbox to create the necessary schema and any seed data
- `make run`
