package rest

import (
	"tracker/handler"
	"tracker/repo"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine, tasksRepo repo.Tasks) {
	h := handler.NewTasksHandler(tasksRepo)

	t := r.Group("/tasks")
	t.POST("", h.Create)
}
