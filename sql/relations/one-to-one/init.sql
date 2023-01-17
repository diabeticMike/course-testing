CREATE DATABASE db;
ALTER DATABASE db SET TIMEZONE = 'UTC';

-- work with a particular db (one PostgreSQL instance could have many DBs)
\c db;

CREATE TABLE users (
    id uuid primary key,
    name text not null unique
);

CREATE TABLE passports (
    id uuid primary key,
    code smallint not null unique,
    user_id uuid not null unique references users (id) on delete cascade
);

INSERT INTO users (
    id,
    name
) VALUES (
    '95285f8f-4880-4258-8712-a622f413bd30',
    'Olivia'
),
(
    'ed75f0e9-ca28-4741-af93-a459ddbabe08',
    'German'
);

INSERT INTO passports (
    id,
    code,
    user_id
) VALUES (
    '683624a3-d03a-47d2-b8d1-5712fc199504',
    '21234',
    '95285f8f-4880-4258-8712-a622f413bd30'
),
(
    '75dabb89-f079-4510-94e0-cc8d23f67d19',
    '32210',
    'ed75f0e9-ca28-4741-af93-a459ddbabe08'
);
