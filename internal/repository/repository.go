package repository

import (
	"GOLANGAPIGOVTECH/internal/model"
	"log"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

// RegisterStudents links students to a teacher, creating new entries in the Registration table.
func (r *Repository) RegisterStudents(teacherEmail string, studentEmails []string) error {
	var teacherID int
	// Log the attempt to find the teacher's ID
	log.Printf("Attempting to find teacher ID for email: %s", teacherEmail)
	err := r.db.Get(&teacherID, "SELECT teacher_id FROM teachers WHERE email = $1", teacherEmail)
	if err != nil {
		// Log the error if the teacher's ID couldn't be found
		log.Printf("Error finding teacher ID for email %s: %v", teacherEmail, err)
		return err
	}
	// Log the found teacher's ID
	log.Printf("Found teacher ID: %d for email: %s", teacherID, teacherEmail)

	for _, email := range studentEmails {
		var studentID int
		// Log the attempt to find each student's ID
		log.Printf("Attempting to find student ID for email: %s", email)
		err := r.db.Get(&studentID, "SELECT student_id FROM students WHERE email = $1", email)
		if err != nil {
			// Log the error if a student's ID couldn't be found
			log.Printf("Error finding student ID for email %s: %v", email, err)
			return err // Or consider logging and continuing with the next student
		}
		// Log the found student's ID
		log.Printf("Found student ID: %d for email: %s", studentID, email)

		// Log the attempt to register the student to the teacher
		log.Printf("Attempting to register student %d to teacher %d", studentID, teacherID)
		_, err = r.db.Exec("INSERT INTO registrations (teacher_id, student_id) VALUES ($1, $2) ON CONFLICT (teacher_id, student_id) DO NOTHING", teacherID, studentID)
		if err != nil {
			log.Printf("Unexpected error during registration: %v", err)
			return err
		}
		// Log successful registration
		log.Printf("Successfully registered student %s to teacher %s", email, teacherEmail)
	}
	return nil
}

// GetCommonStudents retrieves students who are registered to all the given teachers.

func (r *Repository) GetCommonStudents(teacherEmails []string) ([]model.Student, error) {
	var students []model.Student

	
	// Only students with a count equal to the number of provided teacher emails are selected.
	query := `
    SELECT s.student_id, s.email
    FROM students s
    JOIN registrations r ON s.student_id = r.student_id
    JOIN teachers t ON r.teacher_id = t.teacher_id
    WHERE t.email IN (?)
    GROUP BY s.student_id
    HAVING COUNT(DISTINCT t.teacher_id) = ?
    `

	// Use sqlx.In to handle IN query with a dynamic number of arguments
	query, args, err := sqlx.In(query, teacherEmails, len(teacherEmails))
	if err != nil {
		return nil, err
	}

	
	query = r.db.Rebind(query)

	// Execute the query
	err = r.db.Select(&students, query, args...)
	if err != nil {
		return nil, err
	}

	return students, nil
}

// SuspendStudent updates a student's is_suspended status to true.
func (r *Repository) SuspendStudent(studentEmail string) error {
	_, err := r.db.Exec("UPDATE students SET is_suspended = $1 WHERE email = $2", true, studentEmail)
	return err
}

// GetStudentsForNotifications retrieves students who are eligible to receive a notification from a teacher.

func (r *Repository) GetStudentsForNotifications(teacherEmail string, mentionedEmails []string) ([]model.Student, error) {
	var students []model.Student

	log.Printf("Getting students for notifications for teacher: %s with mentioned emails: %v", teacherEmail, mentionedEmails)

	
	var teacherID int
	err := r.db.Get(&teacherID, "SELECT teacher_id FROM teachers WHERE email = $1", teacherEmail)
	if err != nil {
		log.Printf("Failed to find teacher ID for email %s: %v", teacherEmail, err)
		return nil, err
	}

	log.Printf("Found teacher ID: %d", teacherID)


	registeredStudentsQuery := `
    SELECT DISTINCT s.student_id, s.email
    FROM students s
    JOIN registrations r ON s.student_id = r.student_id
    WHERE r.teacher_id = $1 AND s.is_suspended = FALSE
    `

	err = r.db.Select(&students, registeredStudentsQuery, teacherID)
	if err != nil {
		log.Printf("Failed to retrieve registered students for teacher ID %d: %v", teacherID, err)
		return nil, err
	}

	log.Printf("Retrieved registered students: %v", students)

	
	for _, email := range mentionedEmails {
		var mentionedStudent model.Student
		mentionedStudentQuery := `
        SELECT student_id, email
        FROM students
        WHERE email = $1 AND is_suspended = FALSE
        `

		err := r.db.Get(&mentionedStudent, mentionedStudentQuery, email)
		if err != nil {
			log.Printf("Failed to retrieve mentioned student with email %s: %v", email, err)
			
			continue 
		}

		log.Printf("Found mentioned student: %s", mentionedStudent.Email)

		// Check if this student is already in the students slice to avoid duplicates.
		alreadyIncluded := false
		for _, stu := range students {
			if stu.Email == email {
				alreadyIncluded = true
				break
			}
		}
		if !alreadyIncluded {
			students = append(students, mentionedStudent)
			log.Printf("Added mentioned student to the list: %s", mentionedStudent.Email)
		}
	}

	return students, nil
}
