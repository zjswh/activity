package model

import "gorm.io/gorm"

type ProgramActivityDrawResult struct {
	Id           int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Uin          int    `gorm:"column:uin;type:int(11);NOT NULL" json:"uin"`
	DrawId       int    `gorm:"column:drawId;type:int(11);NOT NULL" json:"drawId"`
	UserId       int    `gorm:"column:userId;type:int(11);NOT NULL" json:"userId"`
	Phone        string `gorm:"column:phone;type:varchar(11)" json:"phone"`
	UserNick     string `gorm:"column:userNick;type:varchar(20);NOT NULL" json:"userNick"`
	UserIp       string `gorm:"column:userIp;type:varchar(20)" json:"userIp"`
	Address      string `gorm:"column:address;type:varchar(100)" json:"address"` // 收货地址
	Avatar       string `gorm:"column:avatar;type:varchar(200)" json:"avatar"`
	Status       int    `gorm:"column:status;type:tinyint(1);default:0;NOT NULL" json:"status"` // 是否核销 0未1已
	DrawPlay     int    `gorm:"column:drawPlay;type:tinyint(2);NOT NULL" json:"drawPlay"`
	PlayId       int    `gorm:"column:playId;type:int(11);default:0;NOT NULL" json:"playId"` // 场次id
	PrizeId      int    `gorm:"column:prizeId;type:int(11);NOT NULL" json:"prizeId"`
	ExchangeType int    `gorm:"column:exchangeType;type:tinyint(2);default:1;NOT NULL" json:"exchangeType"` // 兑换方式
	PrizeName    string `gorm:"column:prizeName;type:varchar(20);NOT NULL" json:"prizeName"`
	PrizeLevel   int    `gorm:"column:prizeLevel;type:tinyint(2);NOT NULL" json:"prizeLevel"`
	PrizeIcon    string `gorm:"column:prizeIcon;type:varchar(120)" json:"prizeIcon"` // 奖品图标
	DrawTime     int    `gorm:"column:drawTime;type:int(11);NOT NULL" json:"drawTime"`
	Source       string `gorm:"column:source;type:varchar(15);default:live;NOT NULL" json:"source"` // 来源
	SourceId     int    `gorm:"column:sourceId;type:int(11);default:0;NOT NULL" json:"sourceId"`
	Code         string `gorm:"column:code;type:varchar(30)" json:"code"`                 // 核销码
	Name         string `gorm:"column:name;type:varchar(10)" json:"name"`                 // 联系人
	ContactPhone string `gorm:"column:contactPhone;type:varchar(11)" json:"contactPhone"` // 联系电话
	TempId       string `gorm:"column:tempId;type:char(36)" json:"tempId"`                // 临时id
}

func (m *ProgramActivityDrawResult) TableName() string {
	return "program_activity_draw_result"
}

func (m *ProgramActivityDrawResult) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *ProgramActivityDrawResult) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *ProgramActivityDrawResult) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}

func GetProgramActivityDrawResultList(Db *gorm.DB, page, num int) ([]ProgramActivityDrawResult, error) {
    var list []ProgramActivityDrawResult
    sql := Db.Model(ProgramActivityDrawResult{})
    if page > 0 && num > 0 {
    sql = sql.Limit(num).Offset((page - 1) * num)
    }
    err := sql.Order("id desc").Find(&list).Error
    return list, err
}

func GetProgramActivityDrawResultCount(Db *gorm.DB) (int64, error) {
    var count int64
    sql := Db.Model(ProgramActivityDrawResult{})
    err := sql.Count(&count).Error
    return count, err
}
