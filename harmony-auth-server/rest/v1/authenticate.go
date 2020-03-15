package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"harmony-auth-server/authentication"
	"harmony-auth-server/db"
	"harmony-auth-server/rest/hm"
	"harmony-auth-server/types"
	"net/http"
)

// Authenticate takes in a user session and generates an instance-specific session
func Authenticate(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	host, session := ctx.FormValue("host"), ctx.FormValue("session")
	if host == "" || session == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid auth arguments")
	}
	if err := db.VerifySession(session); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
	}
	serverSession := randstr.Hex(16)
	s := &types.Server{IP:host}
	token, err := authentication.MakeServerSessionToken(serverSession, *s.GetIdentity())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating auth token")
	}
	types.Server{IP:host}.SendSession(token) // IMPORTANT : do not ever give the instance a user session!
	return ctx.JSON(http.StatusOK, map[string]string{
		"session": serverSession,
	})
}