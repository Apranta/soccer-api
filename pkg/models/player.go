package models

// Player model.
type Player struct {
	CreatedUpdated

	ID           int64  `json:"id" db:"id"`
	TeamID       int64  `json:"team_id" db:"team_id" valid:"required"`
	Name         string `json:"name" db:"name" valid:"required"`
	JerseyNumber string `json:"jersey_number" db:"jersey_number" valid:"required"`
}
