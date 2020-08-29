package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kookkla/myapi/task"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

func createTodosHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id", t.Title, t.Status)

	err := row.Scan(&t.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, t)
}

func authMiddleware(c *gin.Context) {
	fmt.Println("start #middleware")

	token := c.GetHeader("Authorization")
	if token != "Bearer token1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you don't have the permission!!"})
		c.Abort()
		return
	}

	c.Next()

	fmt.Println("end #middleware")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")

	apiV1.Use(authMiddleware)

	apiV1.GET("/todos", task.GetTodosHandler)
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
