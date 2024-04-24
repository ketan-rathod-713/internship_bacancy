-- Write your PostgreSQL query statement below
WITH usersCount as (SELECT COUNT(user_id) FROM Users)

-- SELECT * FROM usersCount;

SELECT contest_id, (
    SELECT ROUND(((COUNT(rg.contest_id)::float/count::float) * 100)::numeric, 2)  FROM usersCount
) as percentage
FROM Register rg
INNER JOIN 
Users us USING (user_id)
GROUP BY (rg.contest_id)
ORDER BY percentage DESC, contest_id ASC;