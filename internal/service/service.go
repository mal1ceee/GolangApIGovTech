package service

import (
	"GOLANGAPIGOVTECH/internal/model"
	"GOLANGAPIGOVTECH/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

// NewService creates a new instance of Service with a given repository.
func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

// RegisterStudents links a list of students to a specific teacher.
func (s *Service) RegisterStudents(teacherEmail string, studentEmails []string) error {
	return s.repo.RegisterStudents(teacherEmail, studentEmails)
}

// GetCommonStudents retrieves a list of students registered to all provided teachers.
func (s *Service) GetCommonStudents(teacherEmails []string) ([]model.Student, error) {
	return s.repo.GetCommonStudents(teacherEmails)
}

// SuspendStudent marks a specific student as suspended.
func (s *Service) SuspendStudent(studentEmail string) error {
	return s.repo.SuspendStudent(studentEmail)
}

// GetStudentsForNotifications finds students eligible to receive a given notification.
func (s *Service) GetStudentsForNotifications(teacherEmail string, mentionedEmails []string) ([]model.Student, error) {
	return s.repo.GetStudentsForNotifications(teacherEmail, mentionedEmails)
}
