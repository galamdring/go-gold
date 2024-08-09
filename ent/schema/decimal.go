package schema

import (
	"database/sql/driver"
	"fmt"
	"github.com/shopspring/decimal"
)

// Decimal is a custom type for handling decimal values.
type Decimal decimal.Decimal

// Scan implements the sql.Scanner interface.
func (d *Decimal) Scan(value interface{}) error {
	if value == nil {
		*d = Decimal(decimal.Zero) // Handle NULL values
		return nil
	}
	decimalValue, err := decimal.NewFromString(fmt.Sprintf("%v", value))
	if err != nil {
		return err
	}
	*d = Decimal(decimalValue)
	return nil
}

// Value implements the driver.Valuer interface.
func (d Decimal) Value() (driver.Value, error) {
	return decimal.Decimal(d).String(), nil
}