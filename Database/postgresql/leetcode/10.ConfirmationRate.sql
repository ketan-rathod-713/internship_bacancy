-- Write your PostgreSQL query statement below

WITH totalCount as  (select user_id, count(*) as totalCount
from confirmations
group by user_id),
confirmedCount as (select user_id, count(user_id) as totalCount
from confirmations
group by (user_id, action) having action='confirmed'),
finalTable as (select s.user_id,CASE WHEN c.totalCount is null THEN 0 ELSE c.totalCount END as confirm, 
CASE WHEN t.totalCount is null THEN 0 ELSE t.totalCount END as totalCount
from
signups s left join confirmedCount c 
using (user_id)
left join 
totalCount t using (user_id))

select user_id, 
CASE WHEN confirm = 0 THEN ROUND(0::numeric, 2)
ELSE ROUND((confirm::float/totalcount::float)::numeric, 2) END as confirmation_rate
from finalTable
group by (user_id, confirmation_rate);


-- select * from finalTable;
