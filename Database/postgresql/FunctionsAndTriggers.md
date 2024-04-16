# SQL FUNCTIONS
 
You can create SQL scalar functions and SQL table functions.
 
PL/pgSQL is a procedural programming language extension for PostgreSQL, which allows developers to write stored procedures and functions in a more procedural style, akin to languages like PL/SQL in Oracle or T-SQL in Microsoft SQL Server.
 
Variables: You can declare variables to store values within PL/pgSQL functions.
Control Structures: PL/pgSQL supports control structures like IF, ELSEIF, ELSE, LOOP, WHILE, and FOR.
Exception Handling: You can handle exceptions using BEGIN ... EXCEPTION ... END blocks to catch and handle errors gracefully.

Cursors: PL/pgSQL allows you to work with cursors to process a set of rows returned by a query.
Transactions: You can manage transactions explicitly using BEGIN, COMMIT, and ROLLBACK statements within your functions

## Example

```
CREATE OR REPLACE FUNCTION calculate_sum(a INTEGER, b INTEGER)
RETURNS INTEGER AS $$
DECLARE
    result INTEGER;
BEGIN
    result := a + b;
    RETURN result;
END;
$$ LANGUAGE plpgsql;
```

## Example using if else

```
CREATE OR REPLACE FUNCTION is_positive(num INTEGER)
RETURNS BOOLEAN AS $$
BEGIN
    IF num > 0 THEN
        RETURN TRUE;
    ELSE
        RETURN FALSE;
    END IF;
END;
$$ LANGUAGE plpgsql;
```

## Example using loops

```
CREATE OR REPLACE FUNCTION print_numbers()
RETURNS VOID AS $$
DECLARE
    i INTEGER := 1;
BEGIN
    LOOP
        EXIT WHEN i > 5;
        RAISE NOTICE '%', i;
        i := i + 1;
    END LOOP;
END;
$$ LANGUAGE plpgsql;
```

## Example Factorial 

```
CREATE OR REPLACE FUNCTION factorial(n INTEGER)
RETURNS INTEGER AS $$
BEGIN
    IF n <= 1 THEN
        RETURN 1;
    ELSE
        RETURN n * factorial(n - 1);
    END IF;
END;
$$ LANGUAGE plpgsql;
```

-- CREATE OR REPLACE FUNCTION get_student_info()
-- RETURNS SETOF school.student AS $$
-- BEGIN
--     RETURN QUERY SELECT *FROM school.student WHERE school.student.id = 1;
-- END;
-- $$ LANGUAGE plpgsql;

-- SELECT get_student_info();

-- CREATE OR REPLACE FUNCTION divide(a INTEGER, b INTEGER, OUT quotient FLOAT, OUT remainder FLOAT)
-- RETURNS RECORD AS $$
-- BEGIN
--     quotient := a / b;
--     remainder := a % b;
-- END;
-- $$ LANGUAGE plpgsql;

-- -- crreating temporrary tables 
-- CREATE OR REPLACE FUNCTION process_data()
-- RETURNS VOID AS $$
-- BEGIN
--     CREATE TEMP TABLE temp_table AS SELECT * FROM some_table;
--     -- Process temp_table
--     DROP TABLE temp_table;
-- END;
-- $$ LANGUAGE plpgsql;


----
TRIGGER
In PostgreSQL, NEW is a special keyword used in trigger functions. It represents the new row being inserted, updated, or deleted in the trigger's associated table. When a trigger fires, NEW allows access to the values of the columns in the new row being affected by the triggering event.

-- Using plpgsql in triggers

-- CREATE OR REPLACE FUNCTION calculate_age()
-- RETURNS TRIGGER AS $$
-- BEGIN
-- -- Jo new data aaya he usme changes kar do ha ha
--     NEW.age := DATE_PART('year', CURRENT_DATE) - DATE_PART('year', NEW.birth_date);
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER update_age
-- BEFORE INSERT ON school.student
-- FOR EACH ROW
-- EXECUTE FUNCTION calculate_age();

Key points
- Use of NEW keyword in function for refering new inserted row
- Function is returning a Trigger
- At the end functions is inside returnning NEW itself















