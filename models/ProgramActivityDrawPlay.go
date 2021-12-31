package model

import "gorm.io/gorm"

type ProgramActivityDrawPlay struct {
	Id             int     `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	DrawId         int     `gorm:"column:drawId;type:int(11);NOT NULL" json:"drawId"`
	Play           int     `gorm:"column:play;type:tinyint(2);default:1;NOT NULL" json:"play"`                 // 场次
	Banner         string  `gorm:"column:banner;type:varchar(100)" json:"banner"`                              // 宣传图
	Url            string  `gorm:"column:url;type:varchar(255)" json:"url"`                                    // 外链地址
	IsShowBanner   int     `gorm:"column:isShowBanner;type:tinyint(1);default:0;NOT NULL" json:"isShowBanner"` // 是否展示banner
	WinningRate    float64 `gorm:"column:winningRate;type:float;default:0;NOT NULL" json:"winningRate"`        // 中奖率
	IsShowPrize    int     `gorm:"column:isShowPrize;type:tinyint(1);default:1;NOT NULL" json:"isShowPrize"`   // 是否展示奖品
	ShowType       int     `gorm:"column:showType;type:int(4);default:1;NOT NULL" json:"showType"`             // 抽奖样式 1九宫格
	Deleted        int     `gorm:"column:deleted;type:tinyint(1);default:0;NOT NULL" json:"deleted"`           // 是否删除
	StartTime      int     `gorm:"column:startTime;type:int(11);NOT NULL" json:"startTime"`
	EndTime        int     `gorm:"column:endTime;type:int(11);NOT NULL" json:"endTime"`
	Condition      int     `gorm:"column:condition;type:tinyint(4);default:0;NOT NULL" json:"condition"`         // 1倒计时 2参与人数
	CountDown      int     `gorm:"column:countDown;type:int(10);default:0;NOT NULL" json:"countDown"`            // 倒计时
	JoinNum        int     `gorm:"column:joinNum;type:int(10);default:0" json:"joinNum"`                         // 参与人数
	SpecialSwitch  int     `gorm:"column:specialSwitch;type:tinyint(1);default:0;NOT NULL" json:"specialSwitch"` // 特殊开关 如口令开关
	SpecialContent string  `gorm:"column:specialContent;type:varchar(50)" json:"specialContent"`                 // 特殊内容
	PrizeType      int     `gorm:"column:prizeType;type:int(3);default:1;NOT NULL" json:"prizeType"`             // 奖品类型 1实物
	ExchangeType   int     `gorm:"column:exchangeType;type:int(4);default:1;NOT NULL" json:"exchangeType"`       // 兑换方式
	ExpiredTime    int     `gorm:"column:expiredTime;type:int(11);default:0" json:"expiredTime"`                 // 有效期 0为永久
}

func (m *ProgramActivityDrawPlay) TableName() string {
	return "program_activity_draw_play"
}

func (m *ProgramActivityDrawPlay) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *ProgramActivityDrawPlay) Update(Db *gorm.DB) error {
	field := []string{"play", "banner", "url", "isShowBanner", "winningRate", "isShowPrize", "showType", "startTime", "endTime",
		"condition", "countDown", "joinNum", "specialSwitch", "specialContent", "prizeType", "exchangeType", "expiredTime"}
	err := Db.Model(&m).Select(field).Where("id", m.Id).Updates(m).Error
    return err
}

func DeleteDrawPlay(Db *gorm.DB, idArr []int) error {
	sql := Db.Model(ProgramActivityDrawPlay{})
	err := sql.Where("id IN (?)", idArr).Select("deleted").Updates(map[string]int{
		"deleted" : 1,
	}).Error
	return err
}

func (m *ProgramActivityDrawPlay) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}

func GetProgramActivityDrawPlayList(Db *gorm.DB, drawId, page, num int) ([]ProgramActivityDrawPlay, error) {
    var list []ProgramActivityDrawPlay
    sql := Db.Model(ProgramActivityDrawPlay{}).Where("drawId = ? AND deleted = 0", drawId)
    if page > 0 && num > 0 {
    sql = sql.Limit(num).Offset((page - 1) * num)
    }
    err := sql.Order("play asc").Find(&list).Error
    return list, err
}

func GetProgramActivityDrawPlayCount(Db *gorm.DB) (int64, error) {
    var count int64
    sql := Db.Model(ProgramActivityDrawPlay{})
    err := sql.Count(&count).Error
    return count, err
}
