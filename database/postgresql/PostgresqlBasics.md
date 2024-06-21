## Basics

- Key words and unquoted identifiers are case-insensitive.

- By Convention keywords in uppercase and names in lowercase.

- delimited identifier or quoted identifier. -> "select" will be taken as identifier while select would be taken as keyword.

#### making table name cases sensitive
`UPDATE "my_table" SET "a" = 5;`

- Quoting an identifier also makes it case-sensitive, whereas unquoted names are always folded to lower case.

- "Foo" and "FOO" are different

#### strings and constants
- Three kinds of implicitly-typed constants in PostgreSQL: strings, bit strings, and numbers.

- Two string constants that are only separated by whitespace with at least one newline are concatenated and effectively treated as if the string had been written as one constant.
```
SELECT 'foo' 
'bar';
is same as
SELECT 'foobar';

SELECT 'foo'      'bar'; // but this is not valid
```

- PostgreSQL also accepts “escape” string constants just like c
Can be done by writting E (upper or lower case) just before the opening single quote.
 e.g., E'foo'. 
 Then i can use \b \t \n etc.
 
- Dollar quoting ( when too many single quotes in string contstant )

```
$$Dianne's horse$$
$SomeTag$Dianne's horse$SomeTag$
```

#### Column Reference in query 
correlation.columnname

#### Positional Parameter
- Using $ dollar sign
- $1 it takes the first parameter passed in function argument.

```
 CREATE FUNCTION dept(INTEGER) RETURNS text
   AS $$ SELECT name from school.student where id = $1; $$
   LANGUAGE SQL;

   SELECT dept(5);
   DROP FUNCTION dept;
```

#### Subscripts
If expression yields type array then we can access element by
expression[subscript]

##### For array slice
expression[lower_subscript:upper_subscript]

```
CREATE TABLE IF NOT EXISTS arrays.schools(
	id SERIAL,
  names TEXT[],
  ranks INTEGER[]
);

INSERT INTO arrays.schools (ranks, names) VALUES
  (ARRAY[1, 2, 3, 4, 5], ARRAY['John', 'Alice', 'Bob']),
  (ARRAY[10, 20, 30], ARRAY['Mary', 'David']);
  
SELECT ranks[2] AS third_element FROM arrays.schools WHERE id = 1;

SELECT ranks[2:3] AS third_element FROM arrays.schools WHERE id = 1;
```

#### FIeld Selection

expression.fieldname

We can also create a composite type in sql

```
CREATE TYPE address AS (
    street TEXT,
    city TEXT,
    zipcode TEXT
);
```

- It is very handy when we are using complex data structures.

```
-- Create a composite type
CREATE TYPE address AS (
    street TEXT,
    city TEXT,
    zipcode TEXT
);

-- Create a table with a composite column
CREATE TABLE my_table (
    id SERIAL PRIMARY KEY,
    person_name TEXT,
    person_address address
);

-- Insert some data into the table
INSERT INTO my_table (person_name, person_address) VALUES
  ('John Doe', ROW('123 Main St', 'Anytown', '12345')),
  ('Alice Smith', ROW('456 Elm St', 'Sometown', '67890'));

-- Accessing fields within the composite type
-- Access the city field within the person_address composite type in the first row
SELECT (person_address).city AS person_city FROM my_table WHERE id = 1;

-- Access the street field within the person_address composite type in the second row
SELECT (person_address).street AS person_street FROM my_table WHERE id = 2;
```

#### Operator Invocations

Function calls
function_name ([expression [, expression ... ]] )

#### Aggregate Expressions

An aggregate expression represents the application of an aggregate function across the rows selected by a query. An aggregate function reduces multiple inputs to a single output value, such as the sum or average of the inputs

```
SYNTAX
aggregate_name (expression [ , ... ] [ order_by_clause ] ) [ FILTER ( WHERE filter_clause ) ] // Default is ALL
aggregate_name (ALL expression [ , ... ] [ order_by_clause ] ) [ FILTER ( WHERE filter_clause ) ] // for ALL expression invokes once for each row
aggregate_name (DISTINCT expression [ , ... ] [ order_by_clause ] ) [ FILTER ( WHERE filter_clause ) ] // for DISTINCT : invokes the aggregate once for each distinct value of the expression.
aggregate_name ( * ) [ FILTER ( WHERE filter_clause ) ] // generally used for count(*) only
aggregate_name ( [ expression [ , ... ] ] ) WITHIN GROUP ( order_by_clause ) [ FILTER ( WHERE filter_clause ) // used for ordered-set aggregate functions
```

- Most aggregate functions ignore null inputs, so that rows in which one or more of the expression(s) yield null are discarded. This can be assumed to be true, unless otherwise specified, for all built-in aggregates.

```
for eg. SELECT count(city) from school.student;  -- it ignores null values in city column
```

#### by default ignore null values
- For example, count(*) yields the total number of input rows; count(f1) yields the number of input rows in which f1 is non-null, since count ignores nulls; and count(distinct f1) yields the number of distinct non-null values of f1.

- Some aggregate func don't require ordering for eg. sum, min, max etc.
While some require such as array_agg, string_agg : Hence for them there is special clause which is ORDER BY. because ordering do matter in such aggregate functions.

```
SELECT array_agg(a ORDER BY b DESC) FROM table;

SELECT string_agg(a, ',' ORDER BY a) FROM table; // ORDER BY goes after all the arguments completed.

SELECT string_agg(a ORDER BY a, ',') FROM table;  -- This is incorrect

SELECT string_agg(DISTINCT city, ',') FROM school.student;
SELECT array_agg(DISTINCT city) FROM school.student;
```

TODO: see extra details after this doc 4.2.7

#### Window Function Calls
- A window function call represents the application of an aggregate-like function over some portion of the rows selected by a query.Just like sliding window ha ha.
- Here we are not grouping all selected rows just like non window aggregate calls ->  each row remains separate in the query output
- However the window function has access to all the rows that would be part of the current row's group according to the grouping specification PARTITION BY

SYNTAX 
function_name ([expression [, expression ... ]]) [ FILTER ( WHERE filter_clause ) ] OVER window_name
function_name ([expression [, expression ... ]]) [ FILTER ( WHERE filter_clause ) ] OVER ( window_definition )
function_name ( * ) [ FILTER ( WHERE filter_clause ) ] OVER window_name
function_name ( * ) [ FILTER ( WHERE filter_clause ) ] OVER ( window_definition )

window_definition syntax
[ existing_window_name ]
[ PARTITION BY expression [, ...] ]
[ ORDER BY expression [ ASC | DESC | USING operator ] [ NULLS { FIRST | LAST } ] [, ...] ]
[ frame_clause ]

where frame_start and frame_end can be one of

UNBOUNDED PRECEDING
offset PRECEDING
CURRENT ROW
offset FOLLOWING
UNBOUNDED FOLLOWING

and frame_exclusion can be one of

EXCLUDE CURRENT ROW
EXCLUDE GROUP
EXCLUDE TIES
EXCLUDE NO OTHERS

- Similar to an aggregate function, a window function operates on a set of rows. However, it does not reduce the number of rows returned by the query.

```
SELECT
	product_name,
	price,
	group_name,
	AVG (price) OVER (
	   PARTITION BY group_name
	)
FROM
	products
	INNER JOIN 
		product_groups USING (group_id);
```
		
It gives avg price for each column and thus not reduce the columns.

##### IMPORTANT NOTE
A window function always performs the calculation on the result set after the JOIN, WHERE, GROUP BY and HAVING clause and before the final ORDER BY clause in the evaluation order.

```
window_function(arg1, arg2,..) OVER (
   [PARTITION BY partition_expression]
   [ORDER BY sort_expression [ASC | DESC] [NULLS {FIRST | LAST }])  
```
   
- The PARTITION BY clause is optional. If you skip the PARTITION BY clause, the window function will treat the whole result set as a single partition.

- The ORDER BY clause specifies the order of rows in each partition to which the window function is applied. 

- The ORDER BY clause uses the NULLS FIRST or NULLS LAST option to specify whether nullable values should be first or last in the result set. The default is NULLS LAST option.

- The frame_clause defines a subset of rows in the current partition to which the window function is applied. This subset of rows is called a frame.

- TODO: more window function examples.

#### Type Casts

- Syntax
CAST ( expression AS type )
expression::type

```
SELECT CAST('42' AS INTEGER);

SELECT '42'::INTEGER; // only in postgresql
```


#### Collation Expressions
used for comparison and sorting you can change it.
TODO: ??

#### Scalar Subqueries 
- DOC_REF: 4.2.11

- returns exactly one row, error if more then one row.
- if subquery returns no row then no error ha ha.

- subquery can refer to variables from the surrounding query which will act as constants during any one evalution of the subquery.

```
// finding largest city population in state, here state name for all it's cities is constant and so on...
SELECT name, (SELECT max(pop) FROM cities WHERE cities.state = states.name)
    FROM states;
```
    
#### Array Constructors
- DOC_REF: 4.2.12

- TODO: https://www.postgresql.org/docs/current/sql-expressions.html#SQL-SYNTAX-COLLATE-EXPRS

```
SELECT ARRAY[1,2,22.7];
// This creates an array.

// We can also cast sub query result to array as follows

SELECT ARRAY(SELECT age from school.student);

// The subquery must return a single column.
```

#### ROW constructor

```
SELECT ROW(1,2.5,'this is a test');

SELECT ROW(school.student.*, 42) FROM school.student; // It takes all row values from student and then add 42 at the end of each row.
```

- By default, the value created by a ROW expression is of an anonymous record type. If necessary, it can be cast to a named composite type — either the row type of a table, or a composite type created with CREATE TYPE AS. An explicit cast might be needed to avoid ambiguity

TODO: doubt in below query.
```
SELECT ROW(school.student.*) IS NULL FROM school.student; -- Detect all null rows
```
----
#### Expression Evaluation Rules 
- DOC_REF: 4.2.14
- TODO:

#### CALLING FUNCTION USING 
- TODO:

<hr>

### Data Definition

#### create table

```
CREATE TABLE products (
    product_no integer,
    name text,
    price numeric
);

DROP TABLE IF EXISTS products;
```

#### default values

- A column can be assigned a default value. if no values passed then default value used.
- by default, default value is null.

```
CREATE TABLE products (
    product_no integer,
    name text,
    price numeric DEFAULT 9.99
);

// Default value can also be an expression, which will be evaluated whenever the default value is inserted. 

CREATE TABLE products (
    product_no integer DEFAULT nextval('products_product_no_seq'),
    ...
);

// above syntax so common so that there is shorthand for it.

CREATE TABLE products (
    product_no SERIAL,
    ...
);
```

#### Generated Columns 

- special column that is always calculated from other columns.

##### Types
- stored - when doing insert / update it get stored in table as if it was a normal column.
- virtual - for read purposes

```
CREATE TABLE people (
    ...,
    height_cm numeric,
    height_in numeric GENERATED ALWAYS AS (height_cm / 2.54) STORED
);

CREATE TABLE data_definition.generate_column(
 data_km int,
 data_meter int GENERATED ALWAYS AS (data_km * 1000) STORED
 );

// Try various operations and check data updated or not !!??
INSERT INTO data_definition.generate_column(data_km) VALUES(34);

UPDATE data_definition.generate_column
SET data_km = 45;
```

#### Contraints

##### 1. Check Constraint
- most basic one for boolean expreession checking.

```
CREATE TABLE products (
    product_no integer,
    name text,
    price numeric CHECK (price > 0)
);
```

- we can give constaint a name so that we can refer it when error message produced.
eg. price numeric CONSTRAINT postive_price (price > 0)

- We can also let it refer to several columns.

```
CREATE TABLE products (
    product_no integer,
    name text,
    price numeric CHECK (price > 0), -- column constaint
    discounted_price numeric CHECK (discounted_price > 0), -- column constaint
    CHECK (price > discounted_price)  -- table constraint

    // OR
    //  CONSTRAINT valid_discount CHECK (price > discounted_price)
);
```
- It should be noted that a check constraint is satisfied if the check expression evaluates to true or the null value. For this problem you should use not null constraint.
- several conditions can be matched by AND, OR, etc.

##### 2. Not Null Constraint

- column must not assume the null value.

```
product_no integer NOT NULL

// For more then one constraint : price numeric NOT NULL CHECK (price > 0)
```

##### 3. Unique Constriant

- To define a unique constraint for a group of columns, write it as a table constraint with the column names separated by commas:
for eg. UNIQUE (a, c)

- by default we can still insert NULL values if UNIQUE constraint is there.
The default behavior can be specified explicitly using NULLS DISTINCT. T

- if we don't want null values to be more then one then 
UNIQUE NULLS NOT DISTINCT (product_no)

```
// with constrain name
CREATE TABLE products (
    product_no integer CONSTRAINT must_be_different UNIQUE,
    name text,
    price numeric

    // OR
    // you can write this in table definition too
    // UNIQUE NULLS NOT DISTINCT (product_no)
    // becaue of NULLS NOT DISTINCT it will assume all nulls to be same, default behaviour is NULLS DISTINCT.
);

```

##### 4. Primary Key
- UNIQUE + NOT NULL

```
CREATE TABLE example (
    a integer, // OR // write here // PRIMARY KEY
    b integer,
    c integer,
    PRIMARY KEY (a)
);
```

##### 5. Foreign Keys

```
CREATE TABLE orders (
    order_id integer PRIMARY KEY,
    product_no integer REFERENCES products (product_no),
    quantity integer
);
```

- Now it is impossible to create orders with non-NULL product_no entries that do not appear in the products table.

We can also shorten above code.
```
CREATE TABLE orders (
    order_id integer PRIMARY KEY,
    product_no integer REFERENCES products,
    quantity integer
); 
// it will point to primary key of referencing table

// Also it can be group of columns.
CREATE TABLE t1 (
  a integer PRIMARY KEY,
  b integer,
  c integer,
  FOREIGN KEY (b, c) REFERENCES other_table (c1, c2)
);

```

- A foreign key can also constrain and reference a group of columns. As usual, it then needs to be written in table constraint form. Here is a contrived syntax example:

```
CREATE TABLE t1 (
  a integer PRIMARY KEY,
  b integer,
  c integer,
  FOREIGN KEY (b, c) REFERENCES other_table (c1, c2)
);
```

- Action to take when referenced row is deleted...

- RESTRICT prevents deletion of a referenced row. 
- NO ACTION is default, it only raises an error.
- CASCADE specifies that when a referenced row is deleted, row(s) referencing it should be automatically deleted as well.

###### What is referencing row and what is referenced row ??

- referenced means parent and referencing means child one.

- The appropriate choice of ON DELETE action depends on what kinds of objects the related tables represent. When the referencing table represents something that is a component of what is represented by the referenced table and cannot exist independently, then CASCADE could be appropriate. If the two tables represent independent objects, then RESTRICT or NO ACTION is more appropriate
- The actions SET NULL or SET DEFAULT can be appropriate if a foreign-key relationship represents optional information. For example, if the products table contained a reference to a product manager, and the product manager entry gets deleted, then setting the product's product manager to null or a default might be useful.

##### 6. Exclusion Constraint ??


#### System Columns

- every table has several system columns, hence we can't name our column name to this.
- COLUMNS
    - tableoid 
    - xmin : The identity (transaction ID) of the inserting transaction for this row version. (A row version is an individual state of a row; each update of a row creates a new row version for the same logical row.)
    - cmin
    - xmax
    - cmax
    - ctid

#### Modifying Tables

All this above actions are performed using `ALTER TABLE` command.

- Add columns

    ```
    ALTER TABLE products ADD COLUMN description text;
    // can also specify Check contraints and all but default value should satisfy it. default value is null. you can change it.
    ```

- Remove columns
    - if the column is referenced by a foreign key constraint of another table, PostgreSQL will not silently drop that constraint. You can authorize dropping everything that depends on the column by adding CASCADE:


    ``` 
        ALTER TABLE products DROP COLUMN description CASCADE;
    ```

- Add constraints
    - To Add table constraints
    ```
    ALTER TABLE products ADD CHECK (name <> '');
    ALTER TABLE products ADD CONSTRAINT some_name UNIQUE (product_no);
    ALTER TABLE products ADD FOREIGN KEY (product_group_id) REFERENCES product_groups;
    ```

    - To add column level contraints
    ```
    ALTER TABLE products ALTER COLUMN product_no SET NOT NULL;
    ```
- Remove constraints
    - need to know contraint name to drop it. you can also add cascade if some other contraints depend on it. for eg. foreign key.
    ```
    ALTER TABLE products DROP CONSTRAINT some_name;
    ```
    - if don't know contraint name
    ```
    ALTER TABLE products ALTER COLUMN product_no DROP NOT NULL;
    // (Recall that not-null constraints do not have names.)
    ```

- Change default values
    - `ALTER TABLE products ALTER COLUMN price SET DEFAULT 7.77;` It doesn' affect existing rows. it is for future inserts.
    - `ALTER TABLE products ALTER COLUMN price DROP DEFAULT;` To remove it.

- Change column data types
    - `ALTER TABLE products ALTER COLUMN price TYPE numeric(10,2);`
    - it will only succeed if it is feasible to convert data type by implicit cast.
    - If a more complex conversion is needed, you can add a USING clause that specifies how to compute the new values from the old.
    - drop contraints before doing it as it can produce an error.

- Rename columns
    - `ALTER TABLE products RENAME COLUMN product_no TO product_number;`

- Rename tables
    - `ALTER TABLE products RENAME TO items;`


#### Privileges

- When an object is created, it is assigned an owner.
- to allow other users to use it, priviledges needs to be assigned.
- There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, USAGE, SET and ALTER SYSTEM.

- assign object a new owner
    - `ALTER TABLE table_name OWNER TO new_owner;`
- assign priviledge with grant
    - `GRANT UPDATE ON accounts TO joe;`
- To revoke previously granted priviledge
    - `REVOKE ALL ON accounts FROM PUBLIC;`
- Ordinarily, only the object's owner (or a superuser) can grant or revoke privileges on an object. However, it is possible to grant a privilege “with grant option”, which gives the recipient the right to grant it in turn to others. If the grant option is subsequently revoked then all who received the privilege from that recipient (directly or through a chain of grants) will lose the privilege
- 

#### TODO: Row Security Policies

- In addition to priviledges, tables can have row security policies that restrict, on a per-user basis, which rows can be returned by normal queries or inserted, updated, or deleted by data modification commands. This feature is also known as Row-Level Security. By default, tables do not have any policies, so that if a user has access privileges to a table according to the SQL privilege system, all rows within it are equally available for querying or updating.

- Reference : https://www.postgresql.org/docs/current/ddl-rowsecurity.html

#### Schemas

- There are several reasons why one might want to use schemas:
    - To allow many users to use one database without interfering with each other.
    - To organize database objects into logical groups to make them more manageable.
    - Third-party applications can be put into separate schemas so they do not collide with the names of other objects.


- Schema Normal Operations
    - Creating a schema
    - `CREATE SCHEMA myschema;`
    - Access by `database.schema.table`
    - Drop it by `DROP SCHEMA myschema;`
    - Create table in it by `CREATE TABLE myschema.mytable (....);`
    - To also drop all the objects associated with it`DROP SCHEMA myschema CASCADE;`
    - Schema starting with pg_ are only for system use, user can not create it.
    - Often you will want to create a schema owned by someone else (since this is one of the ways to restrict the activities of your users to well-defined namespaces). The syntax for that is:
        - `CREATE SCHEMA schema_name AUTHORIZATION user_name;`
    - by default public schema is used.
    - SHOW search_path; // it will output "public" // it is the path where normally get stored.
    - To put our new schema in search path `SET search_path TO myschema,public;`

##### Schemas and Priviledges

    - By default, users cannot access any objects in schemas they do not own. To allow that, the owner of the schema must grant the USAGE privilege on the schema. By default, everyone has that privilege on the schema public. To allow users to make use of the objects in a schema, additional privileges might need to be granted, as appropriate for the object.
    - `REVOKE CREATE ON SCHEMA public FROM PUBLIC;`

    - system catalog schema : pg_catalog : It is always inside the search path at first. because built in types must be findable.

    - TODO: usage pattern of schema


### Inheritance

- Postgresql implements table inheritance, very useful for database designers.
- for eg. we have data model for cities, now we want to have a data model for capitals of state, then this should implement cities. What doing this ? because we want to quickly retrive capital data of state.
- for eg. 
    - 
    ```
    CREATE TABLE cities (
    name            text,
    population      float,
    elevation       int     -- in feet
    );

    CREATE TABLE capitals (
        state           char(2)
    ) INHERITS (cities);
    ```
- a table can inherit from 0 or more tables and a query can reference either all rows of a table or all rows of a table plus all of its descendant tables.

```
// finds name of all the cities including state capitals
SELECT name, elevation
    FROM cities
    WHERE elevation > 500;

// while this query find cities except the capitals
SELECT name, elevation
    FROM ONLY cities // we can also write cities* to specify all the decedents of it.// it is default behaviour.
    WHERE elevation > 500;
```
- Here the ONLY keyword indicates that the query should apply only to cities, and not any tables below cities in the inheritance hierarchy. Many of the commands that we have already discussed — SELECT, UPDATE and DELETE — support the ONLY keyword.
- tableoid field in tables can tell you the originating table.
- All check constraints and not-null constraints on a parent table are automatically inherited by its children, unless explicitly specified otherwise with NO INHERIT clauses. Other types of constraints (unique, primary key, and foreign key constraints) are not inherited.
- we can not do this
    - ```
    INSERT INTO cities (name, population, elevation, state)
    VALUES ('Albany', NULL, NULL, 'NY');
    ```
    - as cities don't have elevantion field.

- A parent table cannot be dropped while any of its children remain. Neither can columns or check constraints of child tables be dropped or altered if they are inherited from any parent tables. If you wish to remove a table and all of its descendants, one easy way is to drop the parent table with the CASCADE
- ALTER TABLE will propagate any changes in column data definitions and check constraints down the inheritance hierarchy. 
- dropping columns only possible by using cascade

### Partitioning


### Other database objects

Views

Functions, procedures, and operators

Data types and domains

Triggers and rewrite rules

### Dependency Tracking



## Data Manipulation
TODO: know but write docs here.
### Inserting Data
### Updating Data
### Deleting Data
### Returning Data From Modified Rows


## Query

```
SELECT a, b + c FROM table1;

SELECT random(); // for random number between 0 and 1

// syntax
[WITH with_queries] SELECT select_list FROM table_expression [sort_specification]
```

### table expression

- A table expression computes a table. The table expression contains a FROM clause that is optionally followed by WHERE, GROUP BY, and HAVING clauses

- The FROM clause derives a table from one or more other tables given in a comma-separated table reference list. `FROM table_reference [, table_reference [, ...]]`

#### joined tables

- A joined table is a table derived from two other (real or derived) tables according to the rules of the particular join type. Inner, outer, and cross-joins are available. The general syntax of a joined table is
    - `T1 join_type T2 [ join_condition ]`
    - join chaining can be done
    - Parentheses can be used around JOIN clauses to control the join order. In the absence of parentheses, JOIN clauses nest left-to-right.

##### Join Types

- CROSS JOIN
    - equivalent to FROM T1, T2.
    - `T1 CROSS JOIN T2`

- Qualified Joins
    ```
    T1 { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN T2 ON boolean_expression
    T1 { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN T2 USING ( join column list )
    T1 NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN T2

    // The words INNER and OUTER are optional in all forms. INNER is the default; LEFT, RIGHT, and FULL imply an outer join.
    // The join condition is specified in the ON or USING clause, or implicitly by the word NATURAL.
    ```
    - INNER JOIN
        - 
    - LEFT OUTER JOIN
    - RIGHT OUTER JOIN
    - FULL OUTER JOIN

    - The ON clause is the most general kind of join condition: it takes a Boolean value expression of the same kind as is used in a WHERE clause. A pair of rows from T1 and T2 match if the ON expression evaluates to true.

    - The USING clause is a shorthand that allows you to take advantage of the specific situation where both sides of the join use the same name for the joining column(s). It takes a comma-separated list of the shared column names and forms a join condition that includes an equality comparison for each one. For example, joining T1 and T2 with USING (a, b) produces the join condition ON T1.a = T2.a AND T1.b = T2.b.
    - NATURAL is shorthand for USING.


##### Table and column aliases

- you can give aliases to your table name and column names.

##### Subqueries

- Subqueries specifying a derived table must be enclosed in parentheses. They may be assigned a table alias name, and optionally column alias names.
- `FROM (SELECT * FROM table1) AS alias_name;`

##### Table functions

- Table functions are functions that produce a set of rows, made up of either base data types (scalar types) or composite data types (table rows). They are used like a table, view, or subquery in the FROM clause of a query. Columns returned by table functions can be included in SELECT, JOIN, or WHERE clauses in the same manner as columns of a table, view, or subquery.
- 

TODO: table functions

##### Lateral subqueries
##### WHERE clause
##### GROUP BY and HAVING clause
##### Select Lists, Column Lables, DISTINCT


##### Combining queries [union, intersect, except]

```
query1 UNION [ALL] query2
query1 INTERSECT [ALL] query2
query1 EXCEPT [ALL] query2
```

- the two queries must be “union compatible”, which means that they return the same number of columns and the corresponding columns have compatible data types

- UNION effectively appends the result of query2 to the result of query1 (although there is no guarantee that this is the order in which the rows are actually returned). Furthermore, it eliminates duplicate rows from its result, in the same way as DISTINCT, unless UNION ALL is used.
- INTERSECT returns all rows that are both in the result of query1 and in the result of query2. Duplicate rows are eliminated unless INTERSECT ALL is used.
- EXCEPT returns all rows that are in the result of query1 but not in the result of query2. (This is sometimes called the difference between two queries.) Again, duplicates are eliminated unless EXCEPT ALL is used.

##### Sorting rows (ORDER BY)

##### LIMIT and OFFSET

- user ORDER BY to get consistent output with this.

##### VALUES Lists

- way to create a constant table that can be used in query without having to actually create and populate table on disk.
- `VALUES ( expression [, ...] ) [, ...]`
- `SELECT * FROM (VALUES (1, 'one'), (2, 'two'), (3, 'three')) AS t (num,letter);`
- 


##### WITH queries [ Common Table Expressions ]

- WITH provides a way to write auxiliary statements for use in a larger query. These statements, which are often referred to as Common Table Expressions or CTEs, can be thought of as defining temporary tables that exist just for one query. Each auxiliary statement in a WITH clause can be a SELECT, INSERT, UPDATE, or DELETE; and the WITH clause itself is attached to a primary statement that can be a SELECT, INSERT, UPDATE, DELETE, or MERGE.

###### SELECT in WITH

The basic value of SELECT in WITH is to break down complicated queries into simpler parts. 

example :

```
WITH regional_sales AS (
    SELECT region, SUM(amount) AS total_sales
    FROM orders
    GROUP BY region
), top_regions AS (
    SELECT region
    FROM regional_sales
    WHERE total_sales > (SELECT SUM(total_sales)/10 FROM regional_sales)
)
SELECT region,
       product,
       SUM(quantity) AS product_units,
       SUM(amount) AS product_sales
FROM orders
WHERE region IN (SELECT region FROM top_regions)
GROUP BY region, product;
```

This is just to make the life simpler to write sql queries.

###### Recursive queries

The optional RECURSIVE modifier changes WITH from a mere syntactic convenience into a feature that accomplishes things not otherwise possible in standard SQL. Using RECURSIVE, a WITH query can refer to its own output. A very simple example is this query to sum the integers from 1 through 100:

Recursive queries are typically used to deal with hierarchical or tree-structured data. A useful example is this query to find all the direct and indirect sub-parts of a product, given only a table that shows immediate inclusions:

```
WITH RECURSIVE factorial(n) AS(
	VALUES(1) // initially value 1
  
  UNION // here you can do UNION ALL to also take duplicate values
  SELECT n + 1 FROM factorial WHERE n < 100 // when this condition fails we are out of this loop.
--   now recursive
  
)

SELECT * from factorial;
```

Another good example of heirarchy.

```
WITH RECURSIVE EmployeeHierarchy AS (
    -- Anchor query
    SELECT employee_id, employee_name, manager_id
    FROM employees
    WHERE manager_id = 123  -- Starting manager ID
    
    UNION ALL
    
    -- Recursive query
    SELECT e.employee_id, e.employee_name, e.manager_id
    FROM employees e
    INNER JOIN EmployeeHierarchy eh ON eh.employee_id = e.manager_id
)
SELECT employee_id, employee_name, manager_id
FROM EmployeeHierarchy;
```


###### Search Order

- When computing a tree traversal using a recursive query, you might want to order the results in either depth-first or breadth-first order. This can be done by computing an ordering column alongside the other data columns and using that to sort the results at the end. Note that this does not actually control in which order the query evaluation visits the rows; that is as always in SQL implementation-dependent. This approach merely provides a convenient way to order the results afterwards.


###### TODO:Cycle detection
###### TODO:Common table expression materialization
###### TODO:Data modifying statements in WITH


# Data types

## Numeric Types [ 8.1 ]

- Numeric types consist of two-, four-, and eight-byte integers, four- and eight-byte floating-point numbers, and selectable-precision decimals.

- Types
    - smallint - 2
    - integer - 4
    - bigint - 8 
    - decimal  - variable
    - numeric - variable
    - real - 4
    - double precision - 8
    - smallserial - 2
    - serial - 4
    - bigserial - 8 - large autoincrementing integer

- The types smallint, integer, and bigint store whole numbers.

### Arbitrary precision numbers

- numeric type used when required most accuracy. reccommended when storing monetory information.
- calculations on numeric is very slow compared to integer types or floating point.
- SOME TERMINOLOGY : The precision of a numeric is the total count of significant digits in the whole number, that is, the number of digits to both sides of the decimal point. The scale of a numeric is the count of decimal digits in the fractional part, to the right of the decimal point. So the number 23.5141 has a precision of 6 and a scale of 4. Integers can be considered to have a scale of zero.
- max precision and max scale can be configured.
- NUMERIC(precision, scale) or NUMERIC(precision), NUMERIC
- Always wew should explicitly define precision and scale..

- NUMERIC(3, 1) it will round values to 1 decimal point and can store between -99.9 and 99.9, inclusive.
- type decimal and numeric are equivalent.
- When rounding values, the numeric type rounds ties away from zero, while (on most machines) the real and double precision types round ties to the nearest even number.

## FLoating point types


## Serial types

- smallserial, serial, bigserial are not types just a notational convenience for creating unique identifier column.

```
CREATE TABLE tablename (
    colname SERIAL
);
is equivalent to specifying:

CREATE SEQUENCE tablename_colname_seq AS integer;
CREATE TABLE tablename (
    colname integer NOT NULL DEFAULT nextval('tablename_colname_seq')
);
ALTER SEQUENCE tablename_colname_seq OWNED BY tablename.colname;
```

## Monetory types

The money type stores a currency amount with a fixed fractional precision;


## Boolean Type

boolean	1 byte	state of true or false

```
CREATE TABLE test1 (a boolean, b text);
INSERT INTO test1 VALUES (TRUE, 'sic est');

'yes'::boolean
```

## Enumerated Types

Enumerated (enum) types are data types that comprise a static, ordered set of values. 

### Declaration

```
CREATE TYPE mood AS ENUM ('sad', 'ok', 'happy');

CREATE TABLE person (
    name text,
    current_mood mood
);
INSERT INTO person VALUES ('Moe', 'happy');
SELECT * FROM person WHERE current_mood = 'happy';
```

### Ordering of values

According to the sequence in which they are listed. if order by 'enumType' then it will follow that sequence.

### Type safety

- each enum data type is different and can not be comparable with other enum types.
- Enum labels are case sensitive...

### Geometric Types

- Points
- Lines
- Line segments
- Boxes 
- Paths
- Polygons
- Circles

### Network Address Types [ looking good if want to work in this ]


```
// This is not valid if mood and happiness are enums
SELECT person.name, holidays.num_weeks FROM person, holidays
  WHERE person.current_mood = holidays.happiness;

// but if we really want to compare the underlying text then we can do explicit casting
SELECT person.name, holidays.num_weeks FROM person, holidays
  WHERE person.current_mood::text = holidays.happiness::text;

```

## Mathematical functions and operators [9.3]

### Mathematical Operators

```
numeric_type + numeric_type → numeric_type
+ numeric_type → numeric_type
numeric ^ numeric → numeric
|/ double precision → double precision // Square Root
||/ double precision → double precision // Cube root
@ numeric_type → numeric_type // Absolute value
integral_type & integral_type → integral_type // Bitwise AND
integral_type | integral_type → integral_type // Bitwise OR
integral_type # integral_type → integral_type // Bitwise Exclusive OR
~ integral_type → integral_type // Bitwise NOT
integral_type << integer → integral_type // Bitwise Shift Left
integral_type >> integer → integral_type // Bitwise Shift Right
```

### Mathematical Functions

```
abs ( numeric_type ) → numeric_type // absolute value
cbrt ( double precision ) → double precision // cube root
ceil ( numeric ) → numeric // Nearest integer greater than or equal to argument
ceil ( double precision ) → double precision
degrees ( double precision ) → double precision // convert radians to degrees
div ( y numeric, x numeric ) → numeric // 9/2 == 4 // truncates towards 0

and many more functions like exp, power, pi(), log, lcm, gcd.

IMPORTANT : it requires numeric type hence we need to cast floating point numbers to this numeric types in order to use this function.
round ( v numeric, s integer ) → numeric  // Rounds v to s decimal places. Ties are broken by rounding away from zero. // for eg. round(42.4382, 2) → 42.44
```

## Date Time Types

date - 4 bytes - resolution of 1 day
interval [fields] [(p)] - 16 bytes - time interval - resolution of 1 micro second
time [ (p) ] [ without time zone ] - 8 bytes - time of day ( no date )
time [ (p) ] with time zone - 12 bytes - time with timezone ( no date)
timestamp [ (p) ] [ without time zone ] - both date and time
similarly with tiimezone 

-> time, timestamp and interval accepts optional precison value p which specifies no of fractional digits retained in seconds field. The allowed range of p is from 0 to 6.

->Interval has additional option given below.

    -YEAR
    -MONTH
    -DAY
    -HOUR
    -MINUTE
    -SECOND
    -YEAR TO MONTH
    -DAY TO HOUR
    -DAY TO MINUTE
    -DAY TO SECOND
    -HOUR TO MINUTE
    -HOUR TO SECOND
    -MINUTE TO SECOND

if both fields and p are given then fields can only be seconds as precision only applies to seconds.

- Remember that any date or time literal input needs to be enclosed in single quotes, like text strings.

- Possible Inputs for date field : https://www.postgresql.org/docs/current/datatype-datetime.html#DATATYPE-DATETIME-DATE-TABLE


- CURRENT_DATE : to get current date.

- The date/time style can be selected by the user using the SET datestyle command, the DateStyle parameter in the postgresql.conf configuration file, or the PGDATESTYLE environment variable on the server or client.

### Interval input

- Syntax `[@] quantity unit [quantity unit...] [direction]`

- Quantity is a number
- unit is microsecond, second, minute, day, week, month, year etc.
- direction can be `ago` or empty. Ago negates whole date.


### Date time functions and operators

REF : https://www.postgresql.org/docs/current/functions-datetime.html

- Dates and timestamp are all comparable.
- times and intervals can only be compared with their type.

- Date time operators
    - date + integer = date
    - date + interval → timestamp
    - date + time → timestamp
    - interval + interval → interval
    - timestamp + interval → timestamp
    - time + interval → time
    - interval → interval
    - date - date → integer
    - date - integer → date
    - date - interval → timestamp
    - time - time → interval
    - time - interval → time
    - time '05:00' - interval '2 hours' → 03:00:00
    - timestamp - interval → timestamp
    - interval - interval → interval
    - timestamp - timestamp → interval
    - interval * double precision → interval
    - Multiply an interval by a scalar
    - interval '1 second' * 900 → 00:15:00
    - interval '1 day' * 21 → 21 days
    - interval '1 hour' * 3.5 → 03:30:00
    - interval / double precision → interval
    - Divide an interval by a scalar
    - interval '1 hour' / 1.5 → 00:40:00


### Date time functions

- view inside sql practice folder.

### 9.18 Conditional Expressions

#### Case

```
CASE WHEN condition THEN result
     [WHEN ...]
     [ELSE result]
END

example :


SELECT a,
       CASE WHEN a=1 THEN 'one'
            WHEN a=2 THEN 'two'
            ELSE 'other'
       END
    FROM test;
```

# Functions And Operators [9]

## Logical Operators [9.1]

```
boolean AND boolean → boolean
boolean OR boolean → boolean
NOT boolean → boolean
```

## Comparison Functions And Operators [9.2]

-  <> is the standard SQL notation for “not equal”. != is an alias, which is converted to <> at a very early stage of parsing. Hence, it is not possible to implement != and <> operators that do different things.

- Normal Operators : <>, !=, =, >, <, <=, >= etc.

### Comparison Predicates

```
datatype BETWEEN datatype AND datatype → boolean // 2 BETWEEN 1 AND 3 returns TRUE
datatype NOT BETWEEN datatype AND datatype → boolean // negation of BETWEEN
datatype BETWEEN SYMMETRIC datatype AND datatype → boolean // between after sorting the two endpoint values
same for negation

datatype IS DISTINCT FROM datatype → boolean // Not equal, treating null as a comparable value.
for eg. 1 IS DISTINCT FROM NULL → t (rather than NULL) and NULL IS DISTINCT FROM NULL → f (rather than NULL)

datatype IS NULL → boolean
datatype IS NOT NULL → boolean

boolean IS TRUE → boolean and negation

boolean IS UNKNOWN → boolean // for eg. NULL::boolean IS UNKNOWN → t (rather than NULL)
```

## Mathematical Functions and Operators [9.3]

Refere Previous Notes

## String Functions and Operators

Strings in this context include values of the types character, character varying, and text. They will also accept character varying argument which will be converted to text before applying function which can remove trailing spaces in character value.

```
text || text → text // Concatination
text || anynonarray → text // for eg. "good" || 24 -> "good24"  
anynonarray || text → text


-- Returns number of characters in string
char_length ( text ) → integer
character_length ( text ) → integer

-- lower case and upper case
lower ( text ) → text
upper ( text ) → text


```


## Indexes

```
CREATE TABLE test1 (
    id integer,
    content varchar
);

SELECT content FROM test1 WHERE id = constant;
```

- With no advance preparation, the system would have to scan the entire test1 table, row by row, to find all matching entries. If there are many rows in test1 and only a few rows (perhaps zero or one) that would be returned by such a query, this is clearly an inefficient method. But if the system has been instructed to maintain an index on the id column, it can use a more efficient method for locating matching rows. For instance, it might only have to walk a few levels deep into a search tree.

- To create index `CREATE INDEX test1_id_index ON test1 (id);`
- To remove index `DROP INDEX`
- But you might have to run the ANALYZE command regularly to update statistics to allow the query planner to make educated decisions. it automatically plans the query according to requirements after the index is set.
- Indexes can also benefit UPDATE and DELETE commands with search conditions. Indexes can moreover be used in join searches. Thus, an index defined on a column that is part of a join condition can also significantly speed up queries with joins


### Index Types

- Lots of index algorithms
- By default, the CREATE INDEX command creates B-tree indexes
- To create using particular algorithm `CREATE INDEX name ON table USING HASH (column);`

