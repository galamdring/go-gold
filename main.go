package main

import (
	"context"
	"log"

	_ "entgo.io/ent/dialect/sql"
	"github.com/galamdring/go-gold/ent"
	"github.com/galamdring/go-gold/models"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

const (
	UsersEndpoint      = "/users"
	UserByIdEndpoint   = "/users/:id"
	BudgetsEndpoint    = "/budgets"
	BudgetByIdEndpoint = "/budgets/:id"
)

// Initialize the database.
func initDB() *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool to create the schema.
	if schemaErr := client.Schema.Create(context.Background()); schemaErr != nil {
		log.Fatalf("failed creating schema: %v", schemaErr)
	}

	return client
}

func main() {
	client := initDB()
	defer client.Close()

	router := echo.New()

	userModel := models.NewUserModel(client)
	budgetModel := models.NewBudgetModel(client)

	// CRUD routes for User
	router.POST(UsersEndpoint, userModel.CreateUser)      // Create
	router.GET(UsersEndpoint, userModel.GetAllUsers)      // Read All
	router.GET(UserByIdEndpoint, userModel.GetUser)       // Read
	router.PUT(UserByIdEndpoint, userModel.UpdateUser)    // Update
	router.DELETE(UserByIdEndpoint, userModel.DeleteUser) // Delete

	// CRUD routes for budget
	router.POST(BudgetsEndpoint, budgetModel.CreateBudget)      // Create
	router.GET(BudgetsEndpoint, budgetModel.GetAllBudgets)      // Read All
	router.GET(BudgetByIdEndpoint, budgetModel.GetBudget)       // Read
	router.PUT(BudgetByIdEndpoint, budgetModel.UpdateBudget)    // Update
	router.DELETE(BudgetByIdEndpoint, budgetModel.DeleteBudget) // Delete

	router.Logger.Fatal(router.Start(":8080"))

	// Run the auto migration tool to create the schema.
	if err := client.Schema.Create(context.Background()); err != nil {
		router.Logger.Fatal("failed creating schema: %v", err)
	}
}
