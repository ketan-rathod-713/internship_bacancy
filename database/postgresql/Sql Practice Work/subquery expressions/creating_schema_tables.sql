

-- create registrations referencing students and subjects

-- create students

-- create subjects

-- CREATE SCHEMA subquery;

-- CREATE TABLE subquery.students(
--     id SERIAL PRIMARY KEY,
--     name varchar(255),
--     email varchar(255),
--     phone varchar(255)  
-- );


-- CREATE TABLE subquery.subjects(
--     id SERIAL PRIMARY KEY,
--     name varchar(255) UNIQUE,
--     startDate date,
--     endDate date
-- );

-- INSERT INTO subquery.subjects(name) VALUES('Maths') RETURNING *;
-- INSERT INTO subquery.subjects(name) VALUES('Physics') RETURNING *;
-- INSERT INTO subquery.subjects(name) VALUES('Chemistry') RETURNING *;
-- INSERT INTO subquery.subjects(name) VALUES('English') RETURNING *;


-- INSERT INTO subquery.students(name) VALUES('tridip') RETURNING *;
-- INSERT INTO subquery.students(name) VALUES('manav') RETURNING *;
-- INSERT INTO subquery.students(name) VALUES('vatsal') RETURNING *;
-- INSERT INTO subquery.students(name) VALUES('ketan') RETURNING *;

-- CREATE TABLE IF NOT EXISTS subquery.registrations(
--     id SERIAL PRIMARY KEY,
--     student_id INT NOT NULl,
--     subject_id INT NOT NULL,
--     CONSTRAINT uniqueness UNIQUE(student_id, subject_id)
-- );

-- INSERT INTO subquery.registrations(student_id, subject_id) VALUES(1, 1);
-- INSERT INTO subquery.registrations(student_id, subject_id) VALUES(1, 2);
-- INSERT INTO subquery.registrations(student_id, subject_id) VALUES(1, 3);
-- INSERT INTO subquery.registrations(student_id, subject_id) VALUES(2, 3);
-- INSERT INTO subquery.registrations(student_id, subject_id) VALUES(3, 3);
-- INSERT INTO subquery.registrations(student_id, subject_id) VALUES(2, 1);