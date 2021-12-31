package model

import "gorm.io/gorm"

type ProgramActivityDrawPrize struct {
	Id          int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	DrawId      int    `gorm:"column:drawId;type:int(11);NOT NULL" json:"drawId"`           // 抽奖Id
	PlayId      int    `gorm:"column:playId;type:int(11);default:0;NOT NULL" json:"playId"` // 场次id
	PrizeAlias  string `gorm:"column:prizeAlias;type:varchar(10)" json:"prizeAlias"`        // 奖项别名
	Level       int    `gorm:"column:level;type:int(4);default:1;NOT NULL" json:"level"`    // 奖项级别
	Name        string `gorm:"column:name;type:varchar(15);NOT NULL" json:"name"`           // 奖品名称
	Num         int    `gorm:"column:num;type:int(4);default:0;NOT NULL" json:"num"`        // 已抽中的奖品数量
	Sum         int    `gorm:"column:sum;type:int(4);default:0;NOT NULL" json:"sum"`        // 奖品总数
	Type        int    `gorm:"column:type;type:tinyint(4);default:1;NOT NULL" json:"type"`  // 奖品类型 1实物
	TypeInfo    string `gorm:"column:typeInfo;type:varchar(20)" json:"typeInfo"`
	Icon        string `gorm:"column:icon;type:varchar(150)" json:"icon"` // 奖品图标
	Url         string `gorm:"column:url;type:varchar(255)" json:"url"`
	Deleted     int    `gorm:"column:deleted;type:tinyint(1);default:0;NOT NULL" json:"deleted"`     // 是否删除
	WinningRate int    `gorm:"column:winningRate;type:int(4);default:0;NOT NULL" json:"winningRate"` // 中奖率
	CreateTime  int    `gorm:"column:createTime;type:int(11);NOT NULL" json:"createTime"`
	UpdateTime  int    `gorm:"column:updateTime;type:int(11);NOT NULL" json:"updateTime"`
}

func (m *ProgramActivityDrawPrize) TableName() string {
	return "program_activity_draw_prize"
}

func (m *ProgramActivityDrawPrize) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *ProgramActivityDrawPrize) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *ProgramActivityDrawPrize) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}

func GetProgramActivityDrawPrizeList(Db *gorm.DB, page, num int) ([]ProgramActivityDrawPrize, error) {
    var list []ProgramActivityDrawPrize
    sql := Db.Model(ProgramActivityDrawPrize{})
    if page > 0 && num > 0 {
    sql = sql.Limit(num).Offset((page - 1) * num)
    }
    err := sql.Order("id desc").Find(&list).Error
    return list, err
}

func GetProgramActivityDrawPrizeCount(Db *gorm.DB) (int64, error) {
    var count int64
    sql := Db.Model(ProgramActivityDrawPrize{})
    err := sql.Count(&count).Error
    return count, err
}
