package models

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type BaseModel struct{

}

// sendJSON is a helper function to send JSON responses and handle errors.
func (m *BaseModel) sendJSON(c echo.Context, code int, i interface{}) error {
	if err := c.JSON(code, i); err != nil {
		return fmt.Errorf("failed to send response: %w", err)
	}

	return nil
}