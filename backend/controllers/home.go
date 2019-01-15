package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleIndex
func HandleIndex(c *gin.Context) {
	c.String(http.StatusOK, "Please talk to me with api!!")
}
