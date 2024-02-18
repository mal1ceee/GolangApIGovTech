-- Insert seed data into Teachers
INSERT INTO teachers (email) VALUES ('teacherken@gmail.com');
INSERT INTO teachers (email) VALUES ('teacherjoe@gmail.com');

-- Insert seed data into Students
INSERT INTO students (email, is_suspended) VALUES ('studentjon@gmail.com', FALSE);
INSERT INTO students (email, is_suspended) VALUES ('studenthon@gmail.com', FALSE);
INSERT INTO students (email, is_suspended) VALUES ('commonstudent1@gmail.com', FALSE);
INSERT INTO students (email, is_suspended) VALUES ('commonstudent2@gmail.com', FALSE);
INSERT INTO students (email, is_suspended) VALUES ('studentmary@gmail.com', FALSE);
INSERT INTO students (email, is_suspended) VALUES ('studentagnes@gmail.com', FALSE);
INSERT INTO students (email, is_suspended) VALUES ('studentmiche@gmail.com', FALSE);

-- Register students to Teacher Ken
INSERT INTO registrations (teacher_id, student_id) SELECT teacher_id, student_id FROM teachers t, students s WHERE t.email = 'teacherken@gmail.com' AND s.email IN ('studentjon@gmail.com', 'studenthon@gmail.com', 'commonstudent1@gmail.com', 'commonstudent2@gmail.com');

-- Register students to Teacher Joe
INSERT INTO registrations (teacher_id, student_id) SELECT teacher_id, student_id FROM teachers t, students s WHERE t.email = 'teacherjoe@gmail.com' AND s.email IN ('commonstudent1@gmail.com', 'commonstudent2@gmail.com', 'studentmary@gmail.com');





