CREATE TABLE IF NOT EXISTS students (
    id int PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age int NOT NULL,
    course VARCHAR(50) NOT NULL
);

INSERT INTO students(id, name, age, course)
VALUES (1, 'ketan', 20, 'computer');

INSERT INTO students(id, name, age, course)
VALUES (2, 'tridip', 26, 'computer');

INSERT INTO students(id, name, age, course)
VALUES (3, 'darsh joshi', 24, 'information technology');
