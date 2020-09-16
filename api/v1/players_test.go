package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"soccer/pkg/models"
	"soccer/pkg/services/mocks"
)

// type mockRequestValidator struct{}

// func (m *mockRequestValidator) Validate(i interface{}) error {
// 	return nil
// }

func TestAPI_listPlayers(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/players", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	mockPlayersService := &mocks.PlayersService{}
	mockPlayersService.On("ListPlayers", mock.Anything).Return([]models.Player{}, nil)

	api := NewAPI(&mocks.TeamsService{}, mockPlayersService, "", "")
	if assert.NoError(t, api.listPlayers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[]\n", rec.Body.String())
	}
}

func TestAPI_getPlayer(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/players/1", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath("/players/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockPlayersService := &mocks.PlayersService{}
	mockPlayersService.On("GetPlayer", mock.Anything, int64(1)).Return(models.Player{}, nil)

	api := NewAPI(&mocks.TeamsService{}, mockPlayersService, "", "")
	if assert.NoError(t, api.getPlayer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"id\":0,\"team_id\":0,\"name\":\"\",\"jersey_number\":\"\"}\n", rec.Body.String())
	}
}

func TestAPI_getPlayerTeams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/players/1/details/1", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath("/players/:team_id/details/:id")
	c.SetParamNames("team_id")
	c.SetParamValues("1")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockPlayersService := &mocks.PlayersService{}
	mockPlayersService.On("GetPlayer", mock.Anything, int64(1)).Return(models.Player{}, nil)

	api := NewAPI(&mocks.TeamsService{}, mockPlayersService, "", "")
	if assert.NoError(t, api.getPlayer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"id\":0,\"team_id\":0,\"name\":\"\",\"jersey_number\":\"\"}\n", rec.Body.String())
	}
}
func TestAPI_createPlayer(t *testing.T) {
	player := models.Player{
		ID:           1,
		Name:         "player-1",
		JerseyNumber: "10",
	}
	playerJSON, _ := json.Marshal(player)

	req := httptest.NewRequest(http.MethodPost, "/players", bytes.NewReader(playerJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	e.Validator = &mockRequestValidator{}
	c := e.NewContext(req, rec)

	mockPlayersService := &mocks.PlayersService{}
	mockPlayersService.On("CreatePlayer", mock.Anything, player).Return(player, nil)

	api := NewAPI(&mocks.TeamsService{}, mockPlayersService, "", "")
	if assert.NoError(t, api.createPlayer(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "{\"id\":1,\"team_id\":0,\"name\":\"player-1\",\"jersey_number\":\"10\"}\n", rec.Body.String())
	}
}

func TestAPI_deletePlayer(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/players/1", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath("/players/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockPlayersService := &mocks.PlayersService{}
	mockPlayersService.On("DeletePlayer", mock.Anything, int64(1)).Return(nil)

	api := NewAPI(&mocks.TeamsService{}, mockPlayersService, "", "")
	if assert.NoError(t, api.deletePlayer(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}

func TestAPI_updatePlayer(t *testing.T) {
	player := models.Player{
		Name:         "player-update-1",
		JerseyNumber: "11",
	}
	playerJSON, _ := json.Marshal(player)

	req := httptest.NewRequest(http.MethodPut, "/players/1", bytes.NewReader(playerJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	e.Validator = &mockRequestValidator{}
	c := e.NewContext(req, rec)

	mockPlayersService := &mocks.PlayersService{}
	mockPlayersService.On("UpdatePlayer", mock.Anything, player).Return(player, nil)

	api := NewAPI(&mocks.TeamsService{}, mockPlayersService, "", "")
	if assert.NoError(t, api.updatePlayer(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "{\"id\":0,\"team_id\":0,\"name\":\"player-update-1\",\"jersey_number\":\"11\"}\n", rec.Body.String())
	}
}
