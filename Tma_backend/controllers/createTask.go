package controllers 

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"Tma_backend/ent"
)

type CreateTaskRequest struct {
    Title       string    `json:"title" binding:"required"`
    Description string    `json:"description" binding:"required"`
    DueDate     time.Time `json:"due_date" binding:"required"`
}

// This API will create a task with all the given details.
func CreateTaskHandler(c *gin.Context) {
    var req CreateTaskRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    client := c.MustGet("ent").(*ent.Client)
    _, err := client.Task.
        Create().
        SetTitle(req.Title).
        SetDescription(req.Description).
        SetDueDate(req.DueDate).
        Save(context.Background())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully!"})
}
