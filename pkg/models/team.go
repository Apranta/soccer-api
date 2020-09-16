package models

// Team model.
type Team struct {
	CreatedUpdated

	ID          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name" valid:"required"`
	Description string `json:"description" db:"description" valid:"required"`
}
