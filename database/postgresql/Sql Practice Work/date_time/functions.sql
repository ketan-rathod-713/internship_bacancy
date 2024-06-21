-- Date time functions

-- age only works on date or timestamp but on on time with timezon
select age(CURRENT_DATE + 2, CURRENT_DATE);

-- it will not work
-- select age(CURRENT_TIME, CURRENT_TIME + INTERVAL '1.5 hour')

select age(TIMESTAMP '20-10-2021');
