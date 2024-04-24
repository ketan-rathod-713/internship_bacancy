-- average time of process per minute
select 
    machine_id, ROUND(avg(processTime)::numeric, 3) as processing_time
from (select a.machine_id, a.process_id, (b.timestamp - a.timestamp) as processTime
from activity a left join activity b
using (machine_id, process_id)
where a.activity_type = 'start' and b.activity_type = 'end') group by machine_id;


-- In aboce case there were two rows one has start and other has end so i have to do the join operation. is there any way i would have avoided this join operation.