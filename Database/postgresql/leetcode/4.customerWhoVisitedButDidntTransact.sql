select customer_id, count(customer_id) as count_no_trans
from visits where visit_id not in 
(select visit_id from transactions)
group by customer_id;

-- get all the customers visit id and check ki transaction me vo visit id nahi he.
-- and we will get duplicate customer_id too so count it and report it  to show number of time a user visited mall but didnt buy anything.