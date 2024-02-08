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

