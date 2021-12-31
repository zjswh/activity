package DrawService

import (
	"activity/config"
	model "activity/models"
	"activity/types"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

type ListStruct struct {
	ID         int    `json:"id"`
	Uin        int    `json:"uin"`
	Title      string `json:"title"`
	Type       int    `json:"type"`
	Status     int    `json:"status"`
	StartTime  int    `json:"startTime"`
	EndTime    int    `json:"endTime"`
	TotalPlay  int    `json:"totalPlay"`
	CreateTime int    `json:"createTime"`
	PreviewURL string `json:"previewUrl"`
}

const  (
	InstantLottery = 1
	RegularLottery = 2
	MessageLottery = 3
)

func SaveConfig(userInfo config.UserInfo, req types.DrawConfigRequest) error {
	if req.IsLimitNumOfWin != 0 && req.LimitNumOfWin == 0 {
		return errors.New("限制数不得为0")
	}
	playConfigs := types.PlayConfigs{}
	json.Unmarshal([]byte(req.PlayConfigs), &playConfigs)
	if len(playConfigs) == 0 {
		return errors.New("请设置场次")
	}
	var err error
	//获取第一场次的开始时间以及最后场次的结束时间
	startTime := playConfigs[0].StartTime
	endTime := playConfigs[0].EndTime

	for _, v := range playConfigs {
		prizeNum := len(v.PrizeConfigs)
		if prizeNum == 0 {
			err = errors.New("请输入奖品信息")
			break
		}
		if prizeNum > 8 {
			err = errors.New("奖品超过限制")
			break
		}
		if v.StartTime < startTime {
			startTime = v.StartTime
		}
		if v.EndTime > endTime {
			endTime = v.EndTime
		}
	}
	if err != nil {
		return err
	}

	//定时抽奖 抽奖次数为1
	if req.Type == RegularLottery {
		req.Times = 1
	}
	status := checkStatus(time.Now().Unix(), startTime, endTime)
	drawModel := model.ProgramActivityDraw{
		Uin: userInfo.Uin,
		Aid: userInfo.Aid,
		StartTime: startTime,
		EndTime: endTime,
		Status: status,
	}
	copier.Copy(&drawModel, &req)
	//开启事务
	db := config.GVA_DB.Begin()
	if req.Id > 0 {
		if status == 2 {
			return errors.New("该抽奖已结束,不能修改")
		}
		if status == 1 {
			return errors.New("抽奖已开始,不能修改")
		}
		//更新
		err = drawModel.Update(db)
	} else {
		//新建
		err = drawModel.Create(db)
	}
	if err != nil {
		db.Rollback()
		return errors.New("设置失败")
	}
	//设置场次

	//add your code ...
	return nil
}

func setPlay(drawId int, configs types.PlayConfigs, db *gorm.DB) error {
	//获取已配置的场次信息
	oldPlayConfigList, err := model.GetProgramActivityDrawPlayList(config.GVA_DB, drawId,0, 0)
	if err != nil {
		return err
	}
	oldPlayIdMap := map[int]int{}
	if len(oldPlayConfigList) > 0 {
		for _, v := range oldPlayConfigList {
			oldPlayIdMap[v.Id] = 0
		}
	}

	for k, v := range configs {
		info := model.ProgramActivityDrawPlay{}
		copier.Copy(&info, &v)
		info.DrawId = drawId
		info.Play = k + 1
		if v.ID == 0 {
			err = info.Create(db)
			//新建
		} else {
			//更新
			if _, ok := oldPlayIdMap[v.ID]; ok {
				delete(oldPlayIdMap, v.ID)
				err = info.Update(db)
				//更新缓存
				config.GVA_REDIS.Del(fmt.Sprintf("draw_prize_info_%d_%d", drawId, v.ID))
				config.GVA_REDIS.Del(fmt.Sprintf("draw_play_info:%d_%d", drawId, v.ID))
				config.GVA_REDIS.Del(fmt.Sprintf("draw_play_info_by_now_play:%d_%d", drawId, info.Play))
			}
		}
		if err != nil {
			break
		}

		//设置奖品


	}

	if err != nil {
		return err
	}

	if len(oldPlayIdMap) > 0 {
		deletePlayIdArr := []int{}
		for k, _ := range oldPlayIdMap {
			deletePlayIdArr = append(deletePlayIdArr, k)
		}
		//删除场次
		model.DeleteDrawPlay(db, deletePlayIdArr)
	}
	return nil
}

func checkStatus(nowTime, sTime, eTime int64) int {
	status := 0
	if nowTime > eTime {
		status = 2
	} else if nowTime > sTime {
		status = 1
	}
	return status
}

func GetList(userInfo config.UserInfo, req types.ListRequest) (interface{}, error) {
	list := []ListStruct{}
	tmpList, err := model.GetProgramActivityDrawList(config.GVA_DB, req.Title, userInfo.Uin,
		req.Status,  req.Page, req.Num, req.CreateSTime, req.CreateETime)
	if err != nil {
		return list, err
	}
	cname := "https://web.guangdianyun.tv"
	for _ , v := range tmpList {
		info := ListStruct{}
		copier.Copy(&info, &v)
		info.PreviewURL = fmt.Sprintf("%s/lottery/list?id=%d&uin=%d", cname, info.ID, info.Uin)
		if info.Type == MessageLottery {
			//todo 查询绑定关系
			info.PreviewURL = fmt.Sprintf("%s/live/%d?uin=%d", cname, info.ID, info.Uin)
		}
		list = append(list, info)
	}
	return list, err
}

func GetConfigLists(req types.GetConfigListsRequest) error {
	//add your code ...
	return nil
}

func BindDraw(req types.BindDrawRequest) error {
	//add your code ...
	return nil
}

func Delete(req types.InfoRequest) error {
	//add your code ...
	return nil
}

func CloseDraw(req types.CloseDrawRequest) error {
	//add your code ...
	return nil
}

func GetDrawContent(req types.InfoRequest) error {
	//add your code ...
	return nil
}

func GetLiveDraw(req types.GetBindRequest) error {
	//add your code ...
	return nil
}

func Drawing(req types.DrawingRequest) error {
	//add your code ...
	return nil
}

