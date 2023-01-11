CREATE DATABASE db;
ALTER DATABASE db SET TIMEZONE = 'UTC';

-- work with a particular db (one PostgreSQL instance could have many DBs)
\c db;

CREATE TABLE fathers (
    id uuid primary key,
    name text NOT NULL UNIQUE
);

CREATE TABLE children (
    id uuid primary key,
    name text NOT NULL UNIQUE,
    father_id uuid references fathers (id) on delete cascade
);

INSERT INTO fathers (
    id,
    name
) VALUES (
    '95285f8f-4880-4258-8712-a622f413bd30',
    'Jake'
),
(
    'ed75f0e9-ca28-4741-af93-a459ddbabe08',
    'German'
),
(
    'ed85f0e9-ca28-4741-af93-a459ddbabe08',
    'Dave'
);

INSERT INTO children (
    id,
    name,
    father_id
) VALUES (
    '683624a3-d03a-47d2-b8d1-5712fc199504',
    'Bob',
    '95285f8f-4880-4258-8712-a622f413bd30'
),
(
    '75dabb89-f079-4510-94e0-cc8d23f67d19',
    'Mike',
    'ed75f0e9-ca28-4741-af93-a459ddbabe08'
),
(
    'b61756fe-0d4b-452e-9189-0d49664116d8',
    'Stacy',
    '95285f8f-4880-4258-8712-a622f413bd30'
),
(
    'a81756fe-0d4b-452e-9189-0d49664116d8',
    'John',
    NULL
);
