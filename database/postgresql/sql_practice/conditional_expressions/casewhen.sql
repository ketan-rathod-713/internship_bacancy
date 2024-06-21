select  id, name, 
CASE WHEN id = 1 THEN 'one'
    WHEN id = 2 THEN 'two'
    END as casing
 from student;