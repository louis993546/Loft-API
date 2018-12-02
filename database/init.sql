CREATE TABLE IF NOT EXISTS loft(
    id uuid PRIMARY KEY,
    name text,
    join_code text,
    created_at timestamptz DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS join_request(
    id uuid PRIMARY KEY,
    loft_id uuid REFERENCES loft(id),
    name text,
    message text,
    created_at timestamptz DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS member(
    id uuid PRIMARY KEY,
    loft_id uuid REFERENCES loft(id),
    approved_at timestamptz,
    approved_by_member_id uuid,
    name text,
    phone text,
    email text,
    join_request_id uuid,
    FOREIGN KEY (approved_by_member_id) REFERENCES member(id)
);

CREATE TYPE noteformat AS ENUM ('COMMON_MARK_V_0_28');

CREATE TABLE IF NOT EXISTS note(
    id uuid PRIMARY KEY,
    loft_id uuid REFERENCES loft(id),
    creator_id uuid REFERENCES member(id),
    created_at timestamptz DEFAULT NOW(),
    format noteformat DEFAULT 'COMMON_MARK_V_0_28',
    content text
);

CREATE TABLE IF NOT EXISTS task(
    id uuid PRIMARY KEY,
    loft_id uuid REFERENCES loft(id),
    creator_id uuid REFERENCES member(id),
    created_at timestamptz DEFAULT NOW(),
    assignee_id uuid REFERENCES member(id),
    title text,
    due_date date
);

CREATE TABLE IF NOT EXISTS event(
    id	uuid,
    loft_id	uuid REFERENCES loft(id),
    creator_id	uuid REFERENCES member(id),
    created_at	timestamptz DEFAULT NOW(),
    start_time	timestamptz,
    end_time	timestamptz,
    title	text
);

CREATE TYPE messagetype AS ENUM('TEXT', 'IMAGE_REF');

CREATE TABLE IF NOT EXISTS message(
    id	uuid PRIMARY KEY,
    loft_id	uuid REFERENCES loft(id),
    created_at	timestamptz DEFAULT NOW(),
    sender_id	uuid REFERENCES member(id),
    content	text,
    type messagetype
);

CREATE TABLE IF NOT EXISTS session(
    id uuid PRIMARY KEY,
    member_id uuid REFERENCES member(id),
    created_at timestamptz DEFAULT NOW(),
    last_used_at timestamptz,
    last_used_ip cidr
);