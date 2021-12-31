package v1

import (
	"activity/service"
	"activity/service/DrawService"
	"activity/types"
	"activity/types/response"
	"github.com/gin-gonic/gin"
	"github.com/zjswh/go-tool/utils"
)

func SaveConfig(c *gin.Context) {
	var ConfigRequest types.DrawConfigRequest
	err := c.ShouldBind(&ConfigRequest)
	if err != nil {
		response.ParamError("参数缺失,"+ err.Error(), c)
		return
	}
	userInfo := service.GetBUserInfo(c)
	err = DrawService.SaveConfig(userInfo, ConfigRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func GetList(c *gin.Context) {
	var ListRequest types.ListRequest
	err := c.ShouldBind(&ListRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}
	userInfo := service.GetBUserInfo(c)
	status := utils.DefaultIntFormValue("status", -1, c)
	ListRequest.Status = status
	list, err := DrawService.GetList(userInfo, ListRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success(list, c)
	return
}

func GetConfigLists(c *gin.Context) {
	var GetConfigListsRequest types.GetConfigListsRequest
	err := c.ShouldBind(&GetConfigListsRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = DrawService.GetConfigLists(GetConfigListsRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func BindDraw(c *gin.Context) {
	var BindDrawRequest types.BindDrawRequest
	err := c.ShouldBind(&BindDrawRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = DrawService.BindDraw(BindDrawRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func Delete(c *gin.Context) {
	var InfoRequest types.InfoRequest
	err := c.ShouldBind(&InfoRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = DrawService.Delete(InfoRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func CloseDraw(c *gin.Context) {
	var CloseDrawRequest types.CloseDrawRequest
	err := c.ShouldBind(&CloseDrawRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = DrawService.CloseDraw(CloseDrawRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func GetDrawContent(c *gin.Context) {
	var InfoRequest types.InfoRequest
	err := c.ShouldBind(&InfoRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = DrawService.GetDrawContent(InfoRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func GetLiveDraw(c *gin.Context) {
	var GetBindRequest types.GetBindRequest
	err := c.ShouldBind(&GetBindRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = DrawService.GetLiveDraw(GetBindRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

func Drawing(c *gin.Context) {
	var DrawingRequest types.DrawingRequest
	err := c.ShouldBind(&DrawingRequest)
	if err != nil {
		response.ParamError("参数缺失", c)
		return
	}

	err = DrawService.Drawing(DrawingRequest)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.Success("", c)
	return
}

