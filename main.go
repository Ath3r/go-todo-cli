package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ath3r/go-todo-cli.git/models"
)

func main() {
	db, err := models.NewDB()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer models.CloseDB(db)


	fmt.Println("Connected to database")


	todo := &models.Todo{}
	
	err = todo.CreateTable(db)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Get all todos")
		fmt.Println("2. Get todo by id")
		fmt.Println("3. Create todo")
		fmt.Println("4. Update todo")
		fmt.Println("5. Delete todo")
		fmt.Println("6. Exit")

		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		fmt.Println("You chose:", choice)

		switch choice {
		case "1":
			todo.GetAll(db)
		case "2":
			todo.GetById(db, reader)
		case "3":
			todo.CreateTodo(db, reader)
		case "4":
			todo.Update(db, reader)
		case "5":
			todo.Delete(db, reader)
		case "6":
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}