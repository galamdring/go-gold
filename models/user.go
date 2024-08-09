package models

import (
	"context"
	"net/http"
	"strconv"

	_ "entgo.io/ent/dialect/sql"
	"github.com/galamdring/go-gold/ent" // Adjust the import path as necessary
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Budgets  []Budget `json:"budgets"`
}

// UserModel holds the user-related database operations
type UserModel struct {
    client *ent.Client
}

// NewUserModel creates a new UserModel
func NewUserModel(client *ent.Client) *UserModel {
    return &UserModel{client: client}
}

// CreateUser creates a new user
func (m *UserModel) CreateUser(c echo.Context) error {
    var user ent.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    newUser, err := m.client.User.Create().
        SetUsername(user.Username).
        SetEmail(user.Email).
        Save(context.Background())
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusCreated, newUser)
}

// GetUser retrieves a user by ID
func (m *UserModel) GetUser(c echo.Context) error {
    id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid user ID")
    }
    user, err := m.client.User.Get(context.Background(), userID)
    if err != nil {
        return c.JSON(http.StatusNotFound, "User not found")
    }
    return c.JSON(http.StatusOK, user)
}

// UpdateUser updates a user by ID
func (m *UserModel) UpdateUser(c echo.Context) error {
    id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid user ID")
    }
    var user ent.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    updatedUser, err := m.client.User.UpdateOneID(userID).
        SetUsername(user.Username).
        SetEmail(user.Email).
        Save(context.Background())
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser deletes a user by ID
func (m *UserModel) DeleteUser(c echo.Context) error {
    id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid user ID")
    }
    err = m.client.User.DeleteOneID(userID).Exec(context.Background())
    if err != nil {
        return c.JSON(http.StatusNotFound, "User not found")
    }
    return c.NoContent(http.StatusNoContent)
}