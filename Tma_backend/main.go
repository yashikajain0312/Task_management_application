package main

import (
    "context"
    "net/http"
	"fmt"

    "github.com/gin-gonic/gin"
    "Tma_backend/controllers"
	_ "github.com/lib/pq"
    "Tma_backend/ent"
    "Tma_backend/ent/migrate"
)

func main() {
	connStr := "postgres://postgres:ryd12@localhost:5432/taskmanagementapp?sslmode=disable"

    // Initializing Ent client with PostgreSQL driver.
    client, err := ent.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    // Running Ent migrations.
    if err := client.Schema.Create(
        context.Background(),
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true),
    ); err != nil {
        fmt.Println(err)
    }

    // Creating Gin router.
    r := gin.Default()

    // Middleware to inject Ent client into Gin context.
    r.Use(func(c *gin.Context) {
        c.Set("ent", client)
        c.Next()
    })

	// Use JWT middleware.
	r.Use(controllers.JWTMiddleware())

    // Registering handlers.
    r.POST("/user/signup", controllers.SignupHandler)
	r.GET("/user/login", controllers.LoginHandler)
	r.GET("/user/task/list", controllers.TaskListHandler)
	r.POST("/user/task/add", controllers.CreateTaskHandler)
	r.GET("/user/task/:id/get", controllers.GetTaskByID)
	r.PUT("/user/task/:id/update", controllers.UpdateTaskByID)
	r.POST("/user/assign/task/:task_id/user/:user_id", controllers.AssignTaskHandler)
	r.GET("/user/:user_id/tasks/view", controllers.ViewUserTasksHandler)
    
    // Running the server.
    if err := http.ListenAndServe(":4000", r); err != nil {
        panic(err)
    }
}
