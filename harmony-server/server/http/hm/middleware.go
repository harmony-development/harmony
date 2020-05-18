package hm

import (
	"net/http"
	"sync"

	"harmony-server/server/db"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

// A HarmonyContext adds rate limiting and a user ID to an echo.Context
type HarmonyContext struct {
	echo.Context
	Limiter *rate.Limiter
	UserID  uint64
	Data    interface{}
}

// Middlewares contains middlewares for Harmony
type Middlewares struct {
	DB         *db.HarmonyDB
	RateLimits map[string]map[string]*visitor
	RateLock   sync.RWMutex
}

func (h *HarmonyContext) BindAndVerify(v interface{}) error {
	if !h.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many channels being added, please wait a few seconds")
	}
	err := h.Bind(v)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to parse data")
	}
	err = h.Validate(v)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to parse data")
	}
	return nil
}

func (hc *HarmonyContext) VerifyOwner(db *db.HarmonyDB, guildID, userID uint64) error {
	owner, err := db.GetOwner(guildID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to verify ownership, please try again later")
	}
	if owner != userID {
		return echo.NewHTTPError(http.StatusUnauthorized, "insufficient permissions")
	}
	return nil
}

// New instantiates the middlewares for Harmony
func New(db *db.HarmonyDB) *Middlewares {
	m := &Middlewares{
		DB:         db,
		RateLimits: make(map[string]map[string]*visitor),
	}
	go m.RateCleanup()
	return m
}