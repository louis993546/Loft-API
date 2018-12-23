# Database

All the basics of the database

## Basic specs

- RDBMS w/ PostgreSQL
- (Probably) connected with [`github.com/lib/pg`](https://github.com/lib/pq)

## Schema

### `loft`

| column      | type | not null | remark                                    |
|-------------|------|----------|-------------------------------------------|
| `id`        | uuid | true     | Primary key                               |
| `name`      | text | true     |                                           |
| `join_code` | text | true     | (TODO: I have designed a format for that) |




- `notes`
  - `id`
  - `loft_id`
  - `created_at`
- `member`
  - `id`
  - `loft_id`
  - `joined_at`
- `task`
  - `id`
  - `loft_id`
  - `created_at`
- `event`
  - `id`
  - `loft_id`
- `message`
  - `id`
  - `loft_id`

## Useful commands

### In command line

- `sudo -u postgres -i`
  - go to some kind of interactive mode as user `postgres`
- `sudo -u postgres {command}`
  - execute command as user `postgres`
- `postgres --version`
- `pg_dump -s {table name} > {output file name}`
  - dump the whole db into a kind of query
- `psql {table name}`
  - enter `psql` to a specific table

### In `psql`

- `\dt`
  - list of tables
- `SELECT version();`
  - list version of postgres