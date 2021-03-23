package hm

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/harmony-development/legato/server/responses"
)

// WithAuth checks for authentication
func (m *Middlewares) WithAuth(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(HarmonyContext)
		session := ctx.Request().Header.Get("Authorization")
		if session == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, responses.BadSession)
		}
		userID, err := m.DB.SessionToUserID(session)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, responses.BadSession)
		}
		ctx.UserID = userID
		return handler(ctx)
	}
}
