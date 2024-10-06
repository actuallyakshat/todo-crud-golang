package main

import (
	"github.com/actuallyakshat/todo-crud/initialisers"
	"github.com/actuallyakshat/todo-crud/models"
)

func init() {
	initialisers.LoadEnvironmentVariables()
	initialisers.ConnectToDB()
}

func main() {
	initialisers.DB.AutoMigrate(&models.Todo{})
}
