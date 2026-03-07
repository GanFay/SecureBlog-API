package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteBlog(c *gin.Context) {

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found"})
		return
	}
	var post Post
	id := c.Param("id")

	err := h.DB.QueryRow(c.Request.Context(), `SELECT * FROM posts WHERE id=$1`, id).Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.Category, &post.Tags, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	authorID, err := strconv.Atoi(post.AuthorID)
	if err != nil {
		return
	}
	if authorID != userID.(int) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not permission"})
		return
	}

	cmdTag, err := h.DB.Exec(c.Request.Context(), `DELETE FROM posts WHERE id=$1`, id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if cmdTag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	c.JSON(204, gin.H{"message": "deleted post successfully"})
}
