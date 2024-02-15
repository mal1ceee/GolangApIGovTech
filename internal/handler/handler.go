package handler

import (
	"GOLANGAPIGOVTECH/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/students", h.getStudents)
	router.POST("/register", h.registerStudents)
	router.GET("/commonstudents", h.getCommonStudents)
	router.POST("/suspend", h.suspendStudent)
	router.POST("/notifications", h.getStudentsForNotifications)
}

func (h *Handler) getStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the students API!",
	})
}

// Implement additional handlers for your user stories

func (h *Handler) registerStudents(c *gin.Context) {
	// Define a struct to match the expected request body format
	var req struct {
		Teacher  string   `json:"teacher"`
		Students []string `json:"students"`
	}

	// Bind the incoming JSON to the struct
	if err := c.BindJSON(&req); err != nil {
		// If there's an error, return a 400 Bad Request response
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Call the service method with the parsed data
	err := h.svc.RegisterStudents(req.Teacher, req.Students)
	if err != nil {
		// If the service method returns an error, return a 500 Internal Server Error response
		// In a real application, you might want to handle different types of errors differently
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register students"})
		return
	}

	// If successful, return a 204 No Content response
	c.Status(http.StatusNoContent)
}

func (h *Handler) getCommonStudents(c *gin.Context) {
	// Extracting teacher emails from query parameters
	teacherEmails := c.QueryArray("teacher")

	// Check if at least one teacher email is provided
	if len(teacherEmails) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one teacher email must be provided"})
		return
	}

	// Call the service method with the extracted teacher emails
	students, err := h.svc.GetCommonStudents(teacherEmails)
	if err != nil {
		// Handle potential errors, for example, if no common students are found or a database error occurs
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve common students"})
		return
	}

	// If successful, return the list of common students
	c.JSON(http.StatusOK, gin.H{"students": students})
}

func (h *Handler) suspendStudent(c *gin.Context) {
	// Define a struct to match the expected request body format
	var req struct {
		StudentEmail string `json:"student"` // Assuming the payload uses "student" to denote the student's email
	}

	// Bind the incoming JSON to the struct
	if err := c.BindJSON(&req); err != nil {
		// If there's an error in parsing the request, return a 400 Bad Request response
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Validate the extracted email
	if req.StudentEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student email is required"})
		return
	}

	// Call the service method to suspend the student
	err := h.svc.SuspendStudent(req.StudentEmail)
	if err != nil {
		// If suspending the student fails (e.g., student not found or database error), return a 500 Internal Server Error
		// You might also consider more specific error handling to differentiate between not found and server errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to suspend the student"})
		return
	}

	// If successful, return a 204 No Content status as no content needs to be returned
	c.Status(http.StatusNoContent)
}

func (h *Handler) getStudentsForNotifications(c *gin.Context) {
	// Define a struct to match the expected request body format
	var req struct {
		Teacher      string `json:"teacher"`
		Notification string `json:"notification"`
	}

	// Bind the incoming JSON to the struct
	if err := c.BindJSON(&req); err != nil {
		// If there's an error in parsing the request, return a 400 Bad Request response
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Validate the extracted information
	if req.Teacher == "" || req.Notification == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Teacher and notification text are required"})
		return
	}

	// Extract mentioned student emails from the notification text
	// This is a simplified approach; you might need a more robust method for extracting emails
	mentionedEmails := extractMentionedEmails(req.Notification)

	// Call the service method to get eligible students
	students, err := h.svc.GetStudentsForNotifications(req.Teacher, mentionedEmails)
	if err != nil {
		// Handle potential errors, such as database errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students for notifications"})
		return
	}

	// If successful, return the list of students
	c.JSON(http.StatusOK, gin.H{"students": students})
}

// A simplified example function for extracting emails mentioned in the notification text
func extractMentionedEmails(notification string) []string {
	// Placeholder for email extraction logic
	// In a real scenario, you would use regex or similar to find emails in the notification text
	var emails []string
	// TODO: Implement email extraction from the notification text
	return emails
}
