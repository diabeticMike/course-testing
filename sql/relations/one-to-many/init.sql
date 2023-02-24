create database db;
alter database db set timezone = 'utc';

-- work with a particular db (one postgresql instance could have many dbs)
\c db;

create table fathers (
    id uuid primary key,
    name text not null unique
);

create table children (
    id uuid primary key,
    name text not null unique,
    father_id uuid references fathers (id) on delete cascade
);

insert into fathers (
    id,
    name
) values (
    '95285f8f-4880-4258-8712-a622f413bd30',
    'jake'
),
(
    'ed75f0e9-ca28-4741-af93-a459ddbabe08',
    'german'
),
(
    'ed85f0e9-ca28-4741-af93-a459ddbabe08',
    'dave'
);

insert into children (
    id,
    name,
    father_id
) values (
    '683624a3-d03a-47d2-b8d1-5712fc199504',
    'bob',
    '95285f8f-4880-4258-8712-a622f413bd30'
),
(
    '75dabb89-f079-4510-94e0-cc8d23f67d19',
    'mike',
    'ed75f0e9-ca28-4741-af93-a459ddbabe08'
),
(
    'b61756fe-0d4b-452e-9189-0d49664116d8',
    'stacy',
    '95285f8f-4880-4258-8712-a622f413bd30'
),
(
    'a81756fe-0d4b-452e-9189-0d49664116d8',
    'john',
    null
);
