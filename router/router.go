
package router

import (
	"github.com/gin-gonic/gin"
	v1 "activity/api/v1"
	"activity/middleware"
)

func InitRouter(Router *gin.RouterGroup) {
	DrawCheckBLoginRouter := Router.Group("").	Use(middleware.CheckBLogin())
	{
		DrawCheckBLoginRouter.POST("/v1/activity/Draw/saveConfig", v1.SaveConfig)
		DrawCheckBLoginRouter.GET("/v1/activity/Draw/getList", v1.GetList)
		DrawCheckBLoginRouter.GET("/v1/activity/Draw/getConfigLists", v1.GetConfigLists)
		DrawCheckBLoginRouter.POST("/v1/activity/Draw/bindDraw", v1.BindDraw)
		DrawCheckBLoginRouter.POST("/v1/activity/Draw/delete", v1.Delete)
		DrawCheckBLoginRouter.POST("/v1/activity/Draw/closeDraw", v1.CloseDraw)
		DrawCheckBLoginRouter.GET("/v1/activity/Draw/getDrawContent", v1.GetDrawContent)
	}

	DrawCheckCLoginRouter := Router.Group("").	Use(middleware.CheckCLogin())
	{
		DrawCheckCLoginRouter.GET("/v1/activity/Draw/getLiveDraw", v1.GetLiveDraw)
		DrawCheckCLoginRouter.POST("/v1/activity/Draw/drawing", v1.Drawing)
	}


}
