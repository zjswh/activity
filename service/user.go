package service

import (
	"activity/config"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func GetBUserInfo(c *gin.Context) config.UserInfo {
	parse := c.GetString("userInfo")
	userInfo := config.UserInfo{}
	json.Unmarshal([]byte(parse), &userInfo)
	return userInfo
}