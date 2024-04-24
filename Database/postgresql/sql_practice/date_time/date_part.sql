-- get subfield of time

-- date_part ( text, timestamp ) → double precision

select DATE_PART('hour', CURRENT_TIMESTAMP) -- equivalent to extract

-- different syntaxs for it.
-- date_part ( text, timestamp ) → double precision
-- date_part ( text, interval ) → double precision


