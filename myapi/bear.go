package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var err error
var db *sql.DB

func init() {
	db, err = sql.Open("postgres", "postgres://gosctihb:CqOz6dVYlooEBPY4quY9KHvySa2OmADZ@arjuna.db.elephantsql.com:5432/gosctihb")
	if err != nil {
		log.Fatal(err)
	}
}

//Todo ...
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func getTodoByIDHandler(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id")) //convert id from string to int - strconv is string convert
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	row := stmt.QueryRow(id)
	var title, status string
	err = row.Scan(&id, &title, &status)
	if err != nil {
		context.JSON(http.StatusOK, "No data found!")
	} else {
		context.JSON(http.StatusOK, Todo{ID: id, Title: title, Status: status})
	}
}

func getTodosHandler(context *gin.Context) {
	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	rows, err := stmt.Query()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		items := []*Todo{}

		for rows.Next() {
			var id int
			var title, status string
			err := rows.Scan(&id, &title, &status)
			if err != nil {
				log.Fatal("can't Scan row into variable", err)
			} else {
				items = append(items, &Todo{ID: id, Title: title, Status: status})
			}
		}

		context.JSON(http.StatusOK, items)
	}
}

func createTodosHandler(context *gin.Context) {
	t := Todo{}
	if err := context.ShouldBindJSON(&t); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2) RETURNING id", t.Title, t.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusCreated, id)
	}
}

func updateTodosHandler(context *gin.Context) {
	t := Todo{}
	if err := context.ShouldBindJSON(&t); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := db.Prepare("UPDATE todos SET title=$2, status=$3 WHERE id=$1;")
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := stmt.Exec(t.ID, t.Title, t.Status); err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusCreated, t)
	}
}

func authMiddleware(c *gin.Context) {
	log.Println("start middleware")
	authKey := c.GetHeader("Authorization")
	if authKey != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		c.Abort()
		return
	}
	c.Next()
	log.Println("end middleware")
}

func setupRouter() *gin.Engine {
	route := gin.Default()
	route.Use(authMiddleware)

	route.GET("/todos", getTodosHandler)
	route.GET("/todos/:id", getTodoByIDHandler)
	route.POST("/todos", createTodosHandler)
	route.PUT("/todos", updateTodosHandler)

	return route
}

func main() {
	route := setupRouter()

	route.Run(":1234") // listen and serve on 127.0.0.0:8080
}
