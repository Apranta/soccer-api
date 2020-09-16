package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"soccer/pkg/models"
)

// List teams
// @Summary List teams
// @Description Get the list of teams
// @Tags teams
// @ID list-teams
// @Produce json
// @Success 200 {array} models.Team
// @Router /teams [get]
func (api *API) listTeams(c echo.Context) error {
	ctx := c.Request().Context()

	teams, err := api.teamsService.ListTeams(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, teams)
}

// Get an team
// @Summary Get an team
// @Description Get an team by id
// @Tags teams
// @ID get-team
// @Produce json
// @Param id path int true "Team ID"
// @Success 200 {object} models.Team
// @Router /teams/{id} [get]
func (api *API) getTeam(c echo.Context) error {
	ctx := c.Request().Context()

	idString := c.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	team, err := api.teamsService.GetTeam(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, team)
}

// Create a new team
// @Summary Create a new team
// @Description Create a new team
// @Tags teams
// @ID create-team
// @Produce json
// @Param team body models.Team true "Create team"
// @Success 201 {object} models.Team
// @Router /teams [post]
func (api *API) createTeam(c echo.Context) error {
	ctx := c.Request().Context()

	team := new(models.Team)
	if err := c.Bind(team); err != nil {
		return err
	}

	if err := c.Validate(team); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newTeam, err := api.teamsService.CreateTeam(ctx, *team)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, newTeam)
}

// Delete an team
// @Summary Delete an team
// @Description Delete an team by id
// @Tags teams
// @ID delete-team
// @Produce plain
// @Param id path int true "Team ID"
// @Success 204 {string} string ""
// @Router /teams/{id} [delete]
func (api *API) deleteTeam(c echo.Context) error {
	ctx := c.Request().Context()

	idString := c.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	if err := api.teamsService.DeleteTeam(ctx, id); err != nil {
		return err
	}

	return c.String(http.StatusNoContent, "")
}

// Update an team
// @Summary Update an team
// @Description Update an team
// @Tags teams
// @ID update-team
// @Produce plain
// @Param id path int true "Team ID"
// @Param team body models.Team true "Update team"
// @Success 201 {string} string ""
// @Router /teams/{id} [put]
func (api *API) updateTeam(c echo.Context) error {
	ctx := c.Request().Context()

	idString := c.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	team := new(models.Team)
	if err := c.Bind(team); err != nil {
		return err
	}

	if err := c.Validate(team); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	team.ID = id
	updatedTeam, err := api.teamsService.UpdateTeam(ctx, *team)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, updatedTeam)
}
