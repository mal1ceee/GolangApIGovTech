package config

type Student struct {
	ID     int    `db:"student_id"`
	Email  string `db:"email"`
	Status bool   `db:"status"`
	// Adjust fields according to your actual table structure
}

type Teacher struct {
	ID    int    `db:"teacher_id"`
	Email string `db:"email"`
}

type Registration struct {
	ID        int `db:"registration_id"`
	StudentID int `db:"student_id"`
	TeacherID int `db:"teacher_id"`
}
