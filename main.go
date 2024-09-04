package main

import (
	"tracker/config"
	"tracker/model"
	"tracker/repo"
	"tracker/rest"

	"github.com/gin-gonic/gin"
)

func SetupRouter(tasksRepo repo.Tasks) *gin.Engine {
	e := gin.Default()

	e.Use() // Cors
	e.Use() // Authorization
	e.Use() // Request-response logging

	rest.TaskRoutes(e, tasksRepo)

	return e
}

func main() {
	db := config.NewDB()
	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	err = conn.AutoMigrate(&model.Task{})
	if err != nil {
		panic(err)
	}

	tasksRepo := repo.NewTasks(conn)

	r := SetupRouter(tasksRepo)

	r.Run(":8080")
}
