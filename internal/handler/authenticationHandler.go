// Package handler
package handler

import (
	"github.com/AndiVS/game_galaxy/internal/model"
	"github.com/AndiVS/game_galaxy/internal/service"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"time"

	"net/http"
)

// AuthenticationHandler handler for aunt
type AuthenticationHandler struct {
	Service service.Authentication
}

// NewHandlerAuthentication create AuthenticationHandler
func NewHandlerAuthentication(Service service.Authentication) *AuthenticationHandler {
	return &AuthenticationHandler{Service: Service}
}

// SignUp User about cat
func (h *AuthenticationHandler) SignUp(c echo.Context) error {
	account := new(model.Account)

	if err := c.Bind(account); err != nil {
		log.Errorf("Bind fail : %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := h.Service.SignUp(c.Request().Context(), account)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

// SignIn generate token
func (h *AuthenticationHandler) SignIn(c echo.Context) error {
	account := new(model.Account)

	if err := c.Bind(account); err != nil {
		log.Errorf("Bind fail : %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	accessToken, refreshToken, err := h.Service.SignIn(c.Request().Context(), account)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	service.SetTokenCookie(c, "refreshToken", refreshToken, time.Now().Add(1000*time.Second))

	return c.JSON(http.StatusOK, echo.Map{
		"token": accessToken,
	})
}
