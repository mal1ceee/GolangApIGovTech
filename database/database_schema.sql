-- Create Teachers Table
CREATE TABLE Teachers (
    teacher_id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL
);

-- Create Students Table
CREATE TABLE Students (
    student_id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    is_suspended BOOLEAN DEFAULT FALSE
);

-- Create Registrations Table
CREATE TABLE Registrations (
    registration_id SERIAL PRIMARY KEY,
    teacher_id INT NOT NULL,
    student_id INT NOT NULL,
    FOREIGN KEY (teacher_id) REFERENCES Teachers(teacher_id) ON DELETE CASCADE,
    FOREIGN KEY (student_id) REFERENCES Students(student_id) ON DELETE CASCADE,
    UNIQUE (teacher_id, student_id)
);
