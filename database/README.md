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