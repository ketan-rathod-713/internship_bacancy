-- join students and subjects and then check user ne kiski exam di he

-- why count(*) was giving 1 
select s.student_id, s.student_name, sub.subject_name, count(e.subject_name) as attended_exams
from
students s cross join subjects sub
left join examinations e using(student_id, subject_name)
group by (s.student_id,s.student_name,sub.subject_name)


-- Here count(e.subject_name) is important as it will count that thing only. as it will count not null value only.