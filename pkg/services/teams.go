package services

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"soccer/pkg/models"
)

// TeamsService service interface.
type TeamsService interface {
	ListTeams(ctx context.Context) ([]models.Team, error)
	GetTeam(ctx context.Context, id int64) (models.Team, error)
	CreateTeam(ctx context.Context, team models.Team) (models.Team, error)
	DeleteTeam(ctx context.Context, id int64) error
	UpdateTeam(ctx context.Context, team models.Team) (models.Team, error)
}

type teamsService struct {
	db *sqlx.DB
}

// NewTeamsService returns an initialized TeamsService implementation.
func NewTeamsService(db *sqlx.DB) TeamsService {
	return &teamsService{db: db}
}

func (s *teamsService) ListTeams(ctx context.Context) ([]models.Team, error) {
	query := `
		SELECT
			id
			, name
			, description
			, created_at
			, updated_at
		FROM teams`

	var teams []models.Team
	if err := s.db.SelectContext(ctx, &teams, query); err != nil {
		return nil, fmt.Errorf("get the list of teams: %s", err)
	}

	return teams, nil
}

func (s *teamsService) GetTeam(ctx context.Context, id int64) (models.Team, error) {
	query := `
		SELECT
			id
			, name
			, description
			, created_at
			, updated_at
		FROM teams
		WHERE id = $1`

	var team models.Team
	if err := s.db.GetContext(ctx, &team, query, id); err != nil {
		return models.Team{}, fmt.Errorf("get an team: %s", err)
	}

	return team, nil
}

func (s *teamsService) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	query := "INSERT INTO teams (name, description) VALUES ($1, $2) RETURNING id"

	var id int64
	if err := s.db.QueryRowxContext(ctx, query, team.Name, team.Description).Scan(&id); err != nil {
		return models.Team{}, fmt.Errorf("insert new team: %s", err)
	}

	newTeam, err := s.GetTeam(ctx, id)
	if err != nil {
		return models.Team{}, fmt.Errorf("get new team: %s", err)
	}

	return newTeam, nil
}

func (s *teamsService) DeleteTeam(ctx context.Context, id int64) error {
	query := `DELETE FROM teams WHERE id = $1`

	if _, err := s.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("delete an team: %s", err)
	}

	return nil
}

func (s *teamsService) UpdateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	query := `UPDATE teams SET name=$1, description=$2  Where id=$3`

	if _, err := s.db.ExecContext(ctx, query, team.Name, team.Description, team.ID); err != nil {
		return models.Team{}, fmt.Errorf("Update team: %s", err)
	}

	Team, err := s.GetTeam(ctx, team.ID)
	if err != nil {
		return models.Team{}, fmt.Errorf("get team: %s", err)
	}

	return Team, nil
}
