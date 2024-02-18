package repository

import (
	
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	
	"github.com/stretchr/testify/assert"
)

func TestGetCommonStudents(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"student_id", "email", "is_suspended"}).
		AddRow(1, "student1@example.com", false).
		AddRow(2, "student2@gmail.com", false)

	query := `SELECT s.student_id, s.email, s.is_suspended FROM students s JOIN registrations r ON s.student_id = r.student_id JOIN teachers t ON r.teacher_id = t.teacher_id WHERE t.email IN \(\?\) GROUP BY s.student_id HAVING COUNT\(DISTINCT t.teacher_id\) = \?`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)

	repo := NewRepository(sqlxDB)
	teacherEmails := []string{"teacherjoe@gmail.com", "teacherken@gmail.com"}
	students, err := repo.GetCommonStudents(teacherEmails)

	assert.NoError(t, err)
	assert.Len(t, students, 2)
	assert.Equal(t, students[0].Email, "student1@gmail.com")
	assert.Equal(t, students[1].Email, "student2@gmail.com")

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
