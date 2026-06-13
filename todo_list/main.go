package main

import (
	"todolist/scanner"
	"todolist/todo"
)

func main() {
	todoList := todo.NewList()
	scanner := scanner.NewScanner(todoList)
	scanner.Start()
}