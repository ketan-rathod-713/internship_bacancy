-- Write your PostgreSQL query statement below

-- average selling price for each product


-- WITH finalTable as (SELECT us.product_id, us.units, (
--     SELECT pr.price * us.units
--     FROM Prices pr
--     WHERE pr.product_id = us.product_id AND us.purchase_date >= pr.start_date AND us.purchase_date <= pr.end_date
-- ) as price FROM
-- UnitsSold us)

-- SELECT product_id, ROUND((SUM(price)::float/SUM(units)::float)::numeric, 2) as average_price
-- from finalTable
-- group by (product_id)

-- SELECT product_id, 0 as average_price
-- FROM Prices 
-- WHERE product_id IN (SELECT DISTINCT product_id FROM UnitsSold)

WITH finalTable as (SELECT pr.product_id,
CASE WHEN units is null THEN 0 ELSE units * pr.price END as unitsprice,
CASE WHEN units is null THEN 0 ELSE units END as units
FROM Prices pr LEFT JOIN UnitsSold us
ON us.product_id IS NULL OR (pr.product_id = us.product_id AND us.purchase_date >= pr.start_date AND us.purchase_date <= pr.end_date))

-- SELECT product_id, AVG(unitsprice) FROM finalTable;

SELECT tt.product_id, 
CASE WHEN units = 0 THEN 0 
ELSE ROUND(((tt.unitsPrice)::float/(tt.units)::float)::numeric, 2) END 
AS average_price
FROM
(SELECT ft.product_id, SUM(unitsprice) as unitsPrice, SUM(units) as units
FROM finalTable ft
GROUP BY ft.product_id) as tt;