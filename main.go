package main

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	_ "entgo.io/ent/dialect/sql"
	"github.com/galamdring/go-gold/ent" // Adjust the import path as necessary
	_ "github.com/mattn/go-sqlite3" // Use your preferred database driver
	"github.com/galamdring/go-gold/models"
)

var client *ent.Client

// Initialize the database
func initDB() {
    var err error
    client, err = ent.Open("sqlite3","file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    // Run the auto migration tool to create the schema.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema: %v", err)
    }
}

func main() {
	initDB()
	defer client.Close()

	e := echo.New()

    userModel := models.NewUserModel(client)

    // CRUD routes for User
    e.POST("/users", userModel.CreateUser)          // Create
    e.GET("/users/:id", userModel.GetUser)          // Read
    e.PUT("/users/:id", userModel.UpdateUser)       // Update
    e.DELETE("/users/:id", userModel.DeleteUser)    // Delete

    e.Logger.Fatal(e.Start(":8080"))

	// Run the auto migration tool to create the schema.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}

}