package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"Tma_backend/ent"
)

// This API will list down all the task created by the user.
func TaskListHandler(c *gin.Context) {
	client := c.MustGet("ent").(*ent.Client)

	// Fetching all tasks.
	tasks, err := client.Task.Query().
	All(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the tasks"})
		return
	}

	c.JSON(http.StatusCreated, tasks)
}