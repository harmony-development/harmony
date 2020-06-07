// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package queries

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/lib/pq"
)

const addForeignUser = `-- name: AddForeignUser :one
INSERT INTO Foreign_Users (User_ID, Home_Server, Local_User_ID)
VALUES ($1, $2, $3)
ON CONFLICT (Local_User_ID) DO UPDATE
    SET Local_User_ID=Foreign_Users.Local_User_ID
RETURNING Local_User_ID
`

type AddForeignUserParams struct {
	UserID      uint64 `json:"user_id"`
	HomeServer  string `json:"home_server"`
	LocalUserID uint64 `json:"local_user_id"`
}

func (q *Queries) AddForeignUser(ctx context.Context, arg AddForeignUserParams) (uint64, error) {
	row := q.queryRow(ctx, q.addForeignUserStmt, addForeignUser, arg.UserID, arg.HomeServer, arg.LocalUserID)
	var local_user_id uint64
	err := row.Scan(&local_user_id)
	return local_user_id, err
}

const addLocalUser = `-- name: AddLocalUser :exec
INSERT INTO Local_Users (User_ID, Email, Password, Instances)
VALUES ($1, $2, $3, $4)
`

type AddLocalUserParams struct {
	UserID    uint64            `json:"user_id"`
	Email     string            `json:"email"`
	Password  []byte            `json:"password"`
	Instances []json.RawMessage `json:"instances"`
}

func (q *Queries) AddLocalUser(ctx context.Context, arg AddLocalUserParams) error {
	_, err := q.exec(ctx, q.addLocalUserStmt, addLocalUser,
		arg.UserID,
		arg.Email,
		arg.Password,
		pq.Array(arg.Instances),
	)
	return err
}

const addProfile = `-- name: AddProfile :exec
INSERT INTO Profiles(User_ID, Username, Avatar, Bio, Status)
VALUES ($1, $2, $3, $4, $5)
`

type AddProfileParams struct {
	UserID   uint64         `json:"user_id"`
	Username string         `json:"username"`
	Avatar   sql.NullString `json:"avatar"`
	Bio      sql.NullString `json:"bio"`
	Status   Userstatus     `json:"status"`
}

func (q *Queries) AddProfile(ctx context.Context, arg AddProfileParams) error {
	_, err := q.exec(ctx, q.addProfileStmt, addProfile,
		arg.UserID,
		arg.Username,
		arg.Avatar,
		arg.Bio,
		arg.Status,
	)
	return err
}

const addUser = `-- name: AddUser :exec
INSERT INTO Users (User_ID)
VALUES ($1)
`

func (q *Queries) AddUser(ctx context.Context, userID uint64) error {
	_, err := q.exec(ctx, q.addUserStmt, addUser, userID)
	return err
}

const emailExists = `-- name: EmailExists :one
SELECT User_ID
FROM Local_Users
WHERE Email = $1
`

func (q *Queries) EmailExists(ctx context.Context, email string) (uint64, error) {
	row := q.queryRow(ctx, q.emailExistsStmt, emailExists, email)
	var user_id uint64
	err := row.Scan(&user_id)
	return user_id, err
}

const getAvatar = `-- name: GetAvatar :one
SELECT Avatar
FROM Profiles
WHERE User_ID = $1
`

func (q *Queries) GetAvatar(ctx context.Context, userID uint64) (sql.NullString, error) {
	row := q.queryRow(ctx, q.getAvatarStmt, getAvatar, userID)
	var avatar sql.NullString
	err := row.Scan(&avatar)
	return avatar, err
}

const getLocalUserID = `-- name: GetLocalUserID :one
SELECT Local_User_ID
FROM Foreign_Users
WHERE User_ID = $1
  AND Home_Server = $2
`

type GetLocalUserIDParams struct {
	UserID     uint64 `json:"user_id"`
	HomeServer string `json:"home_server"`
}

func (q *Queries) GetLocalUserID(ctx context.Context, arg GetLocalUserIDParams) (uint64, error) {
	row := q.queryRow(ctx, q.getLocalUserIDStmt, getLocalUserID, arg.UserID, arg.HomeServer)
	var local_user_id uint64
	err := row.Scan(&local_user_id)
	return local_user_id, err
}

const getUser = `-- name: GetUser :one
SELECT Users.User_ID, Profiles.Username, Profiles.Avatar, Profiles.Bio, Profiles.Status
FROM Users
         INNER JOIN Profiles ON (Users.User_ID = Profiles.User_ID)
WHERE Users.User_ID = $1
`

type GetUserRow struct {
	UserID   uint64         `json:"user_id"`
	Username string         `json:"username"`
	Avatar   sql.NullString `json:"avatar"`
	Bio      sql.NullString `json:"bio"`
	Status   Userstatus     `json:"status"`
}

func (q *Queries) GetUser(ctx context.Context, userID uint64) (GetUserRow, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, userID)
	var i GetUserRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Avatar,
		&i.Bio,
		&i.Status,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT Users.User_ID, Local_Users.Email, Profiles.Username, Profiles.Avatar, Profiles.Bio, Profiles.Status, Local_Users.Password
FROM Local_Users
         INNER JOIN Users
                    ON (Local_Users.User_ID = Users.User_ID)
         INNER JOIN Profiles
                    ON (Local_Users.User_ID = Profiles.User_ID)
WHERE Local_Users.Email = $1
`

type GetUserByEmailRow struct {
	UserID   uint64         `json:"user_id"`
	Email    string         `json:"email"`
	Username string         `json:"username"`
	Avatar   sql.NullString `json:"avatar"`
	Bio      sql.NullString `json:"bio"`
	Status   Userstatus     `json:"status"`
	Password []byte         `json:"password"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Username,
		&i.Avatar,
		&i.Bio,
		&i.Status,
		&i.Password,
	)
	return i, err
}

const setStatus = `-- name: SetStatus :exec
UPDATE Profiles
SET Status=$1
WHERE User_ID = $2
`

type SetStatusParams struct {
	Status Userstatus `json:"status"`
	UserID uint64     `json:"user_id"`
}

func (q *Queries) SetStatus(ctx context.Context, arg SetStatusParams) error {
	_, err := q.exec(ctx, q.setStatusStmt, setStatus, arg.Status, arg.UserID)
	return err
}

const updateAvatar = `-- name: UpdateAvatar :exec
UPDATE Profiles
SET Avatar=$1
WHERE User_ID = $2
`

type UpdateAvatarParams struct {
	Avatar sql.NullString `json:"avatar"`
	UserID uint64         `json:"user_id"`
}

func (q *Queries) UpdateAvatar(ctx context.Context, arg UpdateAvatarParams) error {
	_, err := q.exec(ctx, q.updateAvatarStmt, updateAvatar, arg.Avatar, arg.UserID)
	return err
}

const updateBio = `-- name: UpdateBio :exec
UPDATE Profiles
SET Bio=$1
WHERE User_ID = $2
`

type UpdateBioParams struct {
	Bio    sql.NullString `json:"bio"`
	UserID uint64         `json:"user_id"`
}

func (q *Queries) UpdateBio(ctx context.Context, arg UpdateBioParams) error {
	_, err := q.exec(ctx, q.updateBioStmt, updateBio, arg.Bio, arg.UserID)
	return err
}

const updateUsername = `-- name: UpdateUsername :exec
UPDATE Profiles
SET Username=$1
WHERE User_ID = $2
`

type UpdateUsernameParams struct {
	Username string `json:"username"`
	UserID   uint64 `json:"user_id"`
}

func (q *Queries) UpdateUsername(ctx context.Context, arg UpdateUsernameParams) error {
	_, err := q.exec(ctx, q.updateUsernameStmt, updateUsername, arg.Username, arg.UserID)
	return err
}
