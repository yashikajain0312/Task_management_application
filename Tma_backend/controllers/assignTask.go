package controllers

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "Tma_backend/ent"
    "Tma_backend/ent/task"
    "Tma_backend/ent/user"
	"Tma_backend/ent/taskassignment"
)

type AssignTaskURIRequest struct {
    TaskID int `uri:"task_id" binding:"required"`
    UserID int `uri:"user_id" binding:"required"`
}

// This API will assign task to a specific user and if you'll try 
// to re-assign the same task to the same user, api will throw error.
func AssignTaskHandler(c *gin.Context) {
    var u AssignTaskURIRequest

    if err := c.ShouldBindUri(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    client := c.MustGet("ent").(*ent.Client)

    // Validating task existence.
    _, err := client.Task.Query().Where(task.ID(uint64(u.TaskID))).Only(context.Background())
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    // Validating user existence.
    _, err = client.User.Query().Where(user.ID(uint64(u.UserID))).Only(context.Background())
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

	// Checking if the task is already assigned to the user.
    existingAssignment, err := client.TaskAssignment.
        Query().
        Where(taskassignment.And(
            taskassignment.TaskIDEQ(uint64(u.TaskID)),
            taskassignment.UserIDEQ(uint64(u.UserID)),
        )).
        Only(context.Background())
    if err == nil && existingAssignment != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Task is already assigned to this user"})
        return
    }

    // Assigning the task to the user.
    _, err = client.TaskAssignment.
        Create().
        SetTaskID(uint64(u.TaskID)).
        SetUserID(uint64(u.UserID)).
		SetStatus(taskassignment.StatusPending).
        Save(context.Background())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task assigned successfully!"})
}
