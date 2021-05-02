package middleware

import (
	"errors"

	"github.com/harmony-development/legato/server/http/responses"
	"github.com/labstack/echo/v4"
	"github.com/ztrue/tracerr"
)

func (m *Middlewares) AuthHandler(c echo.Context) (uint64, error) {
	session := c.Request().Header.Get("Authorization")

	if session == "" {
		session = c.Request().Header.Get("Sec-WebSocket-Protocol")
	}

	userID, err := m.DB.SessionToUserID(session)
	if err != nil {
		println("bad session")
		return 0, errors.New(responses.InvalidSession)
	}
	go func() {
		err := tracerr.Wrap(m.DB.ExtendSession(session))
		if err != nil {
			c.Logger().Error(err)
		}
	}()

	return userID, nil
}
