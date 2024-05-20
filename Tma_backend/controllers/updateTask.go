package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"Tma_backend/ent"
	"Tma_backend/ent/task"
)

type TaskURIRequest struct {
	ID  int `uri:"id" binding:"required"`
}

type TaskPUTRequest struct {
	Title       string    `json:"title" binding:"required"`
    Description string    `json:"description" binding:"required"`
    DueDate     time.Time `json:"due_date" binding:"required"`
}

// This API fetch the data of a specific task.
func GetTaskByID(c *gin.Context) {
	client := c.MustGet("ent").(*ent.Client)

	var u TaskURIRequest
	if err := c.ShouldBindUri(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Querying a task by id.
	task, err := client.Task.Query().
	Where(task.IDEQ(uint64(u.ID))).
	Only(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch a task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// This API will update the exisiting record in the DB.
func UpdateTaskByID(c *gin.Context) {
	client := c.MustGet("ent").(*ent.Client)

	var req TaskPUTRequest
	var u TaskURIRequest
	
	if err := c.ShouldBindUri(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskDB, err := client.Task.Query().
	Where(task.IDEQ(uint64(u.ID))).
	Only(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch a task"})
		return
	}

	// Updating task record.
	_, err = taskDB.Update().
		SetTitle(req.Title).
        SetDescription(req.Description).
        SetDueDate(req.DueDate).
		Save(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task updated successfully!"})
}
