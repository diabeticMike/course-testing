create database db;
alter database db set timezone = 'utc';

-- work with a particular db (one postgresql instance could have many dbs)
\c db;

create table "user" (
    id uuid primary key,
    name text not null unique
);

create table "team" (
    id uuid primary key,
    name text not null unique
);

create table "user_team" (
    user_id uuid not null references "user" (id) on delete cascade,
    team_id uuid not null references "team" (id) on delete cascade,
    unique (user_id, team_id)
);

insert into "user" (
    id,
    name
) values (
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

insert into "team" (
    id,
    name
) values (
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

insert into "user_team" (
    user_id,
    team_id
) values (
    '95285f8f-4880-4258-8712-a622f413bd30',
    '683624a3-d03a-47d2-b8d1-5712fc199504'
),
(
    'ed75f0e9-ca28-4741-af93-a459ddbabe08',
    'b61756fe-0d4b-452e-9189-0d49664116d8'
),
(
    '95285f8f-4880-4258-8712-a622f413bd30',
    'b61756fe-0d4b-452e-9189-0d49664116d8'
),
(
    'ed75f0e9-ca28-4741-af93-a459ddbabe08',
    '75dabb89-f079-4510-94e0-cc8d23f67d19'
),
(
    'd89342dc-9224-441d-a4af-bdd837a3b239',
    '75dabb89-f079-4510-94e0-cc8d23f67d19'
);

