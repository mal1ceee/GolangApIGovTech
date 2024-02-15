my schema code

- Register studentjon and studenthon to teacherken
INSERT INTO Registrations (teacher_id, student_id) VALUES ((SELECT teacher_id FROM Teachers WHERE email = 'teacherken@gmail.com'), (SELECT student_id FROM Students WHERE email = 'studentjon@gmail.com'));
INSERT INTO Registrations (teacher_id, student_id) VALUES ((SELECT teacher_id FROM Teachers WHERE email = 'teacherken@gmail.com'), (SELECT student_id FROM Students WHERE email = 'studenthon@gmail.com'));

-- Register studentmary to teacherjoe
INSERT INTO Registrations (teacher_id, student_id) VALUES ((SELECT teacher_id FROM Teachers WHERE email = 'teacherjoe@gmail.com'), (SELECT student_id FROM Students WHERE email = 'studentmary@gmail.com'));

-- Assume studentagnes and studentmiche are registered to both teachers
INSERT INTO Registrations (teacher_id, student_id) VALUES ((SELECT teacher_id FROM Teachers WHERE email = 'teacherken@gmail.com'), (SELECT student_id FROM Students WHERE email = 'studentagnes@gmail.com'));
INSERT INTO Registrations (teacher_id, student_id) VALUES ((SELECT teacher_id FROM Teachers WHERE email = 'teacherjoe@gmail.com'), (SELECT student_id FROM Students WHERE email = 'studentagnes@gmail.com'));
INSERT INTO Registrations (teacher_id, student_id) VALUES ((SELECT teacher_id FROM Teachers WHERE email = 'teacherken@gmail.com'), (SELECT student_id FROM Students WHERE email = 'studentmiche@gmail.com'));
INSERT INTO Registrations (teacher_id, student_id) VALUES ((SELECT teacher_id FROM Teachers WHERE email = 'teacherjoe@gmail.com'), (SELECT student_id FROM Students WHERE email = 'studentmiche@gmail.com'));


INSERT INTO Students (email, is_suspended) VALUES ('studentjon@gmail.com', FALSE);
INSERT INTO Students (email, is_suspended) VALUES ('studenthon@gmail.com', FALSE);
INSERT INTO Students (email, is_suspended) VALUES ('studentmary@gmail.com', FALSE);
INSERT INTO Students (email, is_suspended) VALUES ('studentagnes@gmail.com', FALSE);
INSERT INTO Students (email, is_suspended) VALUES ('studentmiche@gmail.com', FALSE);

INSERT INTO Teachers (email) VALUES ('teacherken@gmail.com');
INSERT INTO Teachers (email) VALUES ('teacherjoe@gmail.com');
 based on this dummy data, is there any student that is common between the two teacher

CREATE TABLE Teachers (
    teacher_id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE Students (
    student_id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    is_suspended BOOLEAN DEFAULT FALSE
);

CREATE TABLE Registrations (
    registration_id SERIAL PRIMARY KEY,
    teacher_id INT NOT NULL,
    student_id INT NOT NULL,
    CONSTRAINT fk_teacher
        FOREIGN KEY(teacher_id) 
        REFERENCES Teachers(teacher_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_student
        FOREIGN KEY(student_id) 
        REFERENCES Students(student_id)
        ON DELETE CASCADE,
    UNIQUE (teacher_id, student_id)
);


database: pgsql 15 running on port 5433