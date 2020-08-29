package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

func getTodosHandler(c *gin.Context) {
	
	db, err := sql.Open("postgres", "postgres://vttkxspt:sjA5CdRG1tepOQye8KB1ZMsPjQZ273V9@lallah.db.elephantsql.com:5432/vttkxspt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	log.Println("Okay")

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Fatal("can't prepare query one row statment", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)

	var id int
	var title, status string
	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	fmt.Println("one row", id, title, status)
}

func getTodoByIdHandler(c *gin.Context) {
	id := c.Param("id")

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	row := stmt.QueryRow(id)

	t := &Todo{}

	err = row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, t)
}

func updateTodosHandler(c *gin.Context) {
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	row := stmt.QueryRow(id)

	t := &Todo{}

	err = row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := c.ShouldBindJSON(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err = db.Prepare("UPDATE todos SET status=$2, title=$3 WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := stmt.Exec(id, t.Status, t.Title); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, t)
}

func deleteTodosHandler(c *gin.Context) {
	id := c.Param("id")
	stmt, err := db.Prepare("DELETE FROM todos WHERE id = $1")
	if err != nil {
		log.Fatal("can't prepare delete statement", err)
	}

	if _, err := stmt.Exec(id); err != nil {
		log.Fatal("can't execute delete statment", err)
	}

	c.JSON(http.StatusOK, "deleted todo.")
}

func authMiddleware(c *gin.Context) {

	c.Next()

}
func setupRouter() *gin.Engine {
	r := gin.Default()

	//apiV1 := r.Group("/api/v1")

	//apiV1.Use(authMiddleware)

	r.GET("/todos", getTodosHandler)
	r.GET("/todos/:id", getTodoByIdHandler)
	r.POST("/todos", createTodosHandler)
	r.PUT("/todos/:id", updateTodosHandler)
	r.DELETE("/todos/:id", deleteTodosHandler)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":4444")
}
