package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// LeaveGuildData is the data for a guild leave request
type LeaveGuildData struct {
	Guild uint64 `validate:"required"`
}

// LeaveGuild unjoins a user from a guild
func (h Handlers) LeaveGuild(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	var data LeaveGuildData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}

	if isOwner, err := h.Deps.DB.IsOwner(data.Guild, ctx.UserID); err != nil || isOwner {
		return echo.NewHTTPError(http.StatusForbidden, "you cannot leave a guild you own")
	}
	if err := h.Deps.DB.DeleteMember(data.Guild, ctx.UserID); err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to leave guild, this issue has been reported")
	}
	h.Deps.State.GuildsLock.Lock()
	delete(h.Deps.State.Guilds[data.Guild].Clients, ctx.UserID)
	h.Deps.State.GuildsLock.Unlock()
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully left guild",
	})
}