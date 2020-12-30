// Code generated by sqlc. DO NOT EDIT.
// source: sessions.sql

package queries

import (
	"context"
)

const addSession = `-- name: AddSession :exec
INSERT INTO Sessions (User_ID, Session, Expiration)
VALUES ($1, $2, $3)
`

type AddSessionParams struct {
	UserID     uint64 `json:"user_id"`
	Session    string `json:"session"`
	Expiration int64  `json:"expiration"`
}

func (q *Queries) AddSession(ctx context.Context, arg AddSessionParams) error {
	_, err := q.exec(ctx, q.addSessionStmt, addSession, arg.UserID, arg.Session, arg.Expiration)
	return err
}

const expireSessions = `-- name: ExpireSessions :exec
DELETE FROM Sessions
WHERE Expiration <= $1
`

func (q *Queries) ExpireSessions(ctx context.Context, expiration int64) error {
	_, err := q.exec(ctx, q.expireSessionsStmt, expireSessions, expiration)
	return err
}

const sessionToUserID = `-- name: SessionToUserID :one
SELECT User_ID
FROM Sessions
WHERE Session = $1
`

func (q *Queries) SessionToUserID(ctx context.Context, session string) (uint64, error) {
	row := q.queryRow(ctx, q.sessionToUserIDStmt, sessionToUserID, session)
	var user_id uint64
	err := row.Scan(&user_id)
	return user_id, err
}

const setExpiration = `-- name: SetExpiration :exec
UPDATE Sessions
SET Expiration = $1
WHERE User_ID = $2
`

type SetExpirationParams struct {
	Expiration int64  `json:"expiration"`
	UserID     uint64 `json:"user_id"`
}

func (q *Queries) SetExpiration(ctx context.Context, arg SetExpirationParams) error {
	_, err := q.exec(ctx, q.setExpirationStmt, setExpiration, arg.Expiration, arg.UserID)
	return err
}
