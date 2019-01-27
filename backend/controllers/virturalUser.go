package controllers

import (
	"errors"
	"net/http"
	"time"

	"../helper"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
)

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
			if models.ComparePassword(user.Password, reqInfo.Params.Password) {
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

// HandleCreateUser create user
func HandleCreateUser(c *gin.Context) {
	type paramsBody struct {
		ID           string `json:"id"`
		DomainID     string `json:"domainID"`
		UserName     string `json:"userName"`
		NickName     string `json:"nickName"`
		Password     string `json:"password"`
		MaxQuota     int    `json:"maxQuota"`
		MaxMailCount int    `json:"maxMailCount"`
		ExpireTime   string `json:"expireTime"`
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
	// check params
	if reqInfo.Params.MaxQuota > models.Config.MailServer.GlobalDomainMaxUserQuota {
		reqInfo.Params.MaxQuota = models.Config.MailServer.GlobalDomainMaxUserQuota
	}
	if reqInfo.Params.MaxMailCount > models.Config.MailServer.GlobalDomainMaxMailCount {
		reqInfo.Params.MaxMailCount = models.Config.MailServer.GlobalDomainMaxMailCount
	}
	passValidator := helper.Password{
		Pass:   reqInfo.Params.Password,
		Length: len(reqInfo.Params.Password),
	}
	passValidator.ProcessPassword()
	if passValidator.Score < 4 && reqInfo.Params.ID == "" {
		err = errors.New("weak password, must contain upper and special character")
	}
	if CheckError(c, err) {
		return
	}
	var expireTime time.Time

	if reqInfo.Params.ExpireTime == "" {
		expireTime = time.Time{}
	} else {
		if t, e := time.Parse("2006-01-02T15:04:05Z07:00", reqInfo.Params.ExpireTime); e == nil {
			expireTime = t
		}
	}
	user, err := models.DaoCreateVirtualUser(reqInfo.Params.DomainID, reqInfo.Params.ID, reqInfo.Params.UserName, reqInfo.Params.Password, reqInfo.Params.NickName, reqInfo.Params.MaxQuota, reqInfo.Params.MaxMailCount, expireTime)
	if CheckError(c, err) {
		return
	}
	resp.OK = true
	domain, err := models.DaoGetVirtualDomainByID(uint(user.DomainID))
	if reqInfo.Params.ID == "" {
		resp.Info = "create user " + user.UserName + "@" + domain.Name + " OK"
	} else {
		resp.Info = "modify user " + user.UserName + "@" + domain.Name + " OK"
	}
	c.JSON(http.StatusOK, resp)
}

// HandleDeleteDomain delete user
// just set deleted flag to user, then remove the mail data sometime
func HandleDeleteUser(c *gin.Context) {
	type paramsBody struct {
		ID string `json:"id"`
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
	user, err := models.DaoDeleteVirtualUser(reqInfo.Params.ID)
	if CheckError(c, err) {
		return
	}
	resp.OK = true
	resp.Info = "delete user " + user.UserName + "@" + user.Domain.Name + " OK"
	c.JSON(http.StatusOK, resp)
}

// HandleListDomain list user
func HandleListUser(c *gin.Context) {
	type paramsBody struct {
		DomainID string `json:"domainID"`
		UserName string `json:"userName"`
		SortBy   string `json:"sortBy"`
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
	resp.Data, err = models.DaoListVirtualUser(reqInfo.Params.DomainID, reqInfo.Params.UserName, reqInfo.Params.SortBy)
	if CheckError(c, err) {
		return
	}
	resp.OK = true
	c.JSON(http.StatusOK, resp)
}
