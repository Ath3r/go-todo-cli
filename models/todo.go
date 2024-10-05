package models

import (
	"bufio"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Todo) CreateTodo(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("Enter todo title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	todo := &Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	query := `
		INSERT INTO todo (title, completed, created_at) VALUES (?, ?, ?);
	`
	_, err := db.Exec(query, todo.Title, todo.Completed, todo.CreatedAt)
	if err != nil {
		fmt.Println("Failed to create todo:", err)
		return
	}
	fmt.Println("Todo created successfully")
}

func (t *Todo) CreateTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS todo (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			completed BOOLEAN NOT NULL,
			created_at TIMESTAMP NOT NULL
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}

func (t *Todo) GetAll(db *sql.DB){
	query := `
		SELECT id, title, completed, created_at FROM todo;
	`
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println("Failed to get all todos:", err)
		return
	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
		if err != nil {
			fmt.Println("Failed to scan todo:", err)
			return
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Failed to get all todos:", err)
		return
	}

	for _, todo := range todos {
		fmt.Printf("ID: %d, Title: %s, Completed: %v, Created At: %v\n", todo.ID, todo.Title, todo.Completed, todo.CreatedAt)
	}
}

func (t *Todo) GetById(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("Enter todo ID: ")
    idStr, _ := reader.ReadString('\n')
    id, err := strconv.Atoi(strings.TrimSpace(idStr))
    if err != nil {
        fmt.Println("Invalid ID:", err)
        return
    }
	query := `
		SELECT id, title, completed, created_at FROM todo WHERE id = ?;
	`
	row := db.QueryRow(query, id)
	var todo Todo
	err = row.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		fmt.Println("Failed to get todo by id:", err)
		return
	}
	fmt.Printf("ID: %d, Title: %s, Completed: %v, Created At: %v\n", todo.ID, todo.Title, todo.Completed, todo.CreatedAt)
}

func (t *Todo) Update(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("Enter todo ID: ")
    idStr, _ := reader.ReadString('\n')
    id, err := strconv.Atoi(strings.TrimSpace(idStr))
    if err != nil {
        fmt.Println("Invalid ID:", err)
        return
    }
	fmt.Print("Enter todo title: ")
    title, _ := reader.ReadString('\n')
    title = strings.TrimSpace(title)

    fmt.Print("Enter todo completed (true/false): ")
    completedStr, _ := reader.ReadString('\n')
    completed, err := strconv.ParseBool(strings.TrimSpace(completedStr))
    if err != nil {
        fmt.Println("Invalid completed value:", err)
		return
	}
	query := `
		UPDATE todo SET title = ?, completed = ? WHERE id = ?;
	`
	_, err = db.Exec(query, title, completed, id)
	if err != nil {
		fmt.Println("Failed to update todo:", err)
		return
	}
	fmt.Println("Todo updated successfully")
}

func (t *Todo) Delete(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("Enter todo ID: ")
    idStr, _ := reader.ReadString('\n')
    id, err := strconv.Atoi(strings.TrimSpace(idStr))
    if err != nil {
        fmt.Println("Invalid ID:", err)
        return
    }
	query := `
		DELETE FROM todo WHERE id = ?;
	`
	_, err = db.Exec(query, id)
	if err != nil {
		fmt.Println("Failed to delete todo:", err)
		return
	}
	fmt.Println("Todo deleted successfully")
}
