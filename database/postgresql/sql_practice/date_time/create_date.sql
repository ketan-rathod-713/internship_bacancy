-- create date time from the fields specified

-- syntax
-- make_date ( year int, month int, day int ) → date
-- make_interval ( [ years int [, months int [, weeks int [, days int [, hours int [, mins int [, secs double precision ]]]]]]] ) → interval
-- make_time ( hour int, min int, sec double precision ) → time
-- make_timestamp ( year int, month int, day int, hour int, min int, sec double precision ) → timestamp
-- make_timestamptz ( year int, month int, day int, hour int, min int, sec double precision [, timezone text ] ) → timestamp with time zone
-- now ( ) → timestamp with time zone

SELECT MAKE_DATE(2020, 1, 7) + INTERVAL '1 hour' as myBirthday;