-- adding 1 day
select registrationtime + INTERVAL '1 day' as NextDay from datetimes.practice;


-- removing 1 second from time.datetime
select registrationtime + INTERVAL '1 second ago' as NextDay from datetimes.practice;


-- adding 1 day to date
select registrationtime + 1 as NextDay from datetimes.practice;

-- substract 1 day from date
select registrationtime - 1 as NextDay from datetimes.practice;

-- adding interval
select CURRENT_TIME + INTERVAL '1 hour' as tommorow;

-- now another example
Select CURRENT_DATE + INTERVAL '1 days ago' as yesterday;