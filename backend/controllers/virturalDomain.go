package controllers

import (
	"../helper"
	"../models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// HandleCreateDomain create domain
func HandleCreateDomain(c *gin.Context) {
	type paramsBody struct {
		Domain       string `json:"domain"`
		Password     string `json:"password"`
		MaxUserCount int    `json:"maxUserCount"`
		MaxUserQuota int    `json:"maxUserQuota"`
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
	if reqInfo.Params.MaxUserCount < 0 {
		reqInfo.Params.MaxUserCount = 1
	}
	if reqInfo.Params.MaxUserCount > models.Config.MailServer.GlobalDomainMaxUserCount {
		reqInfo.Params.MaxUserCount = models.Config.MailServer.GlobalDomainMaxUserCount
	}
	if reqInfo.Params.MaxUserQuota > models.Config.MailServer.GlobalDomainMaxUserQuota {
		reqInfo.Params.MaxUserQuota = models.Config.MailServer.GlobalDomainMaxUserQuota
	}
	if reqInfo.Params.MaxMailCount > models.Config.MailServer.GlobalDomainMaxMailCount {
		reqInfo.Params.MaxMailCount = models.Config.MailServer.GlobalDomainMaxMailCount
	}
	if !helper.ValidateEmail("user@" + reqInfo.Params.Domain) {
		err = errors.New("invalid domain")
	}
	if CheckError(c, err) {
		return
	}
	if ipList, err1 := helper.QueryDomainMX(reqInfo.Params.Domain); err1 != nil || len(ipList) == 0 {
		err = errors.New("domain does not have MX records")
	}
	if CheckError(c, err) {
		return
	}
	passValidator := helper.Password{
		Pass:   reqInfo.Params.Password,
		Length: len(reqInfo.Params.Password),
	}
	passValidator.ProcessPassword()
	if passValidator.Score < 4 {
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
	_, err = models.DaoCreateVirtualDomain(reqInfo.Params.Domain, reqInfo.Params.Password, reqInfo.Params.MaxUserCount,
		reqInfo.Params.MaxUserQuota, reqInfo.Params.MaxMailCount, expireTime)
	if CheckError(c, err) {
		return
	}
	resp.OK = true
	resp.Info = "create domain " + reqInfo.Params.Domain + " OK"
	c.JSON(http.StatusOK, resp)
}

// HandleCreateDomain create domain
func HandleListDomain(c *gin.Context) {
	type paramsBody struct {
		Domain string `json:"domain"`
		SortBy string `json:"sortBy"`
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
	resp.Data, err = models.DaoListVirtualDomain(reqInfo.Params.Domain, reqInfo.Params.SortBy)
	if CheckError(c, err) {
		return
	}
	resp.OK = true
	c.JSON(http.StatusOK, resp)

}
