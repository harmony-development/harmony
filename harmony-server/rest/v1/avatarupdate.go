package v1

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"gopkg.in/h2non/bimg.v1"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"io/ioutil"
	"net/http"
	"path"
)

func AvatarUpdate(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	form, err := ctx.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Valid form required")
	}
	files := form.File["files"]
	if len(files) == 0 {
		golog.Debugf("Error updating avatar : %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error updating avatar")
	}
	file, err := files[0].Open()
	if err != nil {
		golog.Debugf("Error opening file : %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Error opening uploaded file, does it exist?")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "Too many requests, please try again later")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			golog.Warnf("Error closing file : %v", err)
		}
	}()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		golog.Warnf("Error reading uploaded file : %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading uploaded file")
	}
	scaled, err := bimg.NewImage(fileBytes).Process(bimg.Options{
		Height: 128,
		Width: 128,
		Quality: 60,
		Crop: true,
	})
	if err != nil {
		golog.Warnf("Error reading uploaded file : %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading uploaded file")
	}

	fname := randstr.Hex(16)
	err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", fname), scaled, 0666)
	if err != nil {
		golog.Warnf("Error saving file upload : %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error saving file upload")
	}
	var oldAvatarID string
	err = harmonydb.DBInst.QueryRow("SELECT avatar FROM users WHERE id=$1", *ctx.UserID).Scan(&oldAvatarID)
	_, err = harmonydb.DBInst.Exec("UPDATE users SET avatar=$1 WHERE id=$2", fname, *ctx.UserID)
	if err != nil {
		go deleteFromFilestore(fname)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading uploaded file")
	}
	go deleteFromFilestore(path.Base(oldAvatarID))
	res, err := harmonydb.DBInst.Query("SELECT guildid FROM guildmembers WHERE userid=$1", *ctx.UserID)
	if err != nil {
		golog.Warnf("Error selecting guilds for avatarupdate : %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error broadcasting avatar update")
	}
	for res.Next() {
		var guildid string
		err = res.Scan(&guildid)
		if err != nil {
			golog.Warnf("Error getting guildid from result on avatarupdate : %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Error broadcasting avatar update")
		}
		if globals.Guilds[guildid] != nil {
			for _, client := range globals.Guilds[guildid].Clients  {
				for _, conn := range client {
					conn.Send(&globals.Packet{
						Type: "avatarupdate",
						Data: map[string]interface{}{
							"userid": *ctx.UserID,
							"avatar": fname,
						},
					})
				}
			}
		}
	}
	return nil
}