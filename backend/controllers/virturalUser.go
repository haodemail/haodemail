package controllers

import (
	"errors"
	"net/http"

	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func comparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}

// HandleLogin login system user
func HandleLogin(c *gin.Context) {
	type paramsBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type params struct {
		Params paramsBody `json:"params"`
	}

	var resp ResponseData
	var reqInfo params
	err := c.BindJSON(&reqInfo)
	if CheckError(c, err) {
		return
	}
	if reqInfo.Params.Username != "" && reqInfo.Params.Password != "" {
		if reqInfo.Params.Username != "postmaster@"+models.Config.MailServer.PrimaryDomain {
			err = errors.New("access denied")
		}
		if CheckError(c, err) {
			return
		}
		user, err := models.DaoGetVirtualUserByName(reqInfo.Params.Username)
		if err == nil {
			if comparePasswords(user.Password, reqInfo.Params.Password) {
				token, err := utils.GenerateToken(user.UserName)
				CheckError(c, err)
				resp.OK = true
				resp.Info = "Authentication success"
				resp.Data = map[string]string{
					"token": token,
				}
			} else {
				resp.OK = false
				resp.Info = "Authentication failed"
			}
		}
	} else {
		resp.OK = false
		resp.Info = "Authentication failed"
	}

	c.JSON(http.StatusOK, resp)
}

// HandleLogout client will cleanup authentication session, so here simple redirect to home
func HandleLogout(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/")
}
