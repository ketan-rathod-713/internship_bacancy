
What to learn in dbms

1. Mongodb
    - see notes
    - documentation for crud operations.

2. Postgressql
    - see notes of sql and commands done
    - see joins and aggregations and subquery commands

See mongo and sql package.
'

DDL(Data Definition Language):  It contains commands which are required to define the database.
E.g., CREATE, ALTER, DROP, TRUNCATE, RENAME, etc.
DML(Data Manipulation Language): It contains commands which are required to manipulate the data present in the database.
E.g., SELECT, UPDATE, INSERT, DELETE, etc.
DCL(Data Control Language):  It contains commands which are required to deal with the user permissions and controls of the database system.
E.g., GRANT and REVOKE.
TCL(Transaction Control Language):  It contains commands which are required to deal with the transaction of the database.
E.g., COMMIT, ROLLBACK, and SAVEPOINT.

Referencing table jo foreign key references kar raha he
referenced table which has pk.. which is being referenced by referencing table.

Normalization
1. duplicate ddata remove only atomic
2. if we delete primary key data then non prime data should be retained. hence divide tables accordingly.
    - Every non-prime attribute of the table should be fully functionally dependent on the primary key
    - else create new table for it.
3. There is no transitive functional dependency of one attribute on any attribute in the same table.
    - if saluatations there Mr and MRs. then we should create ites table and refer its id ffrom there..

4. BCNF - For every functional dependency of any attribute A on B
(A->B), A should be the super key of the table. It simply implies that A can’t be a non-prime attribute if B is a prime attribute.


Keys

candidate key - set of properties that uniquely identifies a table.
super key - set of attributes that u same
Primary key - unique not null
unique key
alternate key  - not choosen pk
 The foreign key defines an attribute that can only take the values present in one table common to the attribute present in another table. In the below example courseId from the Student table is a foreign key to the Course table, as both, the tables contain courseId as one of their attributes.