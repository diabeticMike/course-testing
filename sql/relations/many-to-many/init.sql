CREATE DATABASE db;
ALTER DATABASE db SET TIMEZONE = 'UTC';

-- work with a particular db (one PostgreSQL instance could have many DBs)
\c db;

CREATE TABLE users (
    id uuid primary key,
    name text NOT NULL UNIQUE
);

CREATE TABLE teams (
    id uuid primary key,
    name text NOT NULL UNIQUE
);

CREATE TABLE users_teams (
    id uuid primary key,
    user_id uuid not null references users (id) on delete cascade,
    team_id uuid not null references teams (id) on delete cascade,
    unique (user_id, team_id)
);

INSERT INTO users (
    id,
    name
) VALUES (
    '95285f8f-4880-4258-8712-a622f413bd30',
    'user1'
),
(
    'ed75f0e9-ca28-4741-af93-a459ddbabe08',
    'user2'
),
(
    'd89342dc-9224-441d-a4af-bdd837a3b239',
    'user3'
);

INSERT INTO teams (
    id,
    name
) VALUES (
    '683624a3-d03a-47d2-b8d1-5712fc199504',
    'team1'
),
(
    '75dabb89-f079-4510-94e0-cc8d23f67d19',
    'team2'
),
(
    'b61756fe-0d4b-452e-9189-0d49664116d8',
    'team3'
);

INSERT INTO users_teams (
    id,
    user_id,
    team_id
) VALUES (
    'd7c4b4e9-b08e-4b4d-9e71-c90e19b12964',
    '95285f8f-4880-4258-8712-a622f413bd30',
    '683624a3-d03a-47d2-b8d1-5712fc199504'
),
(
    'd8c4b4e7-b08e-4b4d-9e71-c90e19b12964',
    'ed75f0e9-ca28-4741-af93-a459ddbabe08',
    'b61756fe-0d4b-452e-9189-0d49664116d8'
),
(
    'd8c4b4e9-b08e-4b4d-9e71-c90e19b12964',
    '95285f8f-4880-4258-8712-a622f413bd30',
    'b61756fe-0d4b-452e-9189-0d49664116d8'
),
(
    'c385dd14-352e-41e2-aecc-4a3152d51ef2',
    'ed75f0e9-ca28-4741-af93-a459ddbabe08',
    '75dabb89-f079-4510-94e0-cc8d23f67d19'
),
(
    'a385dd14-352e-41e2-aecc-4a3152d51ef1',
    'd89342dc-9224-441d-a4af-bdd837a3b239',
    '75dabb89-f079-4510-94e0-cc8d23f67d19'
);

