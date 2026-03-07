package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

func (h *Handler) CreateBlog(c *gin.Context) {
	var newBlog Blog
	err := c.ShouldBindJSON(&newBlog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found"})
		return
	}
	userIDstr := strconv.Itoa(userID.(int))

	_, err = h.DB.Exec(c.Request.Context(), `
		INSERT INTO posts (author_id, title, content, category, tags)
		VALUES ($1, $2, $3, $4, $5)
	`, userIDstr, newBlog.Title, newBlog.Content, newBlog.Category, newBlog.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create blog: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "post created successfully"})
}
