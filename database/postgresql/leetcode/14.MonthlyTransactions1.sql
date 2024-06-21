-- Write your PostgreSQL query statement below
-- concat(extract('year' from trans_date),'-', extract('month' from trans_date))
SELECT 
substring(trans_date::text from 0 for 8) as month, 
country, 
count(*) as trans_count,
SUM(CASE WHEN state='approved' THEN 1 ELSE 0 END) as approved_count,
SUM(amount) as trans_total_amount,
SUM(CASE WHEN state='approved' THEN amount ELSE 0 END) as approved_total_amount

FROM Transactions
group by (month, country);


