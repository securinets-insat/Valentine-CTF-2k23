package main

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func GetAllTasks(db *gorm.DB) []Task {
	tasks := []Task{}
	db.Select("id", "title").Find(&tasks)
	return tasks
}

func GetTaskById(db *gorm.DB, id int) Task {
	task := Task{}
	db.Where("id = ?", id).Select("id", "title", "description", "url").First(&task)
	return task
}


func SubmitFlag(id int, flag string) string {
	db := GetDb()

	var errorFlag string
	e := db.Raw("select validate('"+flag+"')").Scan(&errorFlag)

	fmt.Println("log:", e.Error, errorFlag)

	if e.Error != nil {
		return e.Error.Error()
	}

	task := Task{}
	db.Where("id = ?", id).Select("flag").First(&task)

	if   task.Flag == flag {
		return "Correct flag"
	}

	return "Wrong flag"
}


type FlagRequestBody struct {
    Flag string
}

func main() {

	db := GetDb()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))


	r.GET("/tasks", func(c *gin.Context) {
		tasks := GetAllTasks(db)
		c.JSON(http.StatusOK, tasks)
	})


	r.GET("/tasks/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		task := GetTaskById(db, id)
		c.JSON(http.StatusOK, task)
	})

	r.POST("/tasks/:id/submit", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var requestBody FlagRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, "Send Flag")
		} else {
			flag := requestBody.Flag
			c.JSON(http.StatusOK, SubmitFlag(id, flag))
		}
	})

	r.Run()
}
