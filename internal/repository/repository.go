package repository

import (
	"GOLANGAPIGOVTECH/internal/model"

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
	err := r.db.Get(&teacherID, "SELECT id FROM teachers WHERE email = $1", teacherEmail)
	if err != nil {
		return err
	}

	for _, email := range studentEmails {
		var studentID int
		err := r.db.Get(&studentID, "SELECT id FROM students WHERE email = $1", email)
		if err != nil {
			return err // Or continue to attempt to register other students
		}

		_, err = r.db.Exec("INSERT INTO registrations (teacher_id, student_id) VALUES ($1, $2)", teacherID, studentID)
		if err != nil {
			return err // Or log this error and continue
		}
	}
	return nil
}

// GetCommonStudents retrieves students who are registered to all the given teachers.
// Adjust this function based on your specific SQL schema and logic requirements.
func (r *Repository) GetCommonStudents(teacherEmails []string) ([]model.Student, error) {
	var students []model.Student

	// Construct a query to count the number of unique teachers per student
	// that matches the list of provided teacher emails.
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

	// sqlx.In returns a query with '?' placeholders, we need to rebind it for our database
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
// This includes students who are not suspended and are either registered to the teacher or mentioned in the notification.
func (r *Repository) GetStudentsForNotifications(teacherEmail string, mentionedEmails []string) ([]model.Student, error) {
	var students []model.Student

	// Step 1: Find the teacher's ID based on their email.
	var teacherID int
	err := r.db.Get(&teacherID, "SELECT teacher_id FROM teachers WHERE email = $1", teacherEmail)
	if err != nil {
		return nil, err
	}

	// Step 2: Query to select students registered to the teacher and not suspended.
	registeredStudentsQuery := `
    SELECT DISTINCT s.student_id, s.email
    FROM students s
    JOIN registrations r ON s.student_id = r.student_id
    WHERE r.teacher_id = ? AND s.is_suspended = FALSE
    `

	err = r.db.Select(&students, registeredStudentsQuery, teacherID)
	if err != nil {
		return nil, err
	}

	// Step 3: Add mentioned students who are not suspended, if they are not already included.
	// This requires checking each mentioned email to see if it's not in the already selected students' list.
	// Then, querying for each mentioned student that is not suspended and adding them to the list.
	for _, email := range mentionedEmails {
		var mentionedStudent model.Student
		mentionedStudentQuery := `
        SELECT student_id, email
        FROM students
        WHERE email = ? AND is_suspended = FALSE
        `

		err := r.db.Get(&mentionedStudent, mentionedStudentQuery, email)
		// If there's no error, we found a student by the mentioned email who is not suspended.
		if err == nil {
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
			}
		}

	}

	return students, nil
}
