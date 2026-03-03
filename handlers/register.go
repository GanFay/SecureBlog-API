package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Register struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Register(c *gin.Context) {
	var RegisterInfo Register
	err := c.ShouldBindJSON(&RegisterInfo)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(RegisterInfo)
}
