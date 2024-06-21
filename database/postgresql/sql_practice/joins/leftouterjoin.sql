-- CREATE SCHEMA IF NOT EXISTS sqlpractice;

-- CREATE TABLE if NOT EXISTS sqlpractice.products(
--     id SERIAL primary key,
--     name varchar(255)
-- );

-- CREATE table if NOT EXISTS sqlpractice.orders(
--     id SERIAL,
--     product_id int,
--     Foreign Key (product_id) REFERENCES  sqlpractice.products(id) ON DELETE SET NULL
-- );

-- DROP TABLE sqlpractice.orders;

-- INSERt INTO sqlpractice.products(name) values('bike') RETURNING *;


-- INSERT into sqlpractice.orders(product_id) values(3) RETURNING *;


-- products je ek pan orders ma nathi then left outer join ig

SELECT p.id as product_id, o.id as order_id, p.name as product_name FROM sqlpractice.products p
FULL OUTER JOIN sqlpractice.orders o
ON p.id = o.product_id;