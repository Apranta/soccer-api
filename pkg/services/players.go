package services

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"soccer/pkg/models"
)

// PlayersService service interface.
type PlayersService interface {
	ListPlayers(ctx context.Context) ([]models.Player, error)
	ListPlayersByTeams(ctx context.Context, team int64) ([]models.Player, error)
	GetPlayer(ctx context.Context, id int64) (models.Player, error)
	CreatePlayer(ctx context.Context, player models.Player) (models.Player, error)
	DeletePlayer(ctx context.Context, id int64) error
	UpdatePlayer(ctx context.Context, player models.Player) (models.Player, error)
}

type playersService struct {
	db *sqlx.DB
}

// NewPlayersService returns an initialized PlayersService implementation.
func NewPlayersService(db *sqlx.DB) PlayersService {
	return &playersService{db: db}
}

func (s *playersService) ListPlayers(ctx context.Context) ([]models.Player, error) {
	query := `
		SELECT
			id
			, name
			, team_id
			, jersey_number
			, created_at
			, updated_at
		FROM players`

	var players []models.Player
	if err := s.db.SelectContext(ctx, &players, query); err != nil {
		return nil, fmt.Errorf("get the list of players: %s", err)
	}

	return players, nil
}

func (s *playersService) ListPlayersByTeams(ctx context.Context, team int64) ([]models.Player, error) {
	query := `
		SELECT
			id
			, name
			, team_id
			, jersey_number
			, created_at
			, updated_at
		FROM players WHERE team_id=$1`

	var players []models.Player
	if err := s.db.SelectContext(ctx, &players, query, team); err != nil {
		return nil, fmt.Errorf("get the list players in team: %s", err)
	}

	return players, nil
}

func (s *playersService) GetPlayer(ctx context.Context, id int64) (models.Player, error) {
	query := `
		SELECT
			id
			, name
			, team_id
			, jersey_number
			, created_at
			, updated_at
		FROM players
		WHERE id = $1`

	var player models.Player
	if err := s.db.GetContext(ctx, &player, query, id); err != nil {
		return models.Player{}, fmt.Errorf("get an player: %s", err)
	}

	return player, nil
}

func (s *playersService) CreatePlayer(ctx context.Context, player models.Player) (models.Player, error) {
	query := "INSERT INTO players (name, team_id ,jersey_number) VALUES ($1, $2 , $3) RETURNING id"

	var id int64
	if err := s.db.QueryRowxContext(ctx, query, player.Name, player.TeamID, player.JerseyNumber).Scan(&id); err != nil {
		return models.Player{}, fmt.Errorf("insert new player: %s", err)
	}

	newPlayer, err := s.GetPlayer(ctx, id)
	if err != nil {
		return models.Player{}, fmt.Errorf("get new player: %s", err)
	}

	return newPlayer, nil
}

func (s *playersService) DeletePlayer(ctx context.Context, id int64) error {
	query := `DELETE FROM players WHERE id = $1`

	if _, err := s.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("delete an player: %s", err)
	}

	return nil
}

func (s *playersService) UpdatePlayer(ctx context.Context, player models.Player) (models.Player, error) {
	query := `UPDATE players SET name=$1, jersey_number=$2 , team_id=$3  Where id=$4`

	if _, err := s.db.ExecContext(ctx, query, player.Name, player.JerseyNumber, player.TeamID, player.ID); err != nil {
		return models.Player{}, fmt.Errorf("Update player: %s", err)
	}

	Player, err := s.GetPlayer(ctx, player.ID)
	if err != nil {
		return models.Player{}, fmt.Errorf("get player: %s", err)
	}

	return Player, nil
}
