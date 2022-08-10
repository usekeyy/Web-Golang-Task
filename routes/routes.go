package routes

import (
	"Golang-Task/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	taskController := controllers.TaskController{}

	r.GET("/task", taskController.GetAll)

	r.GET("/task/:id", taskController.GetDetail)

	r.GET("/task/add", taskController.ViewStore)

	r.POST("/task/add", taskController.Store)

	r.GET("/task/update/:id", taskController.ViewUpdate)

	r.POST("/task/update/:id", taskController.Update)

	r.GET("/task/:id/done", taskController.Done)

	r.DELETE("/task/:id", taskController.Delete)

}
