package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/anuchito/myapi/middleware"
	"github.com/anuchito/myapi/task"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")

	apiV1.Use(middleware.Auth)

	h := task.Handler{
		DB: db
	}

	apiV1.GET("/todos", h.GetTodosHandler)
	apiV1.GET("/todos/:id", task.GetTodoByIdHandler)
	apiV1.POST("/todos", task.CreateTodosHandler)
	apiV1.PUT("/todos/:id", task.UpdateTodosHandler)
	apiV1.DELETE("/todos/:id", task.DeleteTodosHandler)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":1234")
}
