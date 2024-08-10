package models

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	_ "entgo.io/ent/dialect/sql"        // sql dialect for ent
	"github.com/galamdring/go-gold/ent" // Adjust the import path as necessary
	"github.com/labstack/echo/v4"
)

const (
	InvalidUserIdMessage = "Invalid user ID"
	UserNotFoundMessage  = "User not found"
)

type UserRestModel struct {
	ID       int               `json:"id"`
	Username string            `json:"username"`
	Email    string            `json:"email"`
	Budgets  []*BudgetRestModel `json:"budgets,omitempty"`
}

func BuildUserRestModel(userVal *ent.User) *UserRestModel {
	var budgets []*BudgetRestModel
	if userVal.Edges.Budgets != nil {
		budgets = BuildBudgetRestModels(userVal.Edges.Budgets)
	}

	return &UserRestModel{
		ID: userVal.ID,
		Username: userVal.Username,
		Email: userVal.Email,
		Budgets: budgets,
	}
}

// UserModel holds the user-related database operations.
type UserModel struct {
	BaseModel
	client *ent.Client
}

// NewUserModel creates a new UserModel.
func NewUserModel(client *ent.Client) *UserModel {
	return &UserModel{
		BaseModel: BaseModel{},
		client: client,
	}
}

// CreateUser creates a new user.
func (m *UserModel) CreateUser(eContext echo.Context) error {
	var user ent.User
	if err := eContext.Bind(&user); err != nil {
		return m.sendJSON(eContext, http.StatusBadRequest,
			map[string]string{"error": fmt.Sprintf("failed to bind user: %v", err)})
	}

	newUser, err := m.client.User.Create().
		SetUsername(user.Username).
		SetEmail(user.Email).
		Save(context.Background())
	if err != nil {
		return m.sendJSON(eContext, http.StatusInternalServerError,
			map[string]string{"error": fmt.Sprintf("failed to create user: %v", err)})
	}

	return m.sendJSON(eContext, http.StatusCreated, newUser)
}

// GetAllUsers retrieves all users.
func (m *UserModel) GetAllUsers(c echo.Context) error {
	users, err := m.client.User.Query().All(context.Background())
	if err != nil {
		return m.sendJSON(c, http.StatusInternalServerError, map[string]string{"error": "failed to retrieve users"})
	}

	return m.sendJSON(c, http.StatusOK, map[string]interface{}{"users": users})
}

// GetUser retrieves a user by ID.
func (m *UserModel) GetUser(eContext echo.Context) error {
	id := eContext.Param("id")
	fmt.Printf("received get request for id %s\n", id)

	userID, err := strconv.Atoi(id)
	if err != nil {
		return m.sendJSON(eContext, http.StatusBadRequest, map[string]string{"error": InvalidUserIdMessage})
	}

	user, err := m.client.User.Get(context.Background(), userID)
	if err != nil {
		return m.sendJSON(eContext, http.StatusNotFound, map[string]string{"error": InvalidUserIdMessage})
	}

	return m.sendJSON(eContext, http.StatusOK, user)
}

// UpdateUser updates a user by ID.
func (m *UserModel) UpdateUser(eContext echo.Context) error {
	id := eContext.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		return m.sendJSON(eContext, http.StatusBadRequest, map[string]string{"error": InvalidUserIdMessage})
	}

	var user ent.User
	if err = eContext.Bind(&user); err != nil {
		return m.sendJSON(eContext,
			http.StatusBadRequest,
			map[string]string{"error": fmt.Sprintf("failed to bind user: %v", err)})
	}

	updatedUser, err := m.client.User.UpdateOneID(userID).
		SetUsername(user.Username).
		SetEmail(user.Email).
		Save(context.Background())
	if err != nil {
		return m.sendJSON(eContext, http.StatusInternalServerError,
			map[string]string{"error": fmt.Sprintf("failed to update user: %v", err)})
	}

	return m.sendJSON(eContext, http.StatusOK, updatedUser)
}

// DeleteUser deletes a user by ID.
func (m *UserModel) DeleteUser(eContext echo.Context) error {
	id := eContext.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		return m.sendJSON(eContext, http.StatusBadRequest, map[string]string{"error": InvalidUserIdMessage})
	}

	err = m.client.User.DeleteOneID(userID).Exec(context.Background())
	if err != nil {
		return m.sendJSON(eContext, http.StatusNotFound, map[string]string{"error": UserNotFoundMessage})
	}

	return m.sendJSON(eContext, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
