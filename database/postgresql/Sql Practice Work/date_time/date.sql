
create schema IF NOT EXISTS datetimes;

create table IF NOT EXISTS datetimes.practice(
    id serial primary key,
    registrationTime date,
    yesterdayTime date
);


insert into datetimes.practice(registrationTime, yesterdayTime) values(CURRENT_DATE, CURRENT_DATE - INTERVAL '1 day');

SELECT * from datetimes.practice;

-- CURRENT_DATE to get current date and INTERVAL To minus that.