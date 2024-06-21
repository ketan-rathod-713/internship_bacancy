select name from
customer where referee_id is null or referee_id != 2;

-- by default agar me is null nahi likhta then ye usko ginta hi nahi comparison ke liye.


-- In case of atleast 
select name, population, area from 
World
where area >= 3000000 or population >= 25000000;

--  We can also do same table field comparisons

select distinct author_id as id
from views
where author_id = viewer_id;