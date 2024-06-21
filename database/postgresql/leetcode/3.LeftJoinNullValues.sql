select name, unique_id from
employees e1 left join employeeuni e2
on e1.id = e2.id;

-- jo field match nahi hogi vo automatic null aa jaegi

-- https://leetcode.com/problems/replace-employee-id-with-the-unique-identifier/?envType=study-plan-v2&envId=top-sql-50

-- Above query can be written like this also

select name, (select unique_id from employeeuni 
where employeeuni.id = employees.id
) as unique_id from employees;

-- another inner joinn

select product_name, year, price
from sales s inner join product p
using (product_id); 

