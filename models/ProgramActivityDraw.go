package model

import "gorm.io/gorm"

type ProgramActivityDraw struct {
	Id                  int     `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Aid                 int     `gorm:"column:aid;type:int(11);default:0;NOT NULL" json:"aid"` // 频道标志
	Uin                 int     `gorm:"column:uin;type:int(11);NOT NULL" json:"uin"`
	Title               string  `gorm:"column:title;type:varchar(60);NOT NULL" json:"title"`                      // 抽奖主题
	Type                int     `gorm:"column:type;type:tinyint(1);default:1;NOT NULL" json:"type"`               // 抽奖类型  1即时抽奖2定时抽奖
	Times               int     `gorm:"column:times;type:tinyint(2);default:1;NOT NULL" json:"times"`             // 抽奖次数
	Intro               string  `gorm:"column:intro;type:varchar(150)" json:"intro"`                              // 介绍
	WinningInstructions string  `gorm:"column:winningInstructions;type:varchar(150)" json:"winningInstructions"`  // 中奖说明
	LimitNumOfWin       int     `gorm:"column:limitNumOfWin;type:int(7);default:0;NOT NULL" json:"limitNumOfWin"` // 0为不限制
	Condition           int     `gorm:"column:condition;type:tinyint(4);default:0;NOT NULL" json:"condition"`     // 1倒计时 2参与人数
	CountDown           int     `gorm:"column:countDown;type:int(10);default:0;NOT NULL" json:"countDown"`        // 倒计时
	JoinNum             int     `gorm:"column:joinNum;type:int(10);default:0" json:"joinNum"`                     // 参与人数
	WinningRate         float64 `gorm:"column:winningRate;type:float;default:0;NOT NULL" json:"winningRate"`      // 中奖率
	IsShowPrize         int     `gorm:"column:isShowPrize;type:tinyint(1);default:1;NOT NULL" json:"isShowPrize"` // 是否展示奖品
	ShowRate            int     `gorm:"column:showRate;type:int(1);default:0;NOT NULL" json:"showRate"`           // 是否显示中奖率
	Status              int     `gorm:"column:status;type:tinyint(1);default:1;NOT NULL" json:"status"`           // 抽奖状态
	ShowResult          int     `gorm:"column:showResult;type:tinyint(1);default:1;NOT NULL" json:"showResult"`   // 中奖名单是否开启 1开启0关闭
	ShowType            int     `gorm:"column:showType;type:tinyint(1);default:1;NOT NULL" json:"showType"`       // 抽奖样式 1九宫格
	NowPlay             int     `gorm:"column:nowPlay;type:tinyint(1);default:1;NOT NULL" json:"nowPlay"`         // 当前场次
	TotalPlay           int     `gorm:"column:totalPlay;type:tinyint(2);default:1;NOT NULL" json:"totalPlay"`     // 总共场次
	StartTime           int64   `gorm:"column:startTime;type:int(11);NOT NULL" json:"startTime"`                  // 开启时间
	EndTime             int64   `gorm:"column:endTime;type:int(11);NOT NULL" json:"endTime"`                      // 关闭时间
	CreateTime          int64   `gorm:"column:createTime;type:int(11);NOT NULL" json:"createTime"`
	UpdateTime          int64   `gorm:"column:updateTime;type:int(11);NOT NULL" json:"updateTime"`
	STaskId             int     `gorm:"column:sTaskId;type:int(10);default:0" json:"sTaskId"`
	ETaskId             int     `gorm:"column:eTaskId;type:int(10);default:0" json:"eTaskId"`
	IsDeleted           int     `gorm:"column:isDeleted;type:tinyint(1);default:0" json:"isDeleted"`
}

func (m *ProgramActivityDraw) TableName() string {
	return "program_activity_draw"
}

func (m *ProgramActivityDraw) Create(Db *gorm.DB) error {
	err := Db.Model(&m).Create(&m).Error
	return err
}

func (m *ProgramActivityDraw) Update(Db *gorm.DB) error {
	field := []string{"title", "intro", "winningInstructions", "type", "showResult", "limitNumOfWin", "isLimitNumOfWin",
		"totalPlay", "times", "startTime", "endTime", "status", "updateTime"}
	err := Db.Model(&m).Select(field).Where("id", m.Id).Updates(m).Error
	return err
}

func (m *ProgramActivityDraw) GetInfo(Db *gorm.DB) error {
	sql := Db.Model(m).Where("id = ? ", m.Id)
	err := sql.First(&m).Error
	return err
}

func GetProgramActivityDrawList(Db *gorm.DB, title string, uin, status, page, num int, sTime, eTime int64) ([]ProgramActivityDraw, error) {
	var list []ProgramActivityDraw
	sql := Db.Model(ProgramActivityDraw{}).Where("uin = ? AND isDeleted = 0", uin)
	if title != "" {
		sql = sql.Where("title LIKE ?", "%"+title+"%")
	}

	if status != -1 {
		sql = sql.Where("status = ?", status)
	}

	if sTime > 0 && eTime > 0 {
		sql = sql.Where("createTime BETWEEN ? AND ?", sTime, eTime)
	}

	if page > 0 && num > 0 {
		sql = sql.Limit(num).Offset((page - 1) * num)
	}
	err := sql.Order("id desc").Find(&list).Error
	return list, err
}

func GetProgramActivityDrawCount(Db *gorm.DB) (int64, error) {
	var count int64
	sql := Db.Model(ProgramActivityDraw{})
	err := sql.Count(&count).Error
	return count, err
}
