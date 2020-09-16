package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"soccer/pkg/services"
)

// API can register a set of endpoints in a router and handle
// them using the provided storage.
type API struct {
	teamsService   services.TeamsService
	playersService services.PlayersService

	adminUsername string
	adminPassword string
}

// NewAPI returns an initialized API type.
func NewAPI(teamsService services.TeamsService,
	playersService services.PlayersService, adminUsername, adminPassword string) *API {
	return &API{
		teamsService:   teamsService,
		playersService: playersService,

		adminUsername: adminUsername,
		adminPassword: adminPassword,
	}
}

// Register the API's endpoints in the given router.
func (api *API) Register(g *echo.Group) {
	// Teams API
	g.GET("/teams", api.listTeams)
	g.GET("/teams/:id", api.getTeam)
	g.POST("/teams", api.createTeam, middleware.BasicAuth(api.adminValidator))
	g.DELETE("/teams/:id", api.deleteTeam, middleware.BasicAuth(api.adminValidator))
	g.PUT("/teams/:id", api.updateTeam, middleware.BasicAuth(api.adminValidator))

	// Teams API
	g.GET("/players", api.listPlayers)
	g.GET("/players/:team_id", api.listPlayersByTeams)
	g.GET("/players/:team_id/details/:id", api.getPlayer)
	g.POST("/players", api.createPlayer, middleware.BasicAuth(api.adminValidator))
	g.DELETE("/players/:id", api.deletePlayer, middleware.BasicAuth(api.adminValidator))
	g.PUT("/players/:id", api.updatePlayer, middleware.BasicAuth(api.adminValidator))
}

func (api *API) adminValidator(username, password string, c echo.Context) (bool, error) {
	if username == api.adminUsername && password == api.adminPassword {
		return true, nil
	}
	return false, nil
}
