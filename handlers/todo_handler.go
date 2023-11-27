package handlers

import (
	"database/sql"
	"fmt"
	"github.com/davidyapcc/todo-app-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func GetAllTodos(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		fmt.Println("Error getting Todos list", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)

	fmt.Println("here")
	var todos []models.ToDo
	for rows.Next() {
		var todo models.ToDo
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			fmt.Println("Error getting Todos list", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	var todo models.ToDo

	err := db.QueryRow("SELECT * FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Name, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.ToDo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	query := "INSERT INTO todos (name, description, status) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	res, err := stmt.Exec(newTodo.Name, newTodo.Description, newTodo.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	newTodo.ID = int(lastID)
	c.JSON(http.StatusCreated, newTodo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo models.ToDo

	if err := c.BindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := db.Exec("UPDATE todos SET name=?, description=?, status=? WHERE id=?", updatedTodo.Name, updatedTodo.Description, updatedTodo.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	err = db.QueryRow("SELECT * FROM todos WHERE id=?", id).Scan(&updatedTodo.ID, &updatedTodo.Name, &updatedTodo.Description, &updatedTodo.Status, &updatedTodo.CreatedAt, &updatedTodo.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated data"})
		return
	}

	c.JSON(http.StatusOK, updatedTodo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM todos WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
