package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kookkla/golang/myapi/middleware"
	"github.com/kookkla/golang/myapi/task"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://vttkxspt:sjA5CdRG1tepOQye8KB1ZMsPjQZ273V9@lallah.db.elephantsql.com:5432/vttkxspt")
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")

	apiV1.Use(middleware.Auth)

	h := task.Handler{
		DB: db,
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
