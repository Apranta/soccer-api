package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"soccer/pkg/models"
)

// List players
// @Summary List players
// @Description Get the list of players
// @Tags players
// @ID list-players
// @Produce json
// @Success 200 {array} models.Player
// @Router /players [get]
func (api *API) listPlayers(c echo.Context) error {
	ctx := c.Request().Context()

	players, err := api.playersService.ListPlayers(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, players)
}

// List players by Team ID
// @Summary List players by team
// @Description Get the list of players by team
// @Tags players
// @ID list-players-team
// @Produce json
// @Param team_id path int true "Team ID"
// @Success 200 {array} models.Player
// @Router /players/{team_id} [get]
func (api *API) listPlayersByTeams(c echo.Context) error {
	ctx := c.Request().Context()

	idString := c.Param("team_id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	players, err := api.playersService.ListPlayersByTeams(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, players)
}

// Get an player
// @Summary Get an player
// @Description Get an player by id
// @Tags players
// @ID get-player
// @Produce json
// @Param id path int true "Player ID"
// @Success 200 {object} models.Player
// @Router /players/{team_id}/detail/{id} [get]
func (api *API) getPlayer(c echo.Context) error {
	ctx := c.Request().Context()

	_ = c.Param("team_id")
	idString := c.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	player, err := api.playersService.GetPlayer(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, player)
}

// Create a new player
// @Summary Create a new player
// @Description Create a new player
// @Tags players
// @ID create-player
// @Produce json
// @Param player body models.Player true "Create player"
// @Success 201 {object} models.Player
// @Router /players [post]
func (api *API) createPlayer(c echo.Context) error {
	ctx := c.Request().Context()

	player := new(models.Player)
	if err := c.Bind(player); err != nil {
		return err
	}

	if err := c.Validate(player); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newPlayer, err := api.playersService.CreatePlayer(ctx, *player)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, newPlayer)
}

// Delete an player
// @Summary Delete an player
// @Description Delete an player by id
// @Tags players
// @ID delete-player
// @Produce plain
// @Param id path int true "Player ID"
// @Success 204 {string} string ""
// @Router /players/{id} [delete]
func (api *API) deletePlayer(c echo.Context) error {
	ctx := c.Request().Context()

	idString := c.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	if err := api.playersService.DeletePlayer(ctx, id); err != nil {
		return err
	}

	return c.String(http.StatusNoContent, "")
}

// Update an player
// @Summary Update an player
// @Description Update an player
// @Tags players
// @ID update-player
// @Produce plain
// @Param id path int true "Player ID"
// @Param player body models.Player true "Update player"
// @Success 201 {string} string ""
// @Router /players/{id} [put]
func (api *API) updatePlayer(c echo.Context) error {
	ctx := c.Request().Context()

	idString := c.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	player := new(models.Player)
	if err := c.Bind(player); err != nil {
		return err
	}

	if err := c.Validate(player); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	player.ID = id
	updatedPlayer, err := api.playersService.UpdatePlayer(ctx, *player)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, updatedPlayer)
}
