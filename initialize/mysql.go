package initialize

import (
	"activity/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

func Mysql() {
	conf := config.GVA_CONFIG.Mysql
	link := conf.Username + ":" + conf.Password + "@(" + conf.Path + ")/" + conf.Dbname + "?" + conf.Config
	if db, err := gorm.Open(mysql.Open(link), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,                              // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		fmt.Println("mysql connect failed", err.Error())
		os.Exit(0)
	} else {

		config.GVA_DB = db
		sqlDb, _ := db.DB()
		sqlDb.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDb.SetMaxOpenConns(conf.MaxOpenConns)
		//config.GVA_DB.SingularTable(true)
	}

}

func DBTables() {
	//config.GVA_DB.AutoMigrate(
	//	model.User{},
	//	)
}
