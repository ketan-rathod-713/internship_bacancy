SELECT a.id
FROM Weather a, Weather b
WHERE a.recordDate - 1 = b.recordDate
AND a.temperature > b.temperature;

-- id where the temperature for current date is greater then the temperature from previous date.