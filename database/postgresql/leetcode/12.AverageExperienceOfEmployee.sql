
SELECT project_id, ROUND(AVG(experience_years)::numeric, 2) as average_years
FROM 
Project LEFT JOIN EMPLOYEE USING (employee_id)
GROUP BY project_id;