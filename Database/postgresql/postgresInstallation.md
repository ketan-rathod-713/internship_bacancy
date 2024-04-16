# Postgresql Installation

service postgresql // list all the commands available

service postgresql start | status | restart etc.

Default user : postgres

sudo su postgres // going inside root user postgres

// There is a command line tool for the postgresql which is psql

then psql to enter into posgress 

\l - list all the databases

\du - list out database users

For changing password of the default user

ALTER USER postgres WITH PASSWORD '<password-here>'; // dont forget semicolon heree ha ha

// Creating New User
CREATE USER <userName> WITH PASSWORD '<password-here>';

\du // See user1 will be added to list but there are no priviledges to this user.

ALTER USER username WITH SUPERUSER; // haha

// 
man psql // to check for manual 

PG Admin Client
- Postgres Adminstration Tool on software

# Postgresql Short Notes

- By default a user is created postgres. Hence to access psql we need to type # sudo su postgres
- Then we can access psql shell

## PSQL
- For writting sql queries
- Other commands starts from \ for eg. \h for help and \q for quit.

### Important Cheatsheet Working With Users and Databases

SELECT current_user;  -- user name of current execution context
SELECT session_user;  -- session user name

createdb mydb // bahar se database banane ke liye

FOR PSQL
CREATE DATABASE name;

