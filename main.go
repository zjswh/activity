package main

import (
	"activity/config"
	"activity/initialize"
)

func main() {
	config.SetUp()

	//加载数据库
	initialize.Mysql()

	//加载redis
	initialize.Redis()

	initialize.RunServer()
}
