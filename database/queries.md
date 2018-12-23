# Queries

A list of all the queries for the API

- `INSERT INTO loft.loft (name, join_code) VALUES ('{loft name}', '{loft join code}');`
- `INSERT INTO loft.join_request (loft_id, name, message) values ('{loft id}', '{requester name}', '{request message}');`
- `INSERT INTO loft.member (loft_id, name) values ('{loft id}', '{member name}');`
- `INSERT INTO loft.note (loft_id, creator_id, format, content) values ('{loft id}', '{creator id}', '{format id}', '{content}');`
- `INSERT INTO loft.task (loft_id, creator_id, title) values ('{loft id}', '{creator id}', '{task title}');`
- `INSERT INTO loft.event (loft_id, creator_id, title) VALUES ('{loft id}', '{creator id}', '{event title}');`
- `loft=# insert into loft.message (loft_id, sender_id, content, type) values ('{loft id}', '{sender's member id}', '{content}', '{message type id}');`