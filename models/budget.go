package models

import (
	"net/http"
	"strconv"

	"github.com/galamdring/go-gold/ent"
	"github.com/galamdring/go-gold/ent/budget"
	"github.com/galamdring/go-gold/ent/schema"
	"github.com/galamdring/go-gold/ent/user"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

const (
	InvalidBudgetIdMessage = "Invalid budget ID format"
)

type BudgetRestModel struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Amount decimal.Decimal `json:"amount"`
	User *UserRestModel `json:"user,omitempty"`

}

func BuildBudgetRestModel(bud *ent.Budget) *BudgetRestModel {
	var userVal *UserRestModel
	if bud.Edges.User != nil {
		userVal = BuildUserRestModel(bud.Edges.User)
	}
	return &BudgetRestModel{
		ID: 	bud.ID,
		Name:   bud.Name,
		Amount: decimal.Decimal(bud.Amount),
		User: userVal,
	}
}

func BuildBudgetRestModels(buds []*ent.Budget) (result []*BudgetRestModel) {
	for _, bud := range buds {
		result = append(result, BuildBudgetRestModel(bud))
	}

	return result
}

// BudgetModel handles budget-related operations.
type BudgetModel struct {
	BaseModel
	client *ent.Client
}

// NewBudgetModel creates a new BudgetModel.
func NewBudgetModel(client *ent.Client) *BudgetModel {
	return &BudgetModel{
		BaseModel: BaseModel{},
		client: client,
	}
}

// CreateBudget creates a new budget.
func (m *BudgetModel) CreateBudget(eContext echo.Context) error {
	var b BudgetRestModel
	if err := eContext.Bind(&b); err != nil {
		return m.sendJSON(eContext, http.StatusBadRequest, err.Error())
	}

	bud, err := m.client.Budget.Create().
		SetUserID(b.User.ID).
		SetName(b.Name).
		SetAmount(schema.Decimal(b.Amount)).
		Save(eContext.Request().Context())
	if err != nil {
		return m.sendJSON(eContext, http.StatusInternalServerError, err.Error())
	}

	// Create doesn't actually load the edges, so Load it manually to include it in the response
	bud.Edges.User, err = bud.QueryUser().Only(eContext.Request().Context())
	if err != nil {
		return m.sendJSON(eContext, http.StatusInternalServerError, err.Error())
	}

	return m.sendJSON(eContext, http.StatusCreated, BuildBudgetRestModel(bud))
}

// GetAllBudgets retrieves all budgets.
func (m *BudgetModel) GetAllBudgets(eContext echo.Context) error {
	userIdStr := eContext.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return m.sendJSON(eContext, http.StatusBadRequest, "Invalid user ID")
	}

	budgets, err := m.client.Budget.Query().
		Where(budget.HasUserWith(user.ID(userId))).
		WithUser().
		All(eContext.Request().Context())
	if err != nil {
		return m.sendJSON(eContext, http.StatusInternalServerError, err.Error())
	}

	if len(budgets) == 0 {
		return m.sendJSON(eContext, http.StatusOK, []interface{}{})
	}

	return m.sendJSON(eContext, http.StatusOK, BuildBudgetRestModels(budgets))
}

// GetBudget retrieves a budget by ID.
func (m *BudgetModel) GetBudget(eContext echo.Context) error {
	id := eContext.Param("id")

	budgetId, err := strconv.Atoi(id)
	if err != nil {
		return m.sendJSON(eContext, http.StatusBadRequest, map[string]string{"error": InvalidBudgetIdMessage})
	}

	bud, err := m.client.Budget.Get(eContext.Request().Context(), budgetId)
	if err != nil {
		return m.sendJSON(eContext, http.StatusNotFound, err.Error())
	}

	return m.sendJSON(eContext, http.StatusOK, BuildBudgetRestModel(bud))
}

// UpdateBudget updates an existing budget.
func (m *BudgetModel) UpdateBudget(eContext echo.Context) error {
	id := eContext.Param("id")
	var b BudgetRestModel
	if err := eContext.Bind(&b); err != nil {
		return m.sendJSON(eContext, http.StatusBadRequest, err.Error())
	}

	eContext.Logger().Infof("received request to update budget with id#%s and user #%s")

	budgetId, err := strconv.Atoi(id)
	if err != nil {
		return m.sendJSON(eContext, http.StatusBadRequest, map[string]string{"error": InvalidBudgetIdMessage})
	}

	bud, err := m.client.Budget.UpdateOneID(budgetId).
		SetUserID(b.User.ID).
		SetName(b.Name).
		SetAmount(schema.Decimal(b.Amount)).
		Save(eContext.Request().Context())
	if err != nil {
		return m.sendJSON(eContext, http.StatusInternalServerError, err.Error())
	}

	return m.sendJSON(eContext, http.StatusOK, BuildBudgetRestModel(bud))
}

// DeleteBudget deletes a budget by ID.
func (m *BudgetModel) DeleteBudget(eContext echo.Context) error {
	id := eContext.Param("id")

	budgetId, err := strconv.Atoi(id)
	if err != nil {
		return m.sendJSON(eContext, http.StatusBadRequest, map[string]string{"error": InvalidBudgetIdMessage})
	}

	err = m.client.Budget.DeleteOneID(budgetId).Exec(eContext.Request().Context())
	if err != nil {
		return m.sendJSON(eContext, http.StatusNotFound, err.Error())
	}

	return m.sendJSON(eContext, http.StatusOK, map[string]string{"message": "Budget deleted successfully"})
}