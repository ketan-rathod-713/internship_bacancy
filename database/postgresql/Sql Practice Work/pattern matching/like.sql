-- string LIKE pattern [ESCAPE escape-character]
-- string NOT LIKE pattern [ESCAPE escape-character]

-- The LIKE expression returns true if the string matches the supplied pattern. (As expected, the NOT LIKE expression returns false if LIKE returns true, and vice versa. An equivalent expression is NOT (string LIKE pattern).)

SELECT * FROM (
    VALUES (1, 'good'), (2, 'amazing'), (3, 'goooood') 
) AS  t(id, names) WHERE names LIKE '____'
