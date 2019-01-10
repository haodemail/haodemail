package controllers

import (
	"net/http"

	"../models"

	"github.com/gin-gonic/gin"
)

// ResponseData API response to client structure
type ResponseData struct {
	OK   bool        `json:"ok"`
	Info string      `json:"info"`
	Data interface{} `json:"data"`
}

// Handle404 404 response as html
func Handle404(c *gin.Context) {
	HandleMessage(c, "<h3>Page Not Found</h3>")
}

// HandleMessage response as json
func HandleMessage(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    404,
		"message": message,
	})
}

// GetUserFromContext get user object from context
func GetUserFromContext(c *gin.Context) (user *models.VirtualUser) {
	if k, exist := c.Get("Username"); exist {
		if user, err := models.DaoGetVirtualUserByName(k.(string)); err == nil {
			return &user
		}
	}
	return
}

// CheckErr response to client when error occurs,
// response 220, show error information in info field and ok set as false
func CheckError(c *gin.Context, err error) (ok bool) {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":   false,
			"info": err.Error(),
			"data": nil,
		})
		ok = true
	}
	return
}
