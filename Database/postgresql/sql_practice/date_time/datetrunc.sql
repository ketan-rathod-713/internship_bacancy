-- truncate to specific precision

-- day ke bad ki sari fields 00 ho jaegi or the default one
select DATE_TRUNC('day', CURRENT_TIMESTAMP);


-- Different syntaxes
-- date_trunc ( text, timestamp ) → timestamp
-- date_trunc ( text, timestamp with time zone, text ) → timestamp with time zone
-- date_trunc ( text, interval ) → interval