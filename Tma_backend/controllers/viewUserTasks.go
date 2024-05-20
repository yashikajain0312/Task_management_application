package controllers

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "Tma_backend/ent"
    "Tma_backend/ent/taskassignment"
    "Tma_backend/ent/user"
)

type ViewUserTasksURIRequest struct {
    UserID int `uri:"user_id" binding:"required"`
}

// This API will allow you to see all the tasks of a specific user.
// It will also allow you to see tasks based on given filter(i.e pending or completed)
func ViewUserTasksHandler(c *gin.Context) {
    var u ViewUserTasksURIRequest

    if err := c.ShouldBindUri(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    client := c.MustGet("ent").(*ent.Client)

    // Validating user existence.
    _, err := client.User.Query().Where(user.ID(uint64(u.UserID))).Only(context.Background())
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

	status := c.Query("status")

    // Querying to get all task assignments for the user.
    query := client.TaskAssignment.
        Query().
        Where(taskassignment.UserIDEQ(uint64(u.UserID))).
        WithTask() 

    // Apply status filter if provided.
    if status != "" {
        query = query.Where(taskassignment.StatusEQ(taskassignment.Status(status)))
    }

    taskAssignments, err := query.All(context.Background())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var tasks []ent.Task
    for _, assignment := range taskAssignments {
        tasks = append(tasks, *assignment.Edges.Task)
    }

    c.JSON(http.StatusOK, tasks)
}
