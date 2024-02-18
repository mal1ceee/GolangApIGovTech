package handler

import (
	"GOLANGAPIGOVTECH/internal/service"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/api/students", h.getStudents)
	router.POST("/api/register", h.registerStudents)
	router.GET("/api/commonstudents", h.getCommonStudents)
	router.POST("/api/suspend", h.suspendStudent)
	router.POST("/api/notifications", h.getStudentsForNotifications)
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
		StudentEmail string `json:"student"`
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

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to suspend the student"})
		return
	}

	// If successful, return a 204 No Content status as no content needs to be returned
	c.Status(http.StatusNoContent)
}

func (h *Handler) getStudentsForNotifications(c *gin.Context) {
	var req struct {
		Teacher      string `json:"teacher"`
		Notification string `json:"notification"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if req.Teacher == "" || req.Notification == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Teacher and notification text are required"})
		return
	}

	mentionedEmails := extractMentionedEmails(req.Notification)

	students, err := h.svc.GetStudentsForNotifications(req.Teacher, mentionedEmails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students for notifications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recipients": students})
}

// extractMentionedEmails extracts email addresses mentioned in the notification text.
func extractMentionedEmails(notification string) []string {
	var emails []string
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}`)
	matches := emailRegex.FindAllString(notification, -1)

	// Ensure unique emails only
	emailMap := make(map[string]bool)
	for _, email := range matches {
		if _, exists := emailMap[email]; !exists {
			emails = append(emails, email)
			emailMap[email] = true
		}
	}

	return emails
}
