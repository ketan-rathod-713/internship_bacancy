select m.name
from
employee m inner join employee e
on
m.id = e.managerId
group by (m.id, m.name)
having count(m.id) >= 5;


-- managers with atleast 5 employees