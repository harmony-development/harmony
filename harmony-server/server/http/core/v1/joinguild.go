package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// JoinGuildData is the data for a guild join request
type JoinGuildData struct {
	InviteCode string `validate:"required"`
}

// JoinGuild is the request to join a guild
func (h Handlers) JoinGuild(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(*JoinGuildData)
	guildID, err := h.Deps.DB.ResolveInvite(data.InviteCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "error joining guild, invite code may not exist")
	}
	if err := h.Deps.DB.AddMemberToGuild(ctx.UserID, *guildID); err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to join guild, please try again later")
	}
	if err := h.Deps.DB.IncrementInvite(data.InviteCode); err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error updating invite counter")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"guild": *guildID,
	})
}
