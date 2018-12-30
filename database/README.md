# Database

All the basics of the database

## Schema

<!-- TODO -->

## How to initial DB locally

1. Install PostgreSQL
2. Setup user + role
3. Setup DB
4. Other useful commands

### Install PostgreSQL

Just read the [current documentation](https://www.postgresql.org/docs/current/tutorial-install.html) on how to install it

### Setup user + role

1. `createdb <DB name>`
   - If it is successful, congratulations!
   - You will probably see something like "role 'username' does not exist"
2. `sudo -u postgres -i`
   - Start to use bash as user `postgres` interactively
3. `createuser <your user name>`
4. `psql`
   - Enter the `PostgreSQL` CLI (SQL is basically a scripting language afterall)
5. `alter user <your user name> createdb;`
6. `\q`
   - Exit `psql`
7. exit
   - Stop pretending you are user `postgres`

### Setup DB

1. `created <db name>`
2. (TODO: schema)

### Other userful commands

#### In command line

- `dropdb <db name>`
- `sudo -u postgres {command}`
  - execute command as user `postgres`
- `postgres --version`
- `pg_dump -s {table name} > {output file name}`
  - dump the whole db into a kind of query
- `psql {table name}`
  - enter `psql` to a specific table

#### In `psql`

- `\dt`
  - list of tables
- `SELECT version();`
  - list version of postgres
- `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
  - so that `uuid_generate_v4()` works