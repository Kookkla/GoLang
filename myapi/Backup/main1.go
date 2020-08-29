package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status: "active"},
}

func getTodosHandler(c *gin.Context) {
	// log
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}

func getTodoByIdHandler(c *gin.Context) {
	// log
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, t)
}

func createTodosHandler(c *gin.Context) {
	// log
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := len(todos)
	id++
	t.ID = id
	todos[t.ID] = &t

	c.JSON(http.StatusCreated, "created todo.")
}
func helloHandler(c *gin.Context) {
	log.Println("in helloHandler")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		log.Println("start middleware")
		c.Next()
		log.Println("end middleware")
	})

	r.GET("/hello", helloHandler)
	r.GET("/todos", getTodosHandler)
	r.POST("/todos", createTodosHandler)
	r.GET("/todos/:id", getTodoByIdHandler)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":4444")
}
