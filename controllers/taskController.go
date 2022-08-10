package controllers

import (
	"Golang-Task/database"
	"Golang-Task/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct{}

func (TaskController) Root(c *gin.Context) {

	c.Redirect(http.StatusFound, "/task")
}

func (TaskController) GetAll(c *gin.Context) {
	db := database.Instance()

	var task []models.Task
	db.Find(&task)
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Golang Task - Dashboard",
		"task":  task,
	})

}

func (TaskController) GetDetail(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	var task models.Task
	result := db.First(&task, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data successfully!",
			"task":    task,
		})
	}
}

func (TaskController) ViewStore(c *gin.Context) {
	c.HTML(http.StatusOK, "add", gin.H{
		"title": "Golang Task - Input Form",
	})
}

func (TaskController) ViewUpdate(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	var task models.Task

	result := db.First(&task, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		c.HTML(http.StatusOK, "update", gin.H{
			"title": "Golang Task - Update Form",
			"task":  task,
		})
	}
}

func (TaskController) Store(c *gin.Context) {
	db := database.Instance()
	task_detail := c.PostForm("task_detail")
	assigne := c.PostForm("assigne")
	deadline := c.PostForm("deadline")

	db.Create(&models.Task{
		Task_detail: task_detail,
		Assigne:     assigne,
		Deadline:    deadline,
	})

	c.Redirect(http.StatusFound, "/task")
}

func (TaskController) Update(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	task_detail := c.PostForm("task_detail")
	assigne := c.PostForm("assigne")
	deadline := c.PostForm("deadline")

	var task models.Task
	result := db.First(&task, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		task.Task_detail = task_detail
		task.Assigne = assigne
		task.Deadline = deadline
		db.Save(&task)
	}
	c.Redirect(http.StatusFound, "/task")
}

func (TaskController) Done(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	var task models.Task
	result := db.Find(&task, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		task.IsDone = true
		db.Save(&task)
	}
	c.Redirect(http.StatusFound, "/task")
}

func (TaskController) Delete(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	var task models.Task
	result := db.Find(&task, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		db.Delete(&task)
	}
	c.Redirect(http.StatusFound, "/task")
}
