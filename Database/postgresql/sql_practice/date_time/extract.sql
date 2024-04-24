-- extract any field from timestamp

-- syntax
-- extract ( field from timestamp ) → numeric
-- extract ( field from interval ) → numeric

select EXTRACT('hour' from CURRENT_TIMESTAMP);