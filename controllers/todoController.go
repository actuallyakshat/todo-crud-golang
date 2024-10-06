package controllers

import (
	"github.com/actuallyakshat/todo-crud/initialisers"
	"github.com/actuallyakshat/todo-crud/models"
	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	//Bind the request body to a struct
	var body struct {
		Title string `json:"title"`
	}

	//Check if the request body is valid
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid input: " + err.Error(),
		})
		return
	}

	//Create a new todo object with the title from the request body
	todo := models.Todo{
		Title:       body.Title,
		IsCompleted: false,
	}

	//Create the todo in the database
	result := initialisers.DB.Create(&todo)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error creating todo: " + result.Error.Error(),
		})
	} else {
		// If the todo is created successfully, return a 200 status code
		c.JSON(200, gin.H{
			"success": true,
			"message": "Success",
		})
	}
}

func GetTodos(c *gin.Context) {

	//Create a slice to store the resultant items
	var allTodos []models.Todo

	//Query the DB
	result := initialisers.DB.Find(&allTodos)

	//Check for errors
	if result.Error != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error fetching todos: " + result.Error.Error(),
		})
		return
	}

	//Return the response if resposne is valid.
	c.JSON(200, gin.H{
		"success": true,
		"todos":   allTodos,
	})
}

func GetTodoById(c *gin.Context) {

	//Extract id from the params and create a new todo object to store the result
	id := c.Param("id")
	var todo models.Todo

	//Find the todo with the given id
	result := initialisers.DB.First(&todo, id)

	//If the todo is not found, return a 404 error
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Todo not found",
		})
		return
	}

	//If there is an error, return a 500 error
	if result.Error != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error fetching todo: " + result.Error.Error(),
		})
		return
	}

	//If the todo is found, return the todo
	c.JSON(200, gin.H{
		"success": true,
		"todo":    todo,
	})
}

func UpdateTodo(c *gin.Context) {

	//Bind the request body to a struct
	var body struct {
		Title string `json:"title"`
	}

	//Check if the request body is valid
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid input: " + err.Error(),
		})
		return
	}

	//Extract id from the params and create a new todo object to store the result
	id := c.Param("id")
	var todo models.Todo

	//Find the todo with the given id
	result := initialisers.DB.First(&todo, id)

	//If the todo is not found, return a 404 error
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Todo not found",
		})
		return
	}

	//If there is an error, return a 500 error
	if result.Error != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error fetching todo: " + result.Error.Error(),
		})
		return
	}

	//If the todo is not found, return a 404 error
	if todo.ID == 0 {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Todo not found",
		})
		return
	}

	//Update the todo with the new title
	todo.Title = body.Title

	//Save the updated todo to the database
	result = initialisers.DB.Save(&todo)

	//If there is an error, return a 500 error
	if result.Error != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error updating todo: " + result.Error.Error(),
		})
		return
	}

	//If the todo is updated successfully, return a 200 status code
	c.JSON(200, gin.H{
		"success": true,
		"message": "Todo updated successfully",
	})
}

func DeleteTodo(c *gin.Context) {
	// Extract the id from the URL parameters
	id := c.Param("id")
	var todo models.Todo

	// Try to find the todo with the given id
	result := initialisers.DB.First(&todo, id)

	// Check if the todo was found (RowsAffected will be 0 if not)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Todo not found",
		})
		return
	}

	// If there is a database error, return a 500 error
	if result.Error != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error fetching todo: " + result.Error.Error(),
		})
		return
	}

	// Try to delete the todo
	if err := initialisers.DB.Delete(&todo).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error deleting todo: " + err.Error(),
		})
		return
	}

	// If the deletion is successful, return a 200 status code
	c.JSON(200, gin.H{
		"success": true,
		"message": "Todo deleted successfully",
	})
}

func CompleteTodo(c *gin.Context) {
	// Extract the id from the URL parameters
	id := c.Param("id")
	var todo models.Todo

	// Try to find the todo with the given id
	result := initialisers.DB.First(&todo, id)

	// Check if the todo was found (RowsAffected will be 0 if not)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Todo not found",
		})
		return
	}

	// If there is a database error, return a 500 error
	if result.Error != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error fetching todo: " + result.Error.Error(),
		})
		return
	}

	// Update the todo with the new title
	todo.IsCompleted = true

	// Save the updated todo to the database
	if err := initialisers.DB.Save(&todo).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Error updating todo: " + err.Error(),
		})
		return
	}

	// If the todo is updated successfully, return a 200 status code
	c.JSON(200, gin.H{
		"success": true,
		"message": "Todo updated successfully",
	})
}
