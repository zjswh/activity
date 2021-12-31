package middleware

import (
	"activity/config"
	"activity/types/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zjswh/go-tool/utils"
)

func CheckBLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			token = c.Query("token")
		}

		if token == "" {
			response.Result(1, "", "未登录", c)
			c.Abort()
			return
		}
		url := config.GVA_CONFIG.Param.BGatewayHost + "/v1/Passport/Index/getLoginInfo"
		url = fmt.Sprintf("%s?token=%s&path=/%s&method=%s", url, token, c.FullPath(), c.Request.Method)
		res, _ := utils.RequestGet(url)
		var result response.Response
		json.Unmarshal(res, &result)
		if result.Code != 200 || result.ErrorCode != 0 {
			response.Result(result.ErrorCode, "", result.ErrorMessage, c)
			c.Abort()
			return
		}

		userInfo, _ := json.Marshal(result.Data)
		c.Set("userInfo", string(userInfo))
		c.Next()
	}
}


func CheckCLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//edit your code...
		
		c.Next()
	}
}


