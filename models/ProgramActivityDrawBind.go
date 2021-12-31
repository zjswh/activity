package model

import "gorm.io/gorm"

type ProgramActivityDrawBind struct {
	Id     int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	DrawId int    `gorm:"column:drawId;type:int(11);NOT NULL" json:"drawId"`
	Type   string `gorm:"column:type;type:varchar(10);NOT NULL" json:"type"` // 绑定类型
	BindId int    `gorm:"column:bindId;type:int(11);NOT NULL" json:"bindId"`
	Uin    int    `gorm:"column:uin;type:int(11);NOT NULL" json:"uin"`
}

func (m *ProgramActivityDrawBind) TableName() string {
	return "program_activity_draw_bind"
}

func (m *ProgramActivityDrawBind) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *ProgramActivityDrawBind) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *ProgramActivityDrawBind) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}

func GetProgramActivityDrawBindList(Db *gorm.DB, page, num int) ([]ProgramActivityDrawBind, error) {
    var list []ProgramActivityDrawBind
    sql := Db.Model(ProgramActivityDrawBind{})
    if page > 0 && num > 0 {
    sql = sql.Limit(num).Offset((page - 1) * num)
    }
    err := sql.Order("id desc").Find(&list).Error
    return list, err
}

func GetProgramActivityDrawBindCount(Db *gorm.DB) (int64, error) {
    var count int64
    sql := Db.Model(ProgramActivityDrawBind{})
    err := sql.Count(&count).Error
    return count, err
}
