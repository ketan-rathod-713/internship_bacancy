CREATE SCHEMA IF NOT EXISTS migration;

CREATE TABLE IF NOT EXISTS migration.tridip(
    id int,
    name varchar(100) not null,
    email varchar(100) not null,
    primary key (id)
);