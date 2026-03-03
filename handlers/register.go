package handlers

import (
	"blog/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(req.Username) < 4 || len(req.Username) > 32 {
		c.JSON(400, gin.H{"error": "username is too short or too long"})
		return
	} else if len(req.Password) < 5 || len(req.Password) > 128 {
		c.JSON(400, gin.H{"error": "password is too short or too long"})
		return
	} else if len(req.Email) < 6 || len(req.Email) > 256 {
		c.JSON(400, gin.H{"error": "email is too short or too long"})
		return
	}

	hashPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		c.JSON(400, gin.H{"er	ror": err.Error()})
		return
	}

	_, err = h.DB.Exec(c.Request.Context(), `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
	`, req.Username, req.Email, hashPassword)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "register successfully"})
}
